package keystore

import (
	"hyperledger.abchain.org/adapter/hyfabric/bsnclient/bsn-sdk-go/pkg/common/errors"
	ecdsaUtil "hyperledger.abchain.org/adapter/hyfabric/bsnclient/bsn-sdk-go/pkg/util/esdsa"
	"hyperledger.abchain.org/adapter/hyfabric/bsnclient/bsn-sdk-go/third_party/github.com/hyperledger/fabric/bccsp"
	"crypto"
	"crypto/x509"
	"io"
)

type ecdsaSigner struct{}

func (s *ecdsaSigner) Sign(k bccsp.Key, digest []byte, opts bccsp.SignerOpts) ([]byte, error) {
	return ecdsaUtil.SignECDSA(k.(*ecdsaPrivateKey).privKey, digest)
}

type bccspCryptoSigner struct {
	//csp core.CryptoSuite
	key bccsp.Key
	pk  interface{}
}

func (s *bccspCryptoSigner) Public() crypto.PublicKey {
	return s.pk
}

func (s *bccspCryptoSigner) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error) {
	ecdsa := ecdsaSigner{}
	return ecdsa.Sign(s.key, digest, opts)
}

func New(key bccsp.Key) (crypto.Signer, error) {
	// Validate arguments
	if key == nil {
		return nil, errors.New("key must be different from nil.")
	}
	if key.Symmetric() {
		return nil, errors.New("key must be asymmetric.")
	}

	// Marshall the bccsp public key as a crypto.PublicKey
	pub, err := key.PublicKey()
	if err != nil {
		return nil, errors.New("failed getting public key")
	}

	raw, err := pub.Bytes()
	if err != nil {
		return nil, errors.New("failed marshalling public key")
	}

	pk, err := DERToPublicKey(raw)
	if err != nil {
		return nil, errors.New("failed marshalling der to public key")
	}

	return &bccspCryptoSigner{key, pk}, nil
}

func DERToPublicKey(raw []byte) (pub interface{}, err error) {
	if len(raw) == 0 {
		return nil, errors.New("Invalid DER. It must be different from nil.")
	}

	key, err := x509.ParsePKIXPublicKey(raw)

	return key, err
}
