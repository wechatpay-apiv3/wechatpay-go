// Copyright 2021 Tencent Inc. All rights reserved.

// Package validators 微信支付 API v3 Go SDK 应答报文验证器
package validators

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
)

// WechatPayResponseValidator 微信支付 API v3 默认应答报文验证器
type WechatPayResponseValidator struct {
	wechatPayValidator
}

// Validate 使用验证器对微信支付应答报文进行验证
func (v *WechatPayResponseValidator) Validate(ctx context.Context, response *http.Response) error {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("read response body err:[%s]", err.Error())
	}
	response.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return v.validateHTTPMessage(ctx, response.Header, body)
}

// NewWechatPayResponseValidator 使用 auth.Verifier 初始化一个 WechatPayResponseValidator
func NewWechatPayResponseValidator(verifier auth.Verifier) *WechatPayResponseValidator {
	return &WechatPayResponseValidator{
		wechatPayValidator{verifier: verifier},
	}
}
