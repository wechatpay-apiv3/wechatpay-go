package decryptors

import (
	"context"
	"fmt"
	"strings"
)

// MockDecryptor 模拟字符串解密器
type MockDecryptor struct {
}

// Decrypt 对字符串进行模拟解密
func (d *MockDecryptor) Decrypt(ctx context.Context, ciphertext string) (plaintext string, err error) {
	fmt.Printf("[MockDecryptor] Decrypting `%v`\n", ciphertext)
	if strings.HasPrefix(ciphertext, "Encrypted") {
		return strings.TrimPrefix(ciphertext, "Encrypted"), nil
	} else {
		return ciphertext, fmt.Errorf("cannot decrypt invalid cipher string:`%v`", ciphertext)
	}
}
