//nolint
package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testRSACryptoUtilPublicKeyStr = `-----BEGIN PUBLIC KEY-----
-----END PUBLIC KEY-----`
	testRSACryptoUtilPrivateKeyStr = `-----BEGIN PRIVATE KEY-----
-----END PRIVATE KEY-----`
	testRSACryptoUtilMchCertificateStr = `-----BEGIN CERTIFICATE-----
-----END CERTIFICATE-----`
)

var (
	testRSACryptoUtilPrivateKey  *rsa.PrivateKey
	testRSACryptoUtilCertificate *x509.Certificate
	testRSACryptoUtilPublicKey   *rsa.PublicKey
	err                          error
)

func init() {
	testRSACryptoUtilPrivateKey, err = LoadPrivateKey(testRSACryptoUtilPublicKeyStr)
	if err != nil {
		fmt.Printf("load private key err:%s", err)
	}
	testRSACryptoUtilPublicKey, err = LoadPublicKey(testRSACryptoUtilPrivateKeyStr)
	if err != nil {
		fmt.Printf("load public key err:%s", err)
	}
	testRSACryptoUtilCertificate, _ = LoadCertificate(testRSACryptoUtilMchCertificateStr)
	if err != nil {
		fmt.Printf("load certificate err:%s", err)
	}
}

func TestEncryptAndDecrypt(t *testing.T) {
	message := "hello world"
	// 证书中的公钥加密
	cipertext, err := EncryptOAEPWithCertificate(message, testRSACryptoUtilCertificate)
	assert.Equal(t, nil, err)
	// 私钥解密
	decryptMessage, err := DecryptOAEP(cipertext, testRSACryptoUtilPrivateKey)
	assert.Equal(t, nil, err)
	assert.Equal(t, decryptMessage, message)
	// 直接公钥加密
	cipertext, err = EncryptOAEPWithPublicKey(message, testRSACryptoUtilPublicKey)
	assert.Equal(t, nil, err)
	// 私钥解密
	decryptMessage, err = DecryptOAEP(cipertext, testRSACryptoUtilPrivateKey)
	assert.Equal(t, nil, err)
	assert.Equal(t, decryptMessage, message)
}
