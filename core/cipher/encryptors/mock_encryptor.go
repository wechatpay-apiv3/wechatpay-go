// Copyright 2021 Tencent Inc. All rights reserved.

package encryptors

import (
	"context"
	"fmt"
)

// MockEncryptor 模拟字符串加密器
type MockEncryptor struct {
	Serial string
}

// SelectCertificate 模拟选择加密用证书
func (e *MockEncryptor) SelectCertificate(ctx context.Context) (serial string, err error) {
	return e.Serial, nil
}

// Encrypt 使用指定证书进行字符串加密
func (e *MockEncryptor) Encrypt(ctx context.Context, serial, plaintext string) (ciphertext string, err error) {
	if serial != e.Serial {
		return plaintext, fmt.Errorf("invalid certificate serial: `%v`", serial)
	}

	ciphertext = "Encrypted" + plaintext
	return ciphertext, nil
}
