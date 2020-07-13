package client

import (
	fchannel "github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	fledger "github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/spf13/viper"

	"hyperledger.abchain.org/adapter/hyfabric/client/channel"
	"hyperledger.abchain.org/adapter/hyfabric/client/ledger"
	"hyperledger.abchain.org/chaincode/lib/caller"
	"hyperledger.abchain.org/client"
)

func init() {
	if client.Client_Impls == nil {
		client.Client_Impls = make(map[string]func() client.RpcClient)
	}
	client.Client_Impls["hyfabric"] = NewRPCConfig
}

const (
	userType   = "user"
	registered = "is already registered"
)

type hyFabricClient struct {
	chainInfo client.ChainInfo
	caller    rpc.Caller
	sdk       *fabsdk.FabricSDK
}

func NewRPCConfig() client.RpcClient {
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

	// get user, if user is not exist, register and enroll.
	mspClient, err := msp.New(sdk.Context())
	if err != nil {
		return err
	}
	_, err = mspClient.GetSigningIdentity(user)
	if err == msp.ErrUserNotFound {
		secret := vp.GetString("secret")
		err = mspClient.Enroll(user, msp.WithSecret(secret))
		if err != nil {
			return err
		}
	}

	chC := sdk.ChannelContext(channelId, fabsdk.WithUser(user), fabsdk.WithOrg(org))
	// get fabric ledger client
	ledgerCli, err := fledger.New(chC)
	if err != nil {
		return err
	}

	// get fabric channel client
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

// Caller Assign each http request (run cocurrency) a client, which can be adapted to a caller
// the client is "lazy" connect: it just do connect when required (a request has come)
// and wait for connect finish
func (c *hyFabricClient) Caller(spec *client.RpcSpec) (rpc.Caller, error) {
	return c.caller, nil
}

func (c *hyFabricClient) Chain() (client.ChainInfo, error) {
	return c.chainInfo, nil
}

func (c *hyFabricClient) Quit() {
	c.sdk.Close()
}
