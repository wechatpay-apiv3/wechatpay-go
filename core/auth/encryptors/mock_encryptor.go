package encryptors

import (
	"context"
	"fmt"
)

type MockEncryptor struct {
	Serial string
}

func (e *MockEncryptor) SelectCertificate(ctx context.Context) (serial string, err error) {
	return e.Serial, nil
}

func (e *MockEncryptor) Encrypt(ctx context.Context, serial, plaintext string) (ciphertext string, err error) {
	if serial != e.Serial {
		return plaintext, fmt.Errorf("invalid certificate serial: `%v`", serial)
	}

	ciphertext = "Encrypted" + plaintext
	return ciphertext, nil
}
