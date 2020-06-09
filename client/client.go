package hyfabric

import (
	"github.com/spf13/viper"
	"hyperledger.abchain.org/adapter/hyfabric/client/channel"
	"hyperledger.abchain.org/adapter/hyfabric/client/ledger"
)

type hyFabricClient struct {
	chainInfo ChainInfo
	caller    Caller
}

type RpcClient interface {
	Chain() (ChainInfo, error)
	Caller(*channel.RpcSpec) (Caller, error)
	Load(*viper.Viper) error
	Quit()
}

type ChainInfo interface {
	GetChain() (*ledger.Chain, error)
	GetBlock(int64) (*ledger.ChainBlock, error)
	GetTransaction(string) (*ledger.ChainTransaction, error)
	GetTxEvent(string) ([]*ledger.ChainTxEvents, error)
}

type Caller interface {
	Deploy(method string, arg [][]byte) (string, error)
	Invoke(method string, arg [][]byte) (string, error)
	Query(method string, arg [][]byte) ([]byte, error)
}

func NewRPCConfig() *hyFabricClient {
	return &hyFabricClient{
		chainInfo: ledger.NewLedgerClient(),
		caller:    channel.NewChannelClient(nil),
	}
}

//Load 利用viper 加载配置   配置文件参考core.yaml
func (c *hyFabricClient) Load(vp *viper.Viper) error {
	return nil
}

//Caller Assign each http request (run cocurrency) a client, which can be adapted to a caller
//the client is "lazy" connect: it just do connect when required (a request has come)
//and wait for connect finish
func (c *hyFabricClient) Caller(spec *channel.RpcSpec) (Caller, error) {
	return c.caller, nil
}

func (c *hyFabricClient) Chain() (ChainInfo, error) {
	return c.chainInfo, nil
}

func (c *hyFabricClient) Quit() {
}
