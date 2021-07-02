package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testAESUtilAPIV3Key       = "testAPIv3Key0000"
	testAESUtilCiphertext     = "FsdXzxryWfKwvLJKf8LG/ToRPTRh8RN9wROC"
	testAESUtilPlaintext      = "Hello World"
	testAESUtilNonce          = "wCq51qZv4Yfg"
	testAESUtilAssociatedData = "Dl56FrqWQJF1t9LtC3vEUsniZKvbqdR8"
)

func TestDecryptAes256Gcm(t *testing.T) {
	type args struct {
		apiv3Key       string
		associatedData string
		nonce          string
		ciphertext     string
	}
	tests := []struct {
		name      string
		args      args
		plaintext string
		wantErr   bool
	}{
		{
			name: "decrypt success",
			args: args{
				apiv3Key:       testAESUtilAPIV3Key,
				associatedData: testAESUtilAssociatedData,
				nonce:          testAESUtilNonce,
				ciphertext:     testAESUtilCiphertext,
			},
			wantErr:   false,
			plaintext: testAESUtilPlaintext,
		},
		{
			name: "invalid base64 ciphertext",
			args: args{
				apiv3Key:       testAESUtilAPIV3Key,
				associatedData: testAESUtilAssociatedData,
				nonce:          testAESUtilNonce,
				ciphertext:     "invalid cipher",
			},
			wantErr: true,
		},
		{
			name: "invalid ciphertext",
			args: args{
				apiv3Key:       testAESUtilAPIV3Key,
				associatedData: testAESUtilAssociatedData,
				nonce:          testAESUtilNonce,
				ciphertext:     "SGVsbG8gV29ybGQK",
			},
			wantErr: true,
		},
		{
			name: "invalid aes key",
			args: args{
				apiv3Key:       "not a aes key",
				associatedData: testAESUtilAssociatedData,
				nonce:          testAESUtilNonce,
				ciphertext:     testAESUtilCiphertext,
			},
			wantErr: true,
		},
		{
			name: "wrong aes key",
			args: args{
				apiv3Key:       "testAPIv3Key1111",
				associatedData: testAESUtilAssociatedData,
				nonce:          testAESUtilNonce,
				ciphertext:     testAESUtilCiphertext,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plaintext, err := DecryptAES256GCM(tt.args.apiv3Key, tt.args.associatedData, tt.args.nonce, tt.args.ciphertext)
			require.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.plaintext, plaintext)
		})
	}
}
