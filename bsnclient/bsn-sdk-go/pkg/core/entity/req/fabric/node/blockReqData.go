package node

import (
	"bsn-sdk-go/pkg/core/entity/base"
	"strconv"
)

type BlockReqData struct {
	base.BaseReqModel
	Body BlockReqDataBody `json:"body"` // 消息体
}

type BlockReqDataBody struct {
	BlockNumber uint64 `json:"blockNumber"` // 块 号
	BlockHash   string `json:"blockHash"`   // 块Hash
	TxId        string `json:"txId"`        // 交易Id

}

func (f *BlockReqData) GetEncryptionValue() string {

	fp := f.GetBaseEncryptionValue()
	fp = fp + strconv.FormatUint(f.Body.BlockNumber, 10)
	fp = fp + f.Body.BlockHash
	fp = fp + f.Body.TxId
	return fp

}
