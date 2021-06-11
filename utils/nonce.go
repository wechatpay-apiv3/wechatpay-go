package utils

import (
	"crypto/rand"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
)

func GenerateNonce() (string, error) {
	bytes := make([]byte, consts.NonceLength)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	symbolsByteLength := byte(len(consts.NonceSymbols))
	for i, b := range bytes {
		bytes[i] = consts.NonceSymbols[b%symbolsByteLength]
	}
	return string(bytes), nil
}
