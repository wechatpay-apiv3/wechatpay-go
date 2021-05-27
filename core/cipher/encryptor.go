package cipher

import "context"

// Encryptor 字符串加密器
type Encryptor interface {
	// SelectCertificate 选择合适的微信支付平台证书用于加密
	SelectCertificate(ctx context.Context) (serial string, err error)

	// Encrypt 对字符串加密
	Encrypt(ctx context.Context, serial, plaintext string) (ciphertext string, err error)
}
