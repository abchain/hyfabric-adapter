package channel

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
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
		Fcn:         method,
		Args:        arg,
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
		Fcn:         method,
		Args:        arg,
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
