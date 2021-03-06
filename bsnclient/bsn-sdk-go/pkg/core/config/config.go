package config

import (
	"hyperledger.abchain.org/adapter/hyfabric/bsnclient/bsn-sdk-go/pkg/client/app"
	"hyperledger.abchain.org/adapter/hyfabric/bsnclient/bsn-sdk-go/pkg/common/errors"
	"hyperledger.abchain.org/adapter/hyfabric/bsnclient/bsn-sdk-go/pkg/core/entity/base"
	"hyperledger.abchain.org/adapter/hyfabric/bsnclient/bsn-sdk-go/pkg/core/entity/enum"
	"hyperledger.abchain.org/adapter/hyfabric/bsnclient/bsn-sdk-go/pkg/core/entity/req"
	"path"
)

const (
	_KeyStore = "keystore"
)

// Create a profile information
// api:节点网关的地址
// userCode:用户编号
// appCode:应用编号
// puk :节点网关的公钥证书
// prk :应用的私钥证书
// cert:https证书的文件路径
func NewConfig(api, userCode, appCode, puk, prk, mspDir, cert string) (*Config, error) {

	config := &Config{
		nodeApi:  api,
		mspDir:   mspDir,
		httpCert: cert,
		appCert:  certInfo{AppPublicCert: puk, UserAppPrivateCert: prk},
		user:     userInfo{UserCode: userCode},
		app:      appInfo{AppCode: appCode},
	}
	config.Init()
	return config, nil
}

//NewMockConfig
func NewMockConfig() (*Config, error) {

	config := &Config{
		nodeApi:  " ",
		mspDir:   " ",
		httpCert: "cert/bsn_gateway_https.crt",
		appCert: certInfo{
			AppPublicCert: `-----BEGIN CERTIFICATE-----
MIIC+zCCAqGgAwIBAgIUARhAfFSyhzcx9q4LdiYKl2UHo1YwCgYIKoZIzj0EAwIw
TjELMAkGA1UEBhMCQ04xEDAOBgNVBAgTB0JlaWppbmcxDDAKBgNVBAoTA0JTTjEP
MA0GA1UECxMGY2xpZW50MQ4wDAYDVQQDEwVic25jYTAgFw0xOTA5MjYxMDI0MDBa
GA8yMDk5MDkwNTAyMDQwMFowgZYxCzAJBgNVBAYTAkNOMREwDwYDVQQIEwhDaGFu
Z3NoYTEOMAwGA1UEChMFQ21QYXkxPTALBgNVBAsTBHVzZXIwEgYDVQQLEwtob25n
emFvbm9kZTAOBgNVBAsTB2JzbmJhc2UwCgYDVQQLEwNjb20xJTAjBgNVBAMMHG5v
ZGVAaG9uZ3phb25vZGUuYnNuYmFzZS5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMB
BwNCAAQ/X2w5+pJoZXNCO81T4xMR+TxmFoYk6eG1kYML8HBPrUT6QflxtDXYsE9h
SzVAovq5DHww3vD8Xft/mxwsAXyuo4IBEDCCAQwwDgYDVR0PAQH/BAQDAgeAMAwG
A1UdEwEB/wQCMAAwHQYDVR0OBBYEFDPVPRqPANJavkNOg/WhPkUkH6wqMB8GA1Ud
IwQYMBaAFJuwcYba1G07p1ySkpzyes8L79OPMCUGA1UdEQQeMByCGmNhLmhvbmd6
YW9ub2RlLmJzbmJhc2UuY29tMIGEBggqAwQFBgcIAQR4eyJhdHRycyI6eyJoZi5B
ZmZpbGlhdGlvbiI6Imhvbmd6YW9ub2RlLmJzbmJhc2UuY29tIiwiaGYuRW5yb2xs
bWVudElEIjoibm9kZUBob25nemFvbm9kZS5ic25iYXNlLmNvbSIsImhmLlR5cGUi
OiJ1c2VyIn19MAoGCCqGSM49BAMCA0gAMEUCIQD7FBAQJsgS0uhaL4mjJgILdFfY
RKXvNutyKz/MqJ54RQIgNS67sSUCbOZRx1rWDqYEcBF1zypEFik25fNgY3zk5gM=
-----END CERTIFICATE-----`,
			UserAppPrivateCert: ` `,
		},
		user: userInfo{
			UserCode: " ",
		},
		app: appInfo{
			AppCode: " ",
		},
	}

	err :=config.Init()

	if err !=nil {
		return nil,err
	}

	return config, nil
}

//NewMockConfig
func NewMock1Config() (*Config, error) {

	config := &Config{
		nodeApi:  "http://192.168.1.43:17502",
		mspDir:   " test/msp",
		httpCert: "",
		appCert: certInfo{
			AppPublicCert: `-----BEGIN CERTIFICATE-----
MIIC+zCCAqGgAwIBAgIUARhAfFSyhzcx9q4LdiYKl2UHo1YwCgYIKoZIzj0EAwIw
TjELMAkGA1UEBhMCQ04xEDAOBgNVBAgTB0JlaWppbmcxDDAKBgNVBAoTA0JTTjEP
MA0GA1UECxMGY2xpZW50MQ4wDAYDVQQDEwVic25jYTAgFw0xOTA5MjYxMDI0MDBa
GA8yMDk5MDkwNTAyMDQwMFowgZYxCzAJBgNVBAYTAkNOMREwDwYDVQQIEwhDaGFu
Z3NoYTEOMAwGA1UEChMFQ21QYXkxPTALBgNVBAsTBHVzZXIwEgYDVQQLEwtob25n
emFvbm9kZTAOBgNVBAsTB2JzbmJhc2UwCgYDVQQLEwNjb20xJTAjBgNVBAMMHG5v
ZGVAaG9uZ3phb25vZGUuYnNuYmFzZS5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMB
BwNCAAQ/X2w5+pJoZXNCO81T4xMR+TxmFoYk6eG1kYML8HBPrUT6QflxtDXYsE9h
SzVAovq5DHww3vD8Xft/mxwsAXyuo4IBEDCCAQwwDgYDVR0PAQH/BAQDAgeAMAwG
A1UdEwEB/wQCMAAwHQYDVR0OBBYEFDPVPRqPANJavkNOg/WhPkUkH6wqMB8GA1Ud
IwQYMBaAFJuwcYba1G07p1ySkpzyes8L79OPMCUGA1UdEQQeMByCGmNhLmhvbmd6
YW9ub2RlLmJzbmJhc2UuY29tMIGEBggqAwQFBgcIAQR4eyJhdHRycyI6eyJoZi5B
ZmZpbGlhdGlvbiI6Imhvbmd6YW9ub2RlLmJzbmJhc2UuY29tIiwiaGYuRW5yb2xs
bWVudElEIjoibm9kZUBob25nemFvbm9kZS5ic25iYXNlLmNvbSIsImhmLlR5cGUi
OiJ1c2VyIn19MAoGCCqGSM49BAMCA0gAMEUCIQD7FBAQJsgS0uhaL4mjJgILdFfY
RKXvNutyKz/MqJ54RQIgNS67sSUCbOZRx1rWDqYEcBF1zypEFik25fNgY3zk5gM=
-----END CERTIFICATE-----`,
			UserAppPrivateCert: ` `,
		},
		user: userInfo{
			UserCode: " ",
		},
		app: appInfo{
			AppCode: " ",
		},
	}

	err :=config.Init()

	if err !=nil {
		return nil,err
	}

	return config, nil
}

type Config struct {
	nodeApi string
	mspDir  string

	user userInfo
	app  appInfo

	//应用证书【bsn的节点网关公钥、用户的应用私钥】
	appCert certInfo

	//https的连接证书
	httpCert string

	isInit bool
}

func (c *Config) GetAppInfo() appInfo {
	return c.app
}

func (c *Config) GetCert() string {
	return c.httpCert
}

func (c *Config) GetAppCert() certInfo {
	return c.appCert
}

func (c *Config) GetKSPath() string {
	return path.Join(c.mspDir, _KeyStore)
}

func (c *Config) GetUSPath() string {
	return c.mspDir
}

func (c *Config) GetNodeApi() string {
	return c.nodeApi
}

func (c *Config) GetReqHeader() base.ReqHeader {
	header := base.ReqHeader{
		UserCode: c.user.UserCode,
		AppCode:  c.app.AppCode,
	}

	return header
}

func (c *Config) Init() error {
	if !c.isInit {
		reqData := req.AppInfoReqData{}

		reqData.Header = c.GetReqHeader()

		reqData.Body = req.AppInfoReqDataBody{}
		res,err := app.GetAppInfo(&reqData, c.nodeApi,c.httpCert)

		if err !=nil {
			return err
		}

		if res.Header.Code != 0 {
			return errors.New("get app info failed ：" + res.Header.Msg)
		}

		c.app.AppType = res.Body.AppType

		c.app.CAType = enum.App_CaType(res.Body.CaType)
		c.app.AlgorithmType = enum.App_AlgorithmType(res.Body.AlgorithmType)

		c.app.MspId = res.Body.MspId

		c.app.ChannelId = res.Body.ChannelId
		c.isInit = true
	}

	return nil
}

type certInfo struct {
	//bsn的应用公钥证书
	AppPublicCert string

	//用户的私钥证书
	UserAppPrivateCert string
}

type appInfo struct {
	AppCode string
	AppType string

	CAType        enum.App_CaType
	AlgorithmType enum.App_AlgorithmType

	//AppCertPuk string

	MspId     string
	ChannelId string
}

type userInfo struct {
	UserCode string
}

type orgConfig struct {
	OrgCode string
	MspId   string

	NodeApi string
}
