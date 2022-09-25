// Copyright 2021 Tencent Inc. All rights reserved.

// Package notify 微信支付 API v3 Go SDK 商户通知处理库
package notify

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/validators"
	"io"
	"net/http"
)

// Handler 通知处理器，使用前先设置验签和解密的算法套件
type Handler struct {
	cipherSuites map[string]CipherSuite
}

// CipherSuite 算法套件，包括验签和解密
type CipherSuite struct {
	signatureType string
	validator     validators.WechatPayNotifyValidator
	aeadAlgorithm string
	aead          cipher.AEAD
}

// NewEmptyHandler 创建一个不包含算法套件的空通知处理器
func NewEmptyHandler() *Handler {
	h := &Handler{
		cipherSuites: map[string]CipherSuite{},
	}

	return h
}

// AddCipherSuite 添加一个算法套件
func (h *Handler) AddCipherSuite(cipherSuite CipherSuite) *Handler {
	h.cipherSuites[cipherSuite.signatureType] = cipherSuite
	return h
}

// AddRSAWithAESGCM 添加一个 RSA + AES-GCM 的算法套件
func (h *Handler) AddRSAWithAESGCM(verifier auth.Verifier, aesgcm cipher.AEAD) *Handler {
	v := CipherSuite{
		signatureType: "WECHATPAY2-RSA2048-SHA256",
		validator:     *validators.NewWechatPayNotifyValidator(verifier),
		aeadAlgorithm: "AEAD_AES_256_GCM",
		aead:          aesgcm,
	}
	return h.AddCipherSuite(v)
}

// ParseNotifyRequest 从 HTTP 请求(http.Request) 中解析 微信支付通知(notify.Request)
func (h *Handler) ParseNotifyRequest(
	ctx context.Context,
	request *http.Request,
	content interface{},
) (*Request, error) {
	signType := request.Header.Get("Wechatpay-Signature-Type")
	if signType == "" {
		signType = "WECHATPAY2-RSA2048-SHA256"
	}

	suite, ok := h.cipherSuites[signType]
	if !ok {
		return nil, fmt.Errorf("invalid Wechatpay-Signature-Type: %s", signType)
	}

	if err := suite.validator.Validate(ctx, request); err != nil {
		return nil, fmt.Errorf("not valid wechatpay notify request: %v", err)
	}

	body, err := getRequestBody(request)
	if err != nil {
		return nil, err
	}

	return processBody(suite, body, content)
}

func processBody(suite CipherSuite, body []byte, content interface{}) (*Request, error) {
	ret := new(Request)
	if err := json.Unmarshal(body, ret); err != nil {
		return nil, fmt.Errorf("parse request body error: %v", err)
	}

	if ret.Resource.Algorithm != suite.aeadAlgorithm {
		return nil, fmt.Errorf("invalid resource.algorithm: %s, suite.algorithm: %s",
			ret.Resource.Algorithm,
			suite.aeadAlgorithm)
	}

	plaintext, err := doAEADOpen(
		suite.aead,
		ret.Resource.Nonce,
		ret.Resource.Ciphertext,
		ret.Resource.AssociatedData,
	)
	if err != nil {
		return ret, fmt.Errorf("%s decrypt error: %v", ret.Resource.Algorithm, err)
	}

	ret.Resource.Plaintext = plaintext

	if err = json.Unmarshal([]byte(plaintext), &content); err != nil {
		return ret, fmt.Errorf("unmarshal plaintext to content failed: %v", err)
	}

	return ret, nil
}

func doAEADOpen(c cipher.AEAD, nonce, ciphertext, additionalData string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	plaintext, err := c.Open(
		nil,
		[]byte(nonce),
		data,
		[]byte(additionalData),
	)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func getRequestBody(request *http.Request) ([]byte, error) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, fmt.Errorf("read request body err: %v", err)
	}

	_ = request.Body.Close()
	request.Body = io.NopCloser(bytes.NewBuffer(body))

	return body, nil
}

// NewRSANotifyHandler 创建一个 RSA 的通知处理器，它包含 AES-GCM 解密能力
func NewRSANotifyHandler(apiV3Key string, verifier auth.Verifier) (*Handler, error) {
	c, err := aes.NewCipher([]byte(apiV3Key))
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	return NewEmptyHandler().AddRSAWithAESGCM(verifier, aesgcm), nil
}

// NewNotifyHandler 创建通知处理器
// Deprecated: should NewRSANotifyHandler use instead
func NewNotifyHandler(apiV3Key string, verifier auth.Verifier) *Handler {
	h, _ := NewRSANotifyHandler(apiV3Key, verifier)
	return h
}
