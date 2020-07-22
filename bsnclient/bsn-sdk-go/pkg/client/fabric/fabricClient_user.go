package fabric

import (
	userreq "bsn-sdk-go/pkg/core/entity/req/fabric/user"
	userres "bsn-sdk-go/pkg/core/entity/res/fabric/user"

	"bsn-sdk-go/pkg/common/errors"
	"bsn-sdk-go/pkg/core/cert"
	"bsn-sdk-go/pkg/core/entity/base"
	"bsn-sdk-go/pkg/core/entity/msp"
	"bsn-sdk-go/pkg/util/http"
	"bsn-sdk-go/pkg/util/keystore"

	"encoding/json"
	"fmt"

	"github.com/cloudflare/cfssl/csr"
	"github.com/wonderivan/logger"
)

func (c *FabricClient) RegisterUser(body userreq.RegisterReqDataBody) (*userres.RegisterResData, error) {

	url := c.GetURL("/api/fabric/v1/user/register")

	data := &userreq.RegisterReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())

	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}

	res := &userres.RegisterResData{}

	err = json.Unmarshal(resBytes, res)
	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}
	//验证返回签名
	if !c.Verify(res.Mac,res.GetEncryptionValue()){
		return nil,errors.New("签名验证失败")
	}

	return res, nil
}

func (c *FabricClient) EnrollUser(body userreq.RegisterReqDataBody) error {

	enrollBody := userreq.EnrollReqDataBody{
		Name:   body.Name,
		Secret: body.Secret,
	}

	cr := cert.GetCertificateRequest(c.GetCertName(enrollBody.Name))

	key, cspSigner, err := keystore.BCCSPKeyRequestGenerate(c.Ks)

	if err != nil {
		fmt.Println(err)
		return err
	}

	csrPEM, err := csr.Generate(cspSigner, cr)
	if err != nil {
		fmt.Println(err)
		return err
	}

	enrollBody.CsrPem = string(csrPEM)

	res, err := c.enroll(enrollBody)

	if err != nil {
		return err
	}

	if res.Header.Code == base.Header_success_code {

		pk := keystore.GetECDSAPrivateKey(key)
		user := &msp.UserData{
			UserName: enrollBody.Name,
			AppCode:  c.Config.GetAppInfo().AppCode,
			MspId:    c.Config.GetAppInfo().MspId,
			EnrollmentCertificate: []byte(res.Body.Cert),
			PrivateKey:            pk,
		}
		c.Us.Store(user)
		c.Users[user.UserName] = user
	} else {
		return errors.New(res.Header.Msg)
	}

	return nil

}

func (c *FabricClient) enroll(body userreq.EnrollReqDataBody) (*userres.EnrollResData, error) {

	url := c.GetURL("/api/fabric/v1/user/enroll")

	data := &userreq.EnrollReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())

	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}

	res := &userres.EnrollResData{}

	err = json.Unmarshal(resBytes, res)

	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}

	return res, nil

}
