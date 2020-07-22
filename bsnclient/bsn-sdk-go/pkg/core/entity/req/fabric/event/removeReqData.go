package event

import "bsn-sdk-go/pkg/core/entity/base"

type RemoveReqData struct {
	base.BaseReqModel
	Body RemoveReqDataBody `json:"body"` // 消息体

}

type RemoveReqDataBody struct {
	EventId string `json:"eventId"` //事件Id
}

func (f *RemoveReqData) GetEncryptionValue() string {
	return f.GetBaseEncryptionValue() + f.Body.EventId

}
