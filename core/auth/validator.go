//Package auth 微信支付api v3 go http-client 签名校验相关接口
package auth

import (
	"context"
	"net/http"
)

// Validator 回包校验器
type Validator interface {
	Validate(ctx context.Context, response *http.Response) error // 对http回包进行校验
}
