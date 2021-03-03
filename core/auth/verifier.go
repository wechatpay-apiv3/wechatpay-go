//Package auth 微信支付api v3 go http-client 签名校验相关接口
package auth

import "context"

// Verifier 验证器
type Verifier interface {
	Verify(ctx context.Context, serial, message, signature string) error // 对签名信息进行验证
}
