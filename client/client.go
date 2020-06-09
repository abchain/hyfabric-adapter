package hyfabric

import (
	fchannel "github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	fledger "github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/spf13/viper"
	"hyperledger.abchain.org/adapter/hyfabric/client/channel"
	"hyperledger.abchain.org/adapter/hyfabric/client/ledger"
)

type hyFabricClient struct {
	chainInfo ChainInfo
	caller    Caller
	sdk       *fabsdk.FabricSDK
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
	return &hyFabricClient{}
}

//Load 利用viper 加载配置   配置文件参考core.yaml
func (c *hyFabricClient) Load(vp *viper.Viper) error {
	configPath := vp.GetString("configpath")
	sdk, err := fabsdk.New(config.FromFile(configPath))
	if err != nil {
		return err
	}

	channelId := vp.GetString("channelid")
	user := vp.GetString("user")
	org := vp.GetString("org")
	chC := sdk.ChannelContext(channelId, fabsdk.WithUser(user), fabsdk.WithOrg(org))

	ledgerCli, err := fledger.New(chC)
	if err != nil {
		return err
	}

	channelCli, err := fchannel.New(chC)
	if err != nil {
		return err
	}

	ccName := vp.GetString("chaincode")
	c.caller = channel.NewChannelClient(ccName, channelCli)
	c.chainInfo = ledger.NewLedgerClient(ledgerCli)
	c.sdk = sdk
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
	c.sdk.Close()
}
