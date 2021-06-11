package utils

import (
	"crypto/rand"
)

const (
	NonceSymbols           = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // 随机字符串可用字符集
	NonceLength            = 32                                                               // 随机字符串的长度
)

func GenerateNonce() (string, error) {
	bytes := make([]byte, NonceLength)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	symbolsByteLength := byte(len(NonceSymbols))
	for i, b := range bytes {
		bytes[i] = NonceSymbols[b%symbolsByteLength]
	}
	return string(bytes), nil
}
