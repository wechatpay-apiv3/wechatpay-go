//Package auth 微信支付api v3 go http-client 签名校验相关接口
package auth

import "context"

// Credential Authorization信息生成器
type Credential interface {
	GenerateAuthorizationHeader(ctx context.Context, method, canonicalURL,
		signBody string) (authorization string, err error)
}
