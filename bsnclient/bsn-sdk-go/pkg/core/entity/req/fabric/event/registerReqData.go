/**
 * @Author: Gao Chenxi
 * @Description:
 * @File:  registerReqData
 * @Version: 1.0.0
 * @Date: 2020/4/14 15:42
 */

package event

import "hyperledger.abchain.org/adapter/hyfabric/bsnclient/bsn-sdk-go/pkg/core/entity/base"

type RegisterReqData struct {
	base.BaseReqModel
	Body RegisterReqDataBody `json:"body"` // 消息体

}

// 链码事件请求体
type RegisterReqDataBody struct {
	ChainCode  string `json:"chainCode"`  // 链码Code
	EventKey   string `json:"eventKey"`   // 事件key
	NotifyUrl  string `json:"notifyUrl"`  //通知url
	AttachArgs string `json:"attachArgs"` //附加参数
}

func (f *RegisterReqData) GetEncryptionValue() string {

	fp := f.GetBaseEncryptionValue()

	fp = fp + f.Body.ChainCode
	fp = fp + f.Body.EventKey
	fp = fp + f.Body.NotifyUrl
	fp = fp + f.Body.AttachArgs

	return fp

}
