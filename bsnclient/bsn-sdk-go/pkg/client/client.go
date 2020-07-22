package client

import (
	"bsn-sdk-go/pkg/common/errors"
	"bsn-sdk-go/pkg/core/config"
	"bsn-sdk-go/pkg/core/entity/base"
	"bsn-sdk-go/pkg/core/entity/enum"
	"bsn-sdk-go/pkg/core/entity/msp"
	"bsn-sdk-go/pkg/core/sign"
	"bsn-sdk-go/pkg/util/esdsa"
	"bsn-sdk-go/pkg/util/keystore"
	"bsn-sdk-go/pkg/util/sm2"
	"bsn-sdk-go/pkg/util/userstore"
	"bsn-sdk-go/third_party/github.com/hyperledger/fabric/bccsp"

	"github.com/wonderivan/logger"
)

type Client struct {
	Config *config.Config

	Ks bccsp.KeyStore
	Us userstore.UserStore

	Users map[string]*msp.UserData

	sign sign.Crypto
}

func (c *Client) SetAlgorithm(algorithmType enum.App_AlgorithmType, puk, pri string) error {
	switch algorithmType {
	case enum.AppAlgorithmType_SM2:
		sh, err := sm2.NewSM2Handle(puk, pri)
		if err != nil {
			return err
		} else {
			c.sign = sign.NewCrypto(sh)
			return nil
		}
	case enum.AppAlgorithmType_R1:
		sh, err := ecdsa.NewEcdsaR1Handle(puk, pri)
		if err != nil {
			return err
		} else {
			c.sign = sign.NewCrypto(sh)
			return nil
		}
	}

	return errors.New("Invalid certificate type")

}

func (c *Client) GetCertName(name string) string {

	return name + "@" + c.Config.GetAppInfo().AppCode

}

func (c *Client) LoadUser() {

	users := c.Us.LoadAll(c.Config.GetAppInfo().AppCode)

	for i, _ := range users {

		user := users[i]
		user.MspId = c.Config.GetAppInfo().MspId

		err := keystore.LoadKey(user, c.Ks)

		if err == nil {
			c.Users[user.UserName] = user
		}
	}

}

func (c *Client) GetHeader() base.ReqHeader {
	return c.Config.GetReqHeader()
}

func (c *Client) GetURL(url string) string {
	return c.Config.GetNodeApi() + url
}

func (c *Client) GetUser(name string) (*msp.UserData, error) {
	user, ok := c.Users[name]
	if ok {
		return user, nil
	} else {
		return nil, errors.New("user does not exist")
	}

}

func (c *Client)LoadLocalUser(name string) (*msp.UserData, error)  {

	user :=&msp.UserData{
		UserName:name,
		AppCode:c.Config.GetAppInfo().AppCode,
	}

	err :=c.Us.Load(user)
	if err !=nil{
		return nil,err
	}
	err = keystore.LoadKey(user, c.Ks)

	if err != nil {
		return nil,err

	}else {
		c.Users[user.UserName] = user
		return user,nil

	}

}

func (c *Client) Sign(data string) string {
	mac, err := c.sign.Sign(data)

	if err != nil {
		logger.Error("Exception in signature")
	}

	return mac
}

func (c *Client) Verify(mac, data string) bool {
	return c.sign.Verify(mac, data)

}
