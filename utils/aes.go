package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

//  DecryptToString 将下载证书的回包解析成证书
//
//  解析后的证书是这样的
//  -----BEGIN CERTIFICATE-----
//	-----END CERTIFICATE-----
func DecryptToString(apiv3Key, associatedData, nonce, ciphertext string) (certificate string, err error) {
	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	c, err := aes.NewCipher([]byte(apiv3Key))
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}
	certificateByte, err := gcm.Open(nil, []byte(nonce), decodedCiphertext, []byte(associatedData))
	if err != nil {
		return "", err
	}
	certificate = string(certificateByte)
	return certificate, nil
}
