package channel

import (
	"fmt"
	
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/spf13/viper"
)

type RpcSpec struct {
	//notice chaincode name is different to the ccname in txgenerator, the later
	//is used in the hyperledger-project compatible tx
	ChaincodeName string
	Attributes    []string
	Options       *viper.Viper
}

const (
	invokeFunc = "invoke"
	queryFunc  = "query"
)

type Client struct {
	chainCodeId string
	chClient    *channel.Client
}

func NewChannelClient(ccName string, cli *channel.Client) *Client {
	return &Client{
		chainCodeId: ccName,
		chClient:    cli,
	}
}

// Query chaincode using method and request options.
func (c *Client) Query(method string, arg [][]byte) ([]byte, error) {
	resp, err := c.chClient.Query(channel.Request{
		ChaincodeID: c.chainCodeId,
		Fcn:         invokeFunc,
		Args:        append([][]byte{[]byte(queryFunc)}, append([][]byte{[]byte(method)}, arg...)...),
		//TransientMap:    nil,
		//InvocationChain: nil,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

// Invoke executes transaction using method and request options.
func (c *Client) Invoke(method string, arg [][]byte) (string, error) {
	resp, err := c.chClient.Execute(channel.Request{
		ChaincodeID: c.chainCodeId,
		Fcn:         invokeFunc,
		Args:        append([][]byte{[]byte(method)}, arg...),
		//TransientMap:    nil,
		//InvocationChain: nil,
	})
	if err != nil {
		return "", err
	}

	return string(resp.TransactionID), nil
}

// Deploy is actually instantiates chaincode
func (c *Client) Deploy(method string, arg [][]byte) (string, error) {
	return "", fmt.Errorf("deploy no implement")
}
