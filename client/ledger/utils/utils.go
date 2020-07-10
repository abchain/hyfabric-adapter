package utils

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/pkg/errors"
)

// GetEnvelopeFromBlock gets an envelope from a block's Data field.
func GetEnvelopeFromData(data []byte) (*common.Envelope, error) {
	// Block always begins with an envelope
	env := &common.Envelope{}
	err := proto.Unmarshal(data, env)
	return env, errors.Wrapf(err, "error unmarshaling envelope")
}

// GetProposal returns a Proposal message from its bytes
func GetProposal(propBytes []byte) (*peer.Proposal, error) {
	prop := &peer.Proposal{}
	err := proto.Unmarshal(propBytes, prop)
	return prop, errors.Wrapf(err, "error unmarshaling proposal")
}

// GetChaincodeInvocationSpec get the ChaincodeInvocationSpec from the proposal
func GetChaincodeInvocationSpec(prop []byte) (*peer.ChaincodeInvocationSpec, error) {
	if prop == nil {
		return nil, fmt.Errorf("proposal is nil")
	}

	ccPropPayload, err := GetChaincodeProposalPayload(prop)
	if err != nil {
		return nil, err
	}
	cis := &peer.ChaincodeInvocationSpec{}
	err = proto.Unmarshal(ccPropPayload.Input, cis)
	return cis, errors.Wrapf(err, "error unmarshaling chaincodeInvocationSpec")
}

// GetChaincodeAction gets the ChaincodeAction given chaicnode action bytes
func GetChaincodeAction(caBytes []byte) (*peer.ChaincodeAction, error) {
	chaincodeAction := &peer.ChaincodeAction{}
	err := proto.Unmarshal(caBytes, chaincodeAction)
	return chaincodeAction, errors.Wrapf(err, "error unmarshaling ChaincodeAction")
}

// GetChaincodeEvents gets the ChaincodeEvents given chaincode event bytes
func GetChaincodeEvents(eBytes []byte) (*peer.ChaincodeEvent, error) {
	chaincodeEvent := &peer.ChaincodeEvent{}
	err := proto.Unmarshal(eBytes, chaincodeEvent)
	return chaincodeEvent, errors.Wrapf(err, "error unmarshaling ChaicnodeEvent")
}

// GetPayload Get Payload from Envelope message
func GetPayload(e *common.Envelope) (*common.Payload, error) {
	payload := &common.Payload{}
	err := proto.Unmarshal(e.Payload, payload)
	return payload, errors.Wrapf(err, "error unmarshaling Payload")
}

// UnmarshalChannelHeader returns a ChannelHeader from bytes
func UnmarshalChannelHeader(bytes []byte) (*common.ChannelHeader, error) {
	chdr := &common.ChannelHeader{}
	err := proto.Unmarshal(bytes, chdr)
	return chdr, errors.Wrapf(err, "error unmarshaling ChannelHeader")
}

// GetTransaction Get Transaction from bytes
func GetTransaction(txBytes []byte) (*peer.Transaction, error) {
	tx := &peer.Transaction{}
	err := proto.Unmarshal(txBytes, tx)
	return tx, errors.Wrapf(err, "error unmarshaling Transaction")

}

// GetChaincodeActionPayload Get ChaincodeActionPayload from bytes
func GetChaincodeActionPayload(capBytes []byte) (*peer.ChaincodeActionPayload, error) {
	cap := &peer.ChaincodeActionPayload{}
	err := proto.Unmarshal(capBytes, cap)
	return cap, errors.Wrapf(err, "error unmarshaling ChaincodeActionPayload")
}

// GetProposalResponsePayload gets the proposal response payload
func GetProposalResponsePayload(prpBytes []byte) (*peer.ProposalResponsePayload, error) {
	prp := &peer.ProposalResponsePayload{}
	err := proto.Unmarshal(prpBytes, prp)
	return prp, errors.Wrapf(err, "error unmarshaling ProposalResponsePayload")
}

// GetHeader Get Header from bytes
func GetHeader(bytes []byte) (*common.Header, error) {
	hdr := &common.Header{}
	err := proto.Unmarshal(bytes, hdr)
	return hdr, errors.Wrapf(err, "error unmarshaling Header")
}

// GetChaincodeProposalPayload Get ChaincodeProposalPayload from bytes
func GetChaincodeProposalPayload(bytes []byte) (*peer.ChaincodeProposalPayload, error) {
	cpp := &peer.ChaincodeProposalPayload{}
	err := proto.Unmarshal(bytes, cpp)
	return cpp, errors.Wrapf(err, "error unmarshaling ChaincodeProposalPayload")
}
