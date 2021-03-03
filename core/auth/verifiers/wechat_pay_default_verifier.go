//Package verifiers 微信支付api v3 go http-client 验证器
package verifiers

import (
	"context"
)

//WechatPayDefaultVerifier 默认验证器
type WechatPayDefaultVerifier struct {
}

// Verify 不对信息进行验证，仅下载证书的时候使用
func (verifier *WechatPayDefaultVerifier) Verify(ctx context.Context, serialNumber, message, signature string) error {
	return nil
}
