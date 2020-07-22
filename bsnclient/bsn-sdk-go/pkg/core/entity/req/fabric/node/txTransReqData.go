package node

import "bsn-sdk-go/pkg/core/entity/base"

type TxTransReqData struct {
	base.BaseReqModel
	Body TxTransReqDataBody `json:"body"` // 消息体
}

type TxTransReqDataBody struct {
	TxId string `json:"txId"` // 交易Id
}

func (f *TxTransReqData) GetEncryptionValue() string {
	return f.GetBaseEncryptionValue() + f.Body.TxId
}
