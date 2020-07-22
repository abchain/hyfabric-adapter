package node

import "bsn-sdk-go/pkg/core/entity/base"

type LedgerReqData struct {
	base.BaseReqModel
	Body LedgerReqDataBody `json:"body"` // 消息体
}

type LedgerReqDataBody struct {
}

func (f *LedgerReqData) GetEncryptionValue() string {

	return f.GetBaseEncryptionValue()

}
