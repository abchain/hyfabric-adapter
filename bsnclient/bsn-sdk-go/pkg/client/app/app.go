package app

import (
	"hyperledger.abchain.org/adapter/hyfabric/bsnclient/bsn-sdk-go/pkg/core/entity/req"
	"hyperledger.abchain.org/adapter/hyfabric/bsnclient/bsn-sdk-go/pkg/core/entity/res"
	"hyperledger.abchain.org/adapter/hyfabric/bsnclient/bsn-sdk-go/pkg/util/http"
	"encoding/json"
)

func GetAppInfo(data *req.AppInfoReqData, baseApi string,cert string) (*res.AppInfoResData,error) {

	url := baseApi + "/api/app/getAppInfo"

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, cert)

	if err !=nil{
		return nil,err
	}

	res := &res.AppInfoResData{}

	err = json.Unmarshal(resBytes, res)

	if err != nil {
		return nil,err
	}

	return res,nil

}
