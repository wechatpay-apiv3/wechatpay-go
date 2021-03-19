//Package validators 微信支付api v3 go http-client 回包校验器
package validators

import (
	"context"
	"net/http"
)

// WechatPayNullValidator 回包校验器
type WechatPayNullValidator struct {
}

// Validate 使用验证器对回包进行校验
func (validator *WechatPayNullValidator) Validate(ctx context.Context, response *http.Response) error {
	return nil
}