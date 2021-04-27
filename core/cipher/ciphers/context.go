package ciphers

import "context"

type contextKey string

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
