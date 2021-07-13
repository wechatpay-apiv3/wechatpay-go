// Copyright 2021 Tencent Inc. All rights reserved.

// Package validators 微信支付 API v3 Go SDK 应答报文签名验证器
package validators

import (
	"context"
	"net/http"
)

// NullValidator 空验证器，不对报文进行验证，对任意报文均不会返回错误，
// 在不需要对报文签名进行验证的情况（如微信支付账单文件下载）下使用
type NullValidator struct {
}

// Validate 跳过报文签名验证
func (validator *NullValidator) Validate(context.Context, *http.Response) error {
	return nil
}
