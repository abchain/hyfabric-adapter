package msp

import (
	"hyperledger.abchain.org/adapter/hyfabric/bsnclient/bsn-sdk-go/pkg/common/errors"
	pb_msp "hyperledger.abchain.org/adapter/hyfabric/bsnclient/bsn-sdk-go/third_party/github.com/hyperledger/fabric/protos/msp"
	"crypto/ecdsa"
	"github.com/golang/protobuf/proto"
)

type UserData struct {
	UserName              string
	AppCode               string
	MspId                 string
	EnrollmentCertificate []byte

	PrivateKey *ecdsa.PrivateKey
}

func (u *UserData) Serialize() ([]byte, error) {
	serializedIdentity := &pb_msp.SerializedIdentity{
		Mspid:   u.MspId,
		IdBytes: u.EnrollmentCertificate,
	}
	identity, err := proto.Marshal(serializedIdentity)
	if err != nil {
		return nil, errors.New("marshal serializedIdentity failed")
	}
	return identity, nil
}
