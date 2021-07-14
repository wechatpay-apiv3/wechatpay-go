// Copyright 2021 Tencent Inc. All rights reserved.

// Package auth 微信支付 API v3 Go SDK 安全验证相关接口
package auth

import "context"

// Verifier 数字签名验证器
type Verifier interface {
	Verify(ctx context.Context, serial, message, signature string) error // 对签名信息进行验证
}
