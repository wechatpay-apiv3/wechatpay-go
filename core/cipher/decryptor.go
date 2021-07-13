// Copyright 2021 Tencent Inc. All rights reserved.

package cipher

import "context"

// Decryptor 字符串解密器
type Decryptor interface {
	// Decrypt 对字符串解密
	Decrypt(ctx context.Context, ciphertext string) (plaintext string, err error)
}
