package user

import (
	"hyperledger.abchain.org/adapter/hyfabric/bsnclient/bsn-sdk-go/pkg/core/entity/base"
)

type RegisterReqData struct {
	base.BaseReqModel
	Body RegisterReqDataBody `json:"body"` // 消息体
}

type RegisterReqDataBody struct {
	Name   string `json:"name"`   //帐号 小于20位
	Secret string `json:"secret"` //密码
}

func (f *RegisterReqData) GetEncryptionValue() string {
	fp := f.GetBaseEncryptionValue()

	fp = fp + f.Body.Name
	fp = fp + f.Body.Secret

	return fp

}
