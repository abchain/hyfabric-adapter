package event

import "hyperledger.abchain.org/adapter/hyfabric/bsnclient/bsn-sdk-go/pkg/core/entity/base"

type RegisterResData struct {
	base.BaseResModel
	Body *RegisterResDataBody `json:"body"`
}

type RegisterResDataBody struct {
	EventId string `json:"eventId"`
}

func (f *RegisterResData) GetEncryptionValue() string {
	if f.Body == nil {
		return f.GetBaseEncryptionValue()
	}
	return f.GetBaseEncryptionValue() + f.Body.EventId
}
