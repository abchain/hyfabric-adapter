package fabric

import (
	"bsn-sdk-go/pkg/client"
	"bsn-sdk-go/pkg/core/config"
	"bsn-sdk-go/pkg/core/entity/msp"
	"bsn-sdk-go/pkg/util/keystore"
	"bsn-sdk-go/pkg/util/userstore"
	"github.com/wonderivan/logger"
)

//初始化 fabric请求的client
func InitFabricClient(config *config.Config) (*FabricClient, error) {

	//初始化配置信息
	if err := config.Init(); err != nil {
		logger.Error("Configuration initialization failed")
		return nil, err
	}
	//生成一个私钥处理程序
	ks, err := keystore.NewFileBasedKeyStore(nil, config.GetKSPath(), false)

	if err != nil {
		logger.Error("keystore initialization failed")
		return nil, err
	}
	//生成一个证书处理程序
	us := userstore.NewUserStore(config.GetUSPath())

	client := client.Client{
		Ks:     ks,
		Us:     us,
		Config: config,
		Users:  make(map[string]*msp.UserData),
	}

	fabricClient := &FabricClient{client}
	//设置用户签名算法类型，同时生成统一的签名和验签处理程序
	err = fabricClient.SetAlgorithm(config.GetAppInfo().AlgorithmType, config.GetAppCert().AppPublicCert, config.GetAppCert().UserAppPrivateCert)

	if err != nil {
		logger.Error("signHandle initialization failed")
		return nil, err
	}
	//加载本地已经生成的子用户信息
	fabricClient.LoadUser()

	return fabricClient, nil
}

type FabricClient struct {
	client.Client
}
