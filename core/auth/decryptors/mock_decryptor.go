package decryptors

import (
	"context"
	"fmt"
	"strings"
)

type MockDecryptor struct {
}

func (d *MockDecryptor) Decrypt(ctx context.Context, ciphertext string) (plaintext string, err error) {
	fmt.Printf("[MockDecryptor] Decrypting `%v`\n", ciphertext)
	if strings.HasPrefix(ciphertext, "Encrypted") {
		return strings.TrimPrefix(ciphertext, "Encrypted"), nil
	} else {
		return ciphertext, fmt.Errorf("cannot decrypt invalid cipher string:`%v`", ciphertext)
	}
}
