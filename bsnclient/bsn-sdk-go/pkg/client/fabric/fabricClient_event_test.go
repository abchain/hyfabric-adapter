package fabric

import (
	config2 "hyperledger.abchain.org/adapter/hyfabric/bsnclient/bsn-sdk-go/pkg/core/config"
	req "hyperledger.abchain.org/adapter/hyfabric/bsnclient/bsn-sdk-go/pkg/core/entity/req/fabric/event"
	"fmt"
	"testing"
)

func TestFabricClient_EventRegister(t *testing.T) {

	config, err := config2.NewMockConfig()

	if err != nil {
		t.Fatal(err.Error())
	}

	fabricClient, err := InitFabricClient(config)

	if err != nil {
		t.Fatal(err.Error())
	}

	body := req.RegisterReqDataBody{
		ChainCode:  "cc_app0006202004071529586812466_00",
		EventKey:   "test",
		NotifyUrl:  "http://127.0.0.1",
		AttachArgs: "a=1",
	}

	res, _ := fabricClient.EventRegister(body)

	fmt.Println(res)

}
