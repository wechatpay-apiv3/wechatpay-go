// Copyright 2021 Tencent Inc. All rights reserved.

package core

import (
	"fmt"
	"net/http"

	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cipher"
)

// DialSettings 微信支付 API v3 Go SDK core.Client 需要的配置信息
type DialSettings struct {
	HTTPClient *http.Client   // 自定义所使用的 HTTPClient 实例
	Signer     auth.Signer    // 签名器
	Validator  auth.Validator // 应答包签名校验器
	Cipher     cipher.Cipher  // 敏感字段加解密套件
}

// Validate 校验请求配置是否有效
func (ds *DialSettings) Validate() error {
	if ds.Validator == nil {
		return fmt.Errorf("validator is required for Client")
	}
	if ds.Signer == nil {
		return fmt.Errorf("signer is required for Client")
	}
	return nil
}
