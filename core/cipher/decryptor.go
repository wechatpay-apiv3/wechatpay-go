package cipher

import "context"

type Decryptor interface {
	// Decrypt 对字符串解密
	Decrypt(ctx context.Context, ciphertext string) (plaintext string, err error)
}
