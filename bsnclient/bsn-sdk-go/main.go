/**
 * @Author: Gao Chenxi
 * @Description:
 * @Date: 2020/4/1 4:35 PM
 * @File: main
 */
package main

import (
	"bsn-sdk-go/pkg/client/fabric"
	"bsn-sdk-go/pkg/core/config"
	"bsn-sdk-go/pkg/core/entity/req/fabric/user"
	"fmt"
	"log"
)

func main()  {

	fmt.Println("bsn-sdk-go")

	api:="" //节点网关地址
	userCode:="" //用户编号
	appCode :="" //应用编号
	puk :="" //应用公钥
	prk :="" //应用私钥
	mspDir:="" //证书存数目录
	cert :="" //证书

	config,err :=config.NewConfig(api, userCode, appCode, puk, prk, mspDir, cert )
	if err !=nil{
		log.Fatal(err)
	}

	client,err :=fabric.InitFabricClient(config)
	if err !=nil{
		log.Fatal(err)
	}
	req :=user.RegisterReqDataBody{
		Name:"",
		Secret:"",
	}

	res,err :=client.RegisterUser(req)
	if err !=nil{
		log.Fatal(err)
	}

	if res.Header.Code != 0{
		log.Fatal( res.Header.Msg)
	}



}