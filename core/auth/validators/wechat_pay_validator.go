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

// WechatPayValidator 微信支付 API v3 默认应答报文验证器
type WechatPayValidator struct {
	baseWechatPayValidator
}

// Validate 使用验证器对微信支付应答报文进行验证
func (v *WechatPayValidator) Validate(ctx context.Context, response *http.Response) error {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("read response body err:[%s]", err.Error())
	}
	response.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return v.validateHTTPMessage(ctx, response.Header, body)
}

func NewWechatPayValidator(verifier auth.Verifier) *WechatPayValidator {
	return &WechatPayValidator{
		baseWechatPayValidator{verifier: verifier},
	}
}
