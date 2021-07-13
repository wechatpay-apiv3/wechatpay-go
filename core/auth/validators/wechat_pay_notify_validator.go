// Copyright 2021 Tencent Inc. All rights reserved.

package validators

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
)

// WechatPayNotifyValidator 微信支付 API v3 通知请求报文验证器
type WechatPayNotifyValidator struct {
	wechatPayValidator
}

// Validate 对接收到的微信支付 API v3 通知请求报文进行验证
func (v *WechatPayNotifyValidator) Validate(ctx context.Context, request *http.Request) error {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return fmt.Errorf("read request body err: %v", err)
	}

	_ = request.Body.Close()
	request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return v.validateHTTPMessage(ctx, request.Header, body)
}

// NewWechatPayNotifyValidator 使用 auth.Verifier 初始化一个 WechatPayNotifyValidator
func NewWechatPayNotifyValidator(verifier auth.Verifier) *WechatPayNotifyValidator {
	return &WechatPayNotifyValidator{
		wechatPayValidator{verifier: verifier},
	}
}
