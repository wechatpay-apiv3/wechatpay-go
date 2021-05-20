package decryptors

import (
	"context"
	"crypto/rsa"

	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

type WechatPayDecryptor struct {
	// 商户私钥
	PrivateKey *rsa.PrivateKey
}

func (d *WechatPayDecryptor) Decrypt(ctx context.Context, ciphertext string) (plaintext string, err error) {
	if ciphertext == "" {
		return "", nil
	}
	return utils.DecryptOAEP(ciphertext, d.PrivateKey)
}
