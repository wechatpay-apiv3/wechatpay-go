// Copyright 2021 Tencent Inc. All rights reserved.

// Package auth 微信支付 API v3 Go SDK 安全验证相关接口
package auth

import "context"

// Credential 请求报文头 Authorization 信息生成器
type Credential interface {
	GenerateAuthorizationHeader(ctx context.Context, method, canonicalURL, signBody string) (string, error)
}
