// Package credentials 微信支付 API v3 Go SDK 请求报文头 Authorization 信息生成器
package credentials

import (
	"context"
	"fmt"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

// WechatPayCredentials 微信支付请求报文头 Authorization 信息生成器
type WechatPayCredentials struct {
	Signer auth.Signer // 数字签名生成器
}

// GenerateAuthorizationHeader 生成请求报文头中的 Authorization 信息，详见：
// https://wechatpay-api.gitbook.io/wechatpay-api-v3/qian-ming-zhi-nan-1/qian-ming-sheng-cheng
func (c *WechatPayCredentials) GenerateAuthorizationHeader(ctx context.Context,
	method, canonicalURL, signBody string) (authorization string, err error) {
	if c.Signer == nil {
		return "", fmt.Errorf("you must init WechatPayCredentials with signer")
	}
	nonce, err := utils.GenerateNonce()
	if err != nil {
		return "", err
	}
	timestamp := time.Now().Unix()
	message := fmt.Sprintf(consts.SignatureMessageFormat, method, canonicalURL, timestamp, nonce, signBody)
	signatureResult, err := c.Signer.Sign(ctx, message)
	if err != nil {
		return "", err
	}
	authorization = fmt.Sprintf(consts.HeaderAuthorizationFormat, c.getAuthorizationType(),
		signatureResult.MchID, nonce, timestamp, signatureResult.CertificateSerialNo, signatureResult.Signature)
	return authorization, nil
}

func (c *WechatPayCredentials) getAuthorizationType() string {
	return "WECHATPAY2-" + c.Signer.Algorithm()
}
