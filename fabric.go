package adapter_hyfabric

import (
	"github.com/spf13/viper"
	"time"
)

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

type Chain struct {
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
