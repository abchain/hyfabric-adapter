package event

import "hyperledger.abchain.org/adapter/hyfabric/bsnclient/bsn-sdk-go/pkg/core/entity/base"

type QueryReqData struct {
	base.BaseReqModel
	Body interface{} `json:"body"` // 消息体

}

func (f *QueryReqData) GetEncryptionValue() string {
	return f.GetBaseEncryptionValue()

}
