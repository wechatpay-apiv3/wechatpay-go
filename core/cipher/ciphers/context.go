// Copyright 2021 Tencent Inc. All rights reserved.

package ciphers

import "context"

// contextKey WechatPayCipher Context Key Type
//
// 使用强类型避免与其他 Context Key 冲突
type contextKey string

// String contextKey 的字符串描述，区分普通字符串
func (c contextKey) String() string {
	return "WPCipherContext(" + string(c) + ")"
}

const (
	// 加密使用的微信支付平台证书序列号
	contextKeyEncryptSerial contextKey = "EncryptSerial"
)

// setEncryptSerial 往Context中写入用于加密的证书序列号，返回更新后的Context
func setEncryptSerial(ctx context.Context, serial string) context.Context {
	return context.WithValue(ctx, contextKeyEncryptSerial, serial)
}

// getEncryptSerial 从Context中读取用于加密的证书序列号
func getEncryptSerial(ctx context.Context) (string, bool) {
	serial, ok := ctx.Value(contextKeyEncryptSerial).(string)
	return serial, ok
}
