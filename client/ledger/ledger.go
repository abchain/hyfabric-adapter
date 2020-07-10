package ledger

import (
	"fmt"
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/sirupsen/logrus"
	"hyperledger.abchain.org/client"

	"hyperledger.abchain.org/adapter/hyfabric/client/ledger/utils"
)

const BlockMetadataIndex_TRANSACTIONS_FILTER int32 = 2

var log = logrus.New()

type faClient struct {
	*ledger.Client
}

func NewLedgerClient(cli *ledger.Client) *faClient {
	return &faClient{cli}
}

func (c *faClient) GetChain() (*client.Chain, error) {
	info, err := c.QueryInfo()
	if err != nil {
		return nil, err
	}

	return &client.Chain{Height: int64(info.BCI.Height)}, nil
}

func (c *faClient) GetBlock(height int64) (*client.ChainBlock, error) {
	block, err := c.QueryBlock(uint64(height))
	if err != nil {
		return nil, err
	}

	cBlock := &client.ChainBlock{
		Height:       height,
		Hash:         fmt.Sprintf("%.32X", block.GetHeader().GetDataHash()),
		PreviousHash: fmt.Sprintf("%.32X", block.GetHeader().GetPreviousHash()),
	}

	transFilter := block.GetMetadata().GetMetadata()[BlockMetadataIndex_TRANSACTIONS_FILTER]
	for i, data := range block.GetData().GetData() {
		envelope, err := utils.GetEnvelopeFromData(data)
		if err != nil {
			log.Warnf("reconstructing envelope error: (%s)", err)
			continue
		}

		// only the valid transaction will be recorded in the block transactions.
		if peer.TxValidationCode(transFilter[i]) == peer.TxValidationCode_VALID {
			transaction, err := envelopeToTransaction(height, envelope)
			if err != nil {
				log.Warnf("error get transaction from envelope error: (%s)", err)
				continue
			}
			cBlock.Transactions = append(cBlock.Transactions, transaction)
		}

		txEvent, err := envelopeToTxEvents(envelope)
		if err != nil {
			log.Warnf("get transaction from envelope error: (%s)", err)
			continue
		}
		txEvent.Status = int(transFilter[i])
		cBlock.TxEvents = append(cBlock.TxEvents, txEvent)
	}
	return cBlock, nil
}

func (c *faClient) GetTransaction(txid string) (*client.ChainTransaction, error) {
	transactionId := fab.TransactionID(txid)
	txPro, err := c.QueryTransaction(transactionId)
	if err != nil {
		return nil, err
	}

	block, err := c.QueryBlockByTxID(transactionId)
	if err != nil {
		return nil, err
	}

	return envelopeToTransaction(int64(block.GetHeader().GetNumber()), txPro.TransactionEnvelope)
}

// Note! GetTxEvent cannot get the status of tx. this may be added in later versions.
// todo(mh): add tx status?.
func (c *faClient) GetTxEvent(txid string) ([]*client.ChainTxEvents, error) {
	transactionId := fab.TransactionID(txid)
	txPro, err := c.QueryTransaction(transactionId)
	if err != nil {
		return nil, err
	}
	env, err := envelopeToTxEvents(txPro.TransactionEnvelope)
	return []*client.ChainTxEvents{env}, err
}

func envelopeToTransaction(height int64, env *common.Envelope) (*client.ChainTransaction, error) {
	ccActionPayload, txId, isEndorserTransaction, err := getChainCodeActionPayloadFromEnvelope(env)
	if err != nil {
		return nil, fmt.Errorf("invalid chaincode action in payload for tx %v : %v", txId, err)
	}

	// todo(MH): currently only concerned about endorser transaction.
	if !isEndorserTransaction {
		return nil, fmt.Errorf("unfollowed transacction type")
	}

	spec, err := utils.GetChaincodeInvocationSpec(ccActionPayload.GetChaincodeProposalPayload())
	if err != nil {
		return nil, fmt.Errorf("get chaincode invocation spec error: %v", err)
	}

	return &client.ChainTransaction{
		Height:      height,
		TxID:        txId,
		Chaincode:   spec.GetChaincodeSpec().GetChaincodeId().GetName(),
		Method:      string(spec.GetChaincodeSpec().GetInput().GetArgs()[0]),
		CreatedFlag: isEndorserTransaction,
		TxArgs:      spec.GetChaincodeSpec().GetInput().GetArgs()[1:],
	}, nil
}

func envelopeToTxEvents(env *common.Envelope) (*client.ChainTxEvents, error) {
	ccActionPayload, txId, isEndorserTransaction, err := getChainCodeActionPayloadFromEnvelope(env)
	if err != nil {
		return nil, fmt.Errorf("invalid chaincode action in payload for tx %v : %v", txId, err)
	}

	if !isEndorserTransaction {
		return nil, fmt.Errorf("no HeaderType_ENDORSER_TRANSACTION type ")
	}

	resp, err := utils.GetProposalResponsePayload(ccActionPayload.GetAction().GetProposalResponsePayload())
	if err != nil {

	}
	caPayload, err := utils.GetChaincodeAction(resp.GetExtension())
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling chaincode action for block event: %s", err)
	}
	ccEvent, err := utils.GetChaincodeEvents(caPayload.GetEvents())
	if ccEvent != nil {
		return &client.ChainTxEvents{
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
	payload, err := utils.GetPayload(env)
	if err != nil {
		return nil, txId, isEndorserTransaction, fmt.Errorf("invalid payload: %v", err)
	}

	chdr, err := utils.UnmarshalChannelHeader(payload.GetHeader().GetChannelHeader())
	if err != nil {
		return nil, txId, isEndorserTransaction, fmt.Errorf("invalid channel header: %v", err)
	}

	txId = chdr.GetTxId()
	if chdr.Type != int32(common.HeaderType_ENDORSER_TRANSACTION) {
		isEndorserTransaction = false
		return nil, txId, isEndorserTransaction, nil
	}

	tx, err := utils.GetTransaction(payload.GetData())
	if err != nil {
		return nil, txId, isEndorserTransaction, fmt.Errorf("invalid transaction in payload data for tx: %v. error: %v", chdr.TxId, err)
	}

	ccActionPayload, err := utils.GetChaincodeActionPayload(tx.GetActions()[0].GetPayload())
	if err != nil {
		return nil, txId, isEndorserTransaction, fmt.Errorf("invalid chaincode action in payload for tx %v : %v", chdr.TxId, err)
	}

	if ccActionPayload.Action == nil {
		return nil, txId, isEndorserTransaction, fmt.Errorf("action in ChaincodeActionPayload for %v is nil", chdr.TxId)
	}
	return ccActionPayload, txId, isEndorserTransaction, nil
}
