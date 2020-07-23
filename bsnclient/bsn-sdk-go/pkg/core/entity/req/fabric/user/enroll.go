package user

import "hyperledger.abchain.org/adapter/hyfabric/bsnclient/bsn-sdk-go/pkg/core/entity/base"

type EnrollReqData struct {
	base.BaseReqModel
	Body EnrollReqDataBody `json:"body"` // 消息体
}

type EnrollReqDataBody struct {
	Name   string `json:"name"`   //帐号 小于20位
	Secret string `json:"secret"` //密码
	CsrPem string `json:"csrPem"`
}

func (f *EnrollReqData) GetEncryptionValue() string {

	fp := f.GetBaseEncryptionValue()

	fp = fp + f.Body.Name
	fp = fp + f.Body.Secret
	fp = fp + f.Body.CsrPem

	return fp

}
