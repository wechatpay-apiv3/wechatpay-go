// Copyright 2024 Tencent Inc. All rights reserved.

// Package verifiers 微信支付 API v3 Go SDK 数字签名验证器
package verifiers

import (
	"context"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// SHA256WithRSAPubkeyVerifier 数字签名验证器,使用微信支付提供的公钥验证签名
type SHA256WithRSAPubkeyVerifier struct {
	keyID     string
	publicKey rsa.PublicKey
}

// Verify 使用微信支付提供的公钥验证签名
func (v *SHA256WithRSAPubkeyVerifier) Verify(ctx context.Context, serialNumber, message, signature string) error {
	if ctx == nil {
		return fmt.Errorf("verify failed: context is nil")
	}
	if v.keyID != serialNumber {
		return fmt.Errorf("verify failed: key-id[%s] does not match serial number[%s]", v.keyID, serialNumber)
	}

	sigBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return fmt.Errorf("verify failed: signature is not base64 encoded")
	}
	hashed := sha256.Sum256([]byte(message))
	err = rsa.VerifyPKCS1v15(&v.publicKey, crypto.SHA256, hashed[:], sigBytes)
	if err != nil {
		return fmt.Errorf("verify signature with public key error:%s", err.Error())
	}
	return nil
}

// GetSerial 获取可验签的公钥序列号
func (v *SHA256WithRSAPubkeyVerifier) GetSerial(ctx context.Context) (string, error) {
	return v.keyID, nil
}

// NewSHA256WithRSAPubkeyVerifier 使用 rsa.PublicKey 初始化验签器
func NewSHA256WithRSAPubkeyVerifier(keyID string, publicKey rsa.PublicKey) *SHA256WithRSAPubkeyVerifier {
	return &SHA256WithRSAPubkeyVerifier{keyID: keyID, publicKey: publicKey}
}
