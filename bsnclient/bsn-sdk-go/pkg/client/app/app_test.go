package app

import (
	"bsn-sdk-go/pkg/core/entity/base"
	"bsn-sdk-go/pkg/core/entity/req"
	"fmt"
	"testing"
)

func TestGetAppInfo(t *testing.T) {

	api := "http://192.168.1.43:17502"

	reqData := req.AppInfoReqData{}

	header := base.ReqHeader{
		UserCode: "USER0001202004161009309407413",
		AppCode:  "app0001202004161017141233920",
	}

	reqData.Header = header

	res,_ := GetAppInfo(&reqData, api)

	if res.Header.Code != 0 {
		fmt.Println(res.Header.Msg)
	} else {
		fmt.Println(res.Body)
	}
}
