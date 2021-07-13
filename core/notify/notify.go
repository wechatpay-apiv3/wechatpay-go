// Copyright 2021 Tencent Inc. All rights reserved.

// Package notify 微信支付 API v3 Go SDK 商户通知处理库
package notify

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/validators"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

// Handler 微信支付通知 Handler
type Handler struct {
	mchAPIv3Key string
	validator   validators.WechatPayNotifyValidator
}

// ParseNotifyRequest 从 HTTP 请求(http.Request) 中解析 微信支付通知(notify.Request)
func (h *Handler) ParseNotifyRequest(ctx context.Context, request *http.Request, content interface{}) (
	*Request, error,
) {
	if err := h.validator.Validate(ctx, request); err != nil {
		return nil, fmt.Errorf("not valid wechatpay notify request: %v", err)
	}

	body, err := getRequestBody(request)
	if err != nil {
		return nil, err
	}

	ret := new(Request)
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, fmt.Errorf("parse request body error: %v", err)
	}

	plaintext, err := utils.DecryptAES256GCM(
		h.mchAPIv3Key, ret.Resource.AssociatedData, ret.Resource.Nonce, ret.Resource.Ciphertext,
	)
	if err != nil {
		return ret, fmt.Errorf("decrypt request error: %v", err)
	}

	ret.Resource.Plaintext = plaintext

	if err = json.Unmarshal([]byte(plaintext), &content); err != nil {
		return ret, fmt.Errorf("unmarshal plaintext to content failed: %v", err)
	}

	return ret, nil
}

func getRequestBody(request *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, fmt.Errorf("read request body err: %v", err)
	}

	_ = request.Body.Close()
	request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return body, nil
}

// NewNotifyHandler 创建通知处理器
func NewNotifyHandler(mchAPIv3Key string, verifier auth.Verifier) *Handler {
	return &Handler{
		mchAPIv3Key: mchAPIv3Key,
		validator:   *validators.NewWechatPayNotifyValidator(verifier),
	}
}
