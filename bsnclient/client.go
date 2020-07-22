package bsnclient

import (
	msp "bsn-sdk-go/pkg/core/entity/req/fabric/user"
	"encoding/base64"
	"fmt"
	"github.com/spf13/viper"

	"bsn-sdk-go/pkg/client/fabric"
	"bsn-sdk-go/pkg/core/config"
	"bsn-sdk-go/pkg/core/entity/req/fabric/node"
	"hyperledger.abchain.org/chaincode/lib/caller"
	"hyperledger.abchain.org/client"
)

func init() {
	if client.Client_Impls == nil {
		client.Client_Impls = make(map[string]func() client.RpcClient)
	}
	client.Client_Impls["bsnfabric"] = NewRPCConfig
}

type hyFabricClient struct {
	client    *fabric.FabricClient
	name      string
	chaincode string
}

func NewRPCConfig() client.RpcClient {
	return &hyFabricClient{}
}

//Load 利用viper 加载配置   配置文件参考core.yaml
func (c *hyFabricClient) Load(vp *viper.Viper) error {

	cfg, err := config.NewConfig(
		vp.GetString("api"), vp.GetString("usercode"),
		vp.GetString("appcode"), vp.GetString("puk"),
		vp.GetString("prk"), vp.GetString("mspdir"), vp.GetString("cert"))
	if err != nil {
		return err
	}

	client, err := fabric.InitFabricClient(cfg)
	if err != nil {
		return err
	}

	user := vp.GetString("user")

	enrollReq := msp.RegisterReqDataBody{
		Name:   user,
		Secret: vp.GetString("secret"),
	}

	res, err := client.RegisterUser(enrollReq)
	if err != nil {
		return err
	}
	if res.Header.Code != 0 {
		return fmt.Errorf(res.Header.Msg)
	}

	c.name = user
	c.chaincode = vp.GetString("chaincode")
	c.client = client

	return nil
}

// Caller Assign each http request (run concurrency) a client, which can be adapted to a caller
// the client is "lazy" connect: it just do connect when required (a request has come)
// and wait for connect finish
func (c *hyFabricClient) Caller(spec *client.RpcSpec) (rpc.Caller, error) {
	return c, nil
}

func (c *hyFabricClient) Chain() (client.ChainInfo, error) {
	return c, nil
}

func (c *hyFabricClient) Quit() {

}

func (c *hyFabricClient) Deploy(method string, arg [][]byte) (string, error) {
	return "", fmt.Errorf("deploy no implement")
}

func (c *hyFabricClient) Invoke(method string, arg [][]byte) (string, error) {
	var args []string
	for _, v := range arg {
		args = append(args, base64.RawStdEncoding.EncodeToString(v))
	}

	body := node.TransReqDataBody{
		UserName:     c.name,
		ChainCode:    c.chaincode,
		FuncName:     method,
		Args:         args,
		TransientMap: make(map[string]string),
	}

	res, err := c.client.SdkTran(body)
	if err != nil {
		return "", err
	}
	return res.Body.BlockInfo.TxId, nil
}

func (c *hyFabricClient) Query(method string, arg [][]byte) ([]byte, error) {
	var args []string
	for _, v := range arg {
		args = append(args, base64.RawStdEncoding.EncodeToString(v))
	}

	body := node.TransReqDataBody{
		UserName:     c.name,
		ChainCode:    c.chaincode,
		FuncName:     method,
		Args:         args,
		TransientMap: make(map[string]string),
	}

	res, err := c.client.SdkTran(body)
	if err != nil {
		return nil, err
	}
	return []byte(res.Body.CCRes.CCData), nil
}

func (c *hyFabricClient) GetChain() (*client.Chain, error) {
	info, err := c.client.GetLedgerInfo()
	if err != nil {
		return nil, err
	}
	return &client.Chain{
		Height: int64(info.Body.Height),
	}, nil
}

func (c *hyFabricClient) GetBlock(i int64) (*client.ChainBlock, error) {
	req := node.BlockReqDataBody{
		BlockNumber: uint64(i),
	}

	res, err := c.client.GetBlockInfo(req)
	if err != nil {
		return nil, err
	}

	// todo(mh): query tx detail?
	var txs []*client.ChainTransaction
	for _, tx := range res.Body.Transactions {
		txs = append(txs, &client.ChainTransaction{
			Height: i,
			TxID:   tx.TxId,
		})
	}

	return &client.ChainBlock{
		Height:       i,
		Hash:         res.Body.BlockHash,
		PreviousHash: res.Body.PreBlockHash,
		Transactions: txs,
		// TxEvents:     nil,
	}, nil
}

func (c *hyFabricClient) GetTransaction(s string) (*client.ChainTransaction, error) {
	req := node.TxTransReqDataBody{
		TxId: s,
	}

	res, err := c.client.GetTransInfo(req)
	if err != nil {
		return nil, err
	}

	return &client.ChainTransaction{
		Height: int64(res.Body.BlockNumber),
		TxID:   s,
	}, nil
}

func (c *hyFabricClient) GetTxEvent(s string) ([]*client.ChainTxEvents, error) {
	return nil, nil
}
