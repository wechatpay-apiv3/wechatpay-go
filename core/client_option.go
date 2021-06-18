package core

import (
	"net/http"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cipher"
)

// ClientOption 微信支付 API v3 HTTPClient core.Client 初始化参数
type ClientOption interface {
	Apply(settings *DialSettings) error
}

// ErrorOption 错误初始化参数，用于返回错误
type ErrorOption struct{ Error error }

// Apply 返回初始化错误
func (w ErrorOption) Apply(o *DialSettings) error {
	return w.Error
}

// WithSignerOption 为 Client 设置 Signer
type WithSignerOption struct {
	Signer auth.Signer
}

// Apply 将配置添加到 DialSettings 中
func (w WithSignerOption) Apply(o *DialSettings) error {
	o.Signer = w.Signer
	return nil
}

// WithValidatorOption 为 Client 设置 Validator
type WithValidatorOption struct {
	Validator auth.Validator
}

// Apply 将配置添加到 DialSettings 中
func (w WithValidatorOption) Apply(o *DialSettings) error {
	o.Validator = w.Validator
	return nil
}

// WithHTTPClientOption 为 Client 设置 HTTPClient
type WithHTTPClientOption struct {
	Client *http.Client
}

// Apply 将配置添加到 DialSettings 中
func (w WithHTTPClientOption) Apply(o *DialSettings) error {
	o.HTTPClient = w.Client
	return nil
}

// WithTimeoutOption 为 Client 设置 超时时间
type WithTimeoutOption time.Duration

// Apply 将配置添加到DialSettings中
func (w WithTimeoutOption) Apply(o *DialSettings) error {
	o.Timeout = time.Duration(w)
	return nil
}

// WithHeaderOption 为 Client 设置 DefaultHeader
type WithHeaderOption struct {
	Header http.Header
}

// Apply 将配置添加到 DialSettings 中
func (w WithHeaderOption) Apply(o *DialSettings) error {
	o.Header = w.Header
	return nil
}

// WithCipherOption 为 Client 设置 Cipher
type WithCipherOption struct {
	Cipher cipher.Cipher
}

// Apply 将配置添加到 DialSettings 中
func (w WithCipherOption) Apply(o *DialSettings) error {
	o.Cipher = w.Cipher
	return nil
}
