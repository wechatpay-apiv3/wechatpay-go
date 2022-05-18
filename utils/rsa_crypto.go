// Copyright 2021 Tencent Inc. All rights reserved.

package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"fmt"
)

// EncryptOAEPWithPublicKey 使用 OAEP padding方式用公钥进行加密
func EncryptOAEPWithPublicKey(message string, publicKey *rsa.PublicKey) (ciphertext string, err error) {
	if publicKey == nil {
		return "", fmt.Errorf("you should input *rsa.PublicKey")
	}
	ciphertextByte, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, publicKey, []byte(message), nil)
	if err != nil {
		return "", fmt.Errorf("encrypt message with public key err:%s", err.Error())
	}
	ciphertext = base64.StdEncoding.EncodeToString(ciphertextByte)
	return ciphertext, nil
}

// EncryptOAEPWithCertificate 先解析出证书中的公钥，然后使用 OAEP padding方式公钥进行加密
func EncryptOAEPWithCertificate(message string, certificate *x509.Certificate) (ciphertext string, err error) {
	if certificate == nil {
		return "", fmt.Errorf("you should input *x509.Certificate")
	}
	publicKey, ok := certificate.PublicKey.(*rsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("certificate is invalid")
	}
	return EncryptOAEPWithPublicKey(message, publicKey)
}

// EncryptPKCS1v15WithPublicKey 使用PKCS1 padding方式用公钥进行加密
func EncryptPKCS1v15WithPublicKey(message string, publicKey *rsa.PublicKey) (ciphertext string, err error) {
	if publicKey == nil {
		return "", fmt.Errorf("you should input *rsa.PublicKey")
	}
	ciphertextByte, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(message))
	if err != nil {
		return "", fmt.Errorf("encrypt message with public key err:%s", err.Error())
	}
	ciphertext = base64.StdEncoding.EncodeToString(ciphertextByte)
	return ciphertext, nil
}

// EncryptPKCS1v15WithCertificate 先解析出证书中的公钥，然后使用PKCS1 padding方式用公钥进行加密
func EncryptPKCS1v15WithCertificate(message string, certificate *x509.Certificate) (ciphertext string, err error) {
	if certificate == nil {
		return "", fmt.Errorf("you should input *x509.Certificate")
	}
	publicKey, ok := certificate.PublicKey.(*rsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("certificate is invalid")
	}
	return EncryptPKCS1v15WithPublicKey(message, publicKey)
}

// DecryptOAEP 使用私钥进行解密
func DecryptOAEP(ciphertext string, privateKey *rsa.PrivateKey) (message string, err error) {
	if privateKey == nil {
		return "", fmt.Errorf("you should input *rsa.PrivateKey")
	}
	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("base64 decode failed, error=%s", err.Error())
	}
	messageBytes, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, privateKey, decodedCiphertext, nil)
	if err != nil {
		return "", fmt.Errorf("decrypt ciphertext with private key err:%s", err)
	}
	return string(messageBytes), nil
}

// DecryptPKCS1v15 使用私钥对PKCS1 padding方式加密的字符串进行解密
func DecryptPKCS1v15(ciphertext string, privateKey *rsa.PrivateKey) (message string, err error) {
	if privateKey == nil {
		return "", fmt.Errorf("you should input *rsa.PrivateKey")
	}
	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("base64 decode failed, error=%s", err.Error())
	}
	messageBytes, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, decodedCiphertext)
	if err != nil {
		return "", fmt.Errorf("decrypt ciphertext with private key err:%s", err)
	}
	return string(messageBytes), nil
}
