package ledger

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/peer"
	futils "github.com/hyperledger/fabric/protos/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

var log = logrus.New()

type RpcClient interface {
	Chain() (ChainInfo, error)
	Caller(*RpcSpec) (Caller, error)
	Load(*viper.Viper) error
	Quit()
}

type ChainClient interface {
	ViaWeb(*viper.Viper) ChainInfo
}

type ChainInfo interface {
	GetChain() (*Chain, error)
	GetBlock(int64) (*ChainBlock, error)
	GetTransaction(string) (*ChainTransaction, error)
	GetTxEvent(string) ([]*ChainTxEvents, error)
}

type Client struct {
	*ledger.Client
}

//
func (c *Client) GetChain() (*Chain, error) {
	info, err := c.QueryInfo()
	if err != nil {
		return nil, err
	}

	return &Chain{Height: int64(info.BCI.Height)}, nil
}

func (c *Client) GetBlock(height int64) (*ChainBlock, error) {
	block, err := c.QueryBlock(uint64(height))
	if err != nil {
		return nil, err
	}

	cBlock := &ChainBlock{
		Height:       height,
		Hash:         string(block.GetHeader().GetDataHash()),
		PreviousHash: string(block.GetHeader().GetPreviousHash()),
	}

	for _, data := range block.GetData().GetData() {
		envelope, err := futils.GetEnvelopeFromBlock(data)
		if err != nil {
			log.Warnf("reconstructing envelope error: (%s)", err)
			continue
		}

		transaction, err := envelopeToTrasaction(height, envelope)
		if err == nil {
			log.Warnf("error get transaction from envelope error: (%s)", err)
			cBlock.Transactions = append(cBlock.Transactions, transaction)
		}

		txEvent, err := envelopeToTxEvents(envelope)
		if err == nil {
			log.Warnf("get transaction from envelope error: (%s)", err)
			cBlock.TxEvents = append(cBlock.TxEvents, txEvent)
		}
	}
	return cBlock, nil
}

func (c *Client) GetTransaction(txid string) (*ChainTransaction, error) {
	transactionId := fab.TransactionID(txid)
	txPro, err := c.QueryTransaction(transactionId)
	if err != nil {
		return nil, err
	}

	block, err := c.QueryBlockByTxID(transactionId)
	if err != nil {
		return nil, err
	}

	return envelopeToTrasaction(int64(block.GetHeader().GetNumber()), (*common.Envelope)(txPro.TransactionEnvelope))
}

func (c *Client) GetTxEvent(txid string) (*ChainTxEvents, error) {
	transactionId := fab.TransactionID(txid)
	txPro, err := c.QueryTransaction(transactionId)
	if err != nil {
		return nil, err
	}
	return envelopeToTxEvents((*common.Envelope)(txPro.TransactionEnvelope))
}

type Chain struct {
	// may be uint is better
	Height int64
}

type ChainBlock struct {
	Height       int64  `json:",string"`
	Hash         string `json:",omitempty"`
	PreviousHash string
	TimeStamp    time.Time           `json:",omitempty"`
	Transactions []*ChainTransaction `json:"-"`
	TxEvents     []*ChainTxEvents    `json:"-"`
}

type ChainTransaction struct {
	Height                  int64 `json:",string"`
	TxID, Chaincode, Method string
	CreatedFlag             bool
	TxArgs                  [][]byte `json:"-"`
}

type ChainTxEvents struct {
	TxID, Chaincode, Name string
	Status                int
	Payload               []byte `json:"-"`
}

type Caller interface {
	Deploy(method string, arg [][]byte) (string, error)
	Invoke(method string, arg [][]byte) (string, error)
	Query(method string, arg [][]byte) ([]byte, error)
}

type RpcSpec struct {
	//notice chaincode name is different to the ccname in txgenerator, the later
	//is used in the hyperledger-project compatible tx
	ChaincodeName string
	Attributes    []string
	Options       *viper.Viper
}

func envelopeToTrasaction(height int64, env *common.Envelope) (*ChainTransaction, error) {
	ccActionPayload, txId, isEndorserTransaction, err := getChainCodeActionPayloadFromEnvelope(env)
	if err != nil {
		return nil, fmt.Errorf("invalid chaincode action in payload for tx %v : %v", txId, err)
	}

	pro, err := futils.GetProposal(ccActionPayload.GetChaincodeProposalPayload())
	if err != nil {
		return nil, fmt.Errorf("get proposal error: %v", err)
	}

	spec, err := futils.GetChaincodeInvocationSpec(pro)
	if err != nil {
		return nil, fmt.Errorf("get chaincode invocation spec error: %v", err)
	}
	return &ChainTransaction{
		Height:      height,
		TxID:        txId,
		Chaincode:   spec.GetChaincodeSpec().GetChaincodeId().GetName(),
		Method:      string(spec.GetChaincodeSpec().GetInput().GetArgs()[0]),
		CreatedFlag: isEndorserTransaction,
		TxArgs:      spec.GetChaincodeSpec().GetInput().GetArgs()[1:],
	}, nil
}

func envelopeToTxEvents(env *common.Envelope) (*ChainTxEvents, error) {
	ccActionPayload, txId, isEndorserTransaction, err := getChainCodeActionPayloadFromEnvelope(env)
	if err != nil {
		return nil, fmt.Errorf("invalid chaincode action in payload for tx %v : %v", txId, err)
	}

	if !isEndorserTransaction {
		return nil, fmt.Errorf("no HeaderType_ENDORSER_TRANSACTION type ")
	}

	resp, err := futils.GetProposalResponsePayload(ccActionPayload.GetAction().GetProposalResponsePayload())
	if err != nil {

	}
	caPayload, err := futils.GetChaincodeAction(resp.GetExtension())
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling chaincode action for block event: %s", err)
	}
	ccEvent, err := futils.GetChaincodeEvents(caPayload.GetEvents())
	if ccEvent != nil {
		return &ChainTxEvents{
			TxID:      txId,
			Chaincode: ccEvent.GetChaincodeId(),
			Name:      ccEvent.GetEventName(),
			// todo(mh): define status
			Payload: ccEvent.GetPayload(),
		}, nil
	}
	return nil, fmt.Errorf("chaincode event is nil, err: %v", err)
}

func getChainCodeActionPayloadFromEnvelope(env *common.Envelope) (*peer.ChaincodeActionPayload, string, bool, error) {
	var isEndorserTransaction = true
	var txId = ""
	payload, err := futils.GetPayload(env)
	if err != nil {
		return nil, txId, isEndorserTransaction, fmt.Errorf("invalid payload: %v", err)
	}

	chdr, err := futils.UnmarshalChannelHeader(payload.GetHeader().GetChannelHeader())
	if err != nil {
		return nil, txId, isEndorserTransaction, fmt.Errorf("invalid channel header: %v", err)
	}

	txId = chdr.GetTxId()
	if chdr.Type != int32(common.HeaderType_ENDORSER_TRANSACTION) {
		isEndorserTransaction = false
	}

	tx, err := futils.GetTransaction(payload.GetData())
	if err != nil {
		return nil, txId, isEndorserTransaction, fmt.Errorf("invalid transaction in payload data for tx: %v. error: %v", chdr.TxId, err)
	}

	ccActionPayload, err := futils.GetChaincodeActionPayload(tx.GetActions()[0].GetPayload())
	if err != nil {
		return nil, txId, isEndorserTransaction, fmt.Errorf("invalid chaincode action in payload for tx %v : %v", chdr.TxId, err)
	}

	if ccActionPayload.Action == nil {
		return nil, txId, isEndorserTransaction, fmt.Errorf("action in chaincodeActionPayload for %v is nil", chdr.TxId, )
	}
	return ccActionPayload, txId, isEndorserTransaction, nil
}
