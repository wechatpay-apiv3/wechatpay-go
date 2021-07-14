// Copyright 2021 Tencent Inc. All rights reserved.

// Package auth 微信支付 API v3 Go SDK 安全验证相关接口
package auth

import (
	"context"
	"net/http"
)

// Validator 应答报文验证器
type Validator interface {
	Validate(ctx context.Context, response *http.Response) error // 对 HTTP 应答报文进行验证
}
