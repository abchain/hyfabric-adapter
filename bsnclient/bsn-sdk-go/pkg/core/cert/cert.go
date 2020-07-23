package cert

import (
	"hyperledger.abchain.org/adapter/hyfabric/bsnclient/bsn-sdk-go/pkg/util/keystore"
	"hyperledger.abchain.org/adapter/hyfabric/bsnclient/bsn-sdk-go/third_party/github.com/hyperledger/fabric/bccsp"
	"encoding/hex"
	"fmt"
	"github.com/cloudflare/cfssl/csr"
)

//GetCSRPEM ...
func GetCSRPEM(name string, ks bccsp.KeyStore) (string, error) {

	cr := GetCertificateRequest(name)

	key, cspSigner, err := keystore.BCCSPKeyRequestGenerate(ks)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	csrPEM, err := csr.Generate(cspSigner, cr)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Println("key:", hex.EncodeToString(key.SKI()))

	fmt.Println("csrPEM：", string(csrPEM))

	return string(csrPEM), nil
}

//GetCertificateRequest 创建证书请求文件信息
func GetCertificateRequest(name string) *csr.CertificateRequest {

	cr := &csr.CertificateRequest{}
	cr.CN = name
	cr.KeyRequest = newCfsslBasicKeyRequest()

	return cr

}

//证书算法
func newCfsslBasicKeyRequest() *csr.KeyRequest {
	return &csr.KeyRequest{A: "ecdsa", S: 256}
}
