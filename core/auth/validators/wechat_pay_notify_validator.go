package validators

import (
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

func (v *WechatPayNotifyValidator) Validate(ctx context.Context, request *http.Request) error {
	reqBody, err := request.GetBody()
	if err != nil {
		return fmt.Errorf("get request body err: %v", err)
	}

	body, err := ioutil.ReadAll(reqBody)
	if err != nil {
		return fmt.Errorf("read request body err: %v", err)
	}

	return v.validateHTTPMessage(ctx, request.Header, body)
}

func NewWechatPayNotifyValidator(verifier auth.Verifier) *WechatPayNotifyValidator {
	return &WechatPayNotifyValidator{
		wechatPayValidator{verifier: verifier},
	}
}
