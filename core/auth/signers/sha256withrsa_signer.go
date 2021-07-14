// Copyright 2021 Tencent Inc. All rights reserved.

// Package signers 微信支付 API v3 Go SDK 数字签名生成器
package signers

import (
	"context"
	"crypto/rsa"
	"fmt"
	"strings"

	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

// SHA256WithRSASigner Sha256WithRSA 数字签名生成器
type SHA256WithRSASigner struct {
	MchID               string          // 商户号
	CertificateSerialNo string          // 商户证书序列号
	PrivateKey          *rsa.PrivateKey // 商户私钥
}

// Sign 对信息使用 SHA256WithRSA 算法进行签名
func (s *SHA256WithRSASigner) Sign(_ context.Context, message string) (*auth.SignatureResult, error) {
	if s.PrivateKey == nil {
		return nil, fmt.Errorf("you must set privatekey to use SHA256WithRSASigner")
	}
	if strings.TrimSpace(s.CertificateSerialNo) == "" {
		return nil, fmt.Errorf("you must set mch certificate serial no to use SHA256WithRSASigner")
	}
	signature, err := utils.SignSHA256WithRSA(message, s.PrivateKey)
	if err != nil {
		return nil, err
	}
	return &auth.SignatureResult{MchID: s.MchID, CertificateSerialNo: s.CertificateSerialNo, Signature: signature}, nil
}

// Algorithm 返回使用的签名算法：SHA256-RSA2048
func (s *SHA256WithRSASigner) Algorithm() string {
	return "SHA256-RSA2048"
}
