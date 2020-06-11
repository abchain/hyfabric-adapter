package chaincode

import (
	"time"

	fashim "github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
	yashim "hyperledger.abchain.org/chaincode/shim"
)

type faChainCode struct {
	yaCc yashim.Chaincode
}

// AdaptChainCode convert abchain chaincode to fabric chaincode
func AdaptChainCode(yaCc yashim.Chaincode) fashim.Chaincode {
	return &faChainCode{yaCc: yaCc}
}

func (f *faChainCode) Init(stub fashim.ChaincodeStubInterface) peer.Response {
	return fashim.Success(nil)
}

func (f *faChainCode) Invoke(stub fashim.ChaincodeStubInterface) peer.Response {
	function := ""
	args := stub.GetArgs()
	if len(args) > 0 {
		function = string(args[0])
		args = args[1:]
	}
	res, err := f.yaCc.Invoke(&adaptChainCodeStub{stub}, function, args, false)
	if err != nil {
		return fashim.Error(err.Error())
	}
	return fashim.Success(res)
}

// helper
type adaptChainCodeStub struct {
	fashim.ChaincodeStubInterface
}

func (s *adaptChainCodeStub) GetTxTime() (time.Time, error) {
	// todo(mh):  convert timestamp.Timestamp to time.Time
	return time.Time{}, nil
}

func (s *adaptChainCodeStub) RangeQueryState(startKey, endKey string) (yashim.StateRangeQueryIteratorInterface, error) {
	state, err := s.GetStateByRange(startKey, endKey)
	return &adaptState{state: state}, err
}

func (s *adaptChainCodeStub) GetRawStub() interface{} {
	return s.ChaincodeStubInterface
}

type adaptState struct {
	state fashim.StateQueryIteratorInterface
}

func (s *adaptState) HasNext() bool {
	return s.state.HasNext()
}

func (s *adaptState) Close() error {
	return s.state.Close()
}

func (s *adaptState) Next() (string, []byte, error) {
	next, err := s.state.Next()
	return next.Key, next.Value, err
}
