package core

import (
	"fmt"
	"net/http"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
)

// dialSettings 微信支付apiv3 go http-client需要的配置信息
type dialSettings struct {
	HTTPClient *http.Client    // 自定义所使用的 HTTPClient 实例
	Header     http.Header     // 自定义额外请求头
	Credential auth.Credential // 请求头 Authorization 生成器
	Validator  auth.Validator  // 应答包签名校验器
	Timeout    time.Duration   // HTTP 请求超时时间，将覆盖 HTTPClient 中的 Timeout（如果你同步设置了 HTTPClient）
}

// Validate 校验请求配置是否有效
func (ds *dialSettings) Validate() error {
	if ds.Credential == nil {
		return fmt.Errorf("credential is required for Client")
	}
	if ds.Validator == nil {
		return fmt.Errorf("validator is required for Client")
	}
	return nil
}
