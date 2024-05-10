package verifiers

import (
	"context"
	"crypto/rsa"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
)

type SHA256WithRSADuoVerifier struct {
	pubicKeyVerifier SHA256WithRSAPubkeyVerifier
	certVerifier     SHA256WithRSAVerifier
}

func (v *SHA256WithRSADuoVerifier) Verify(ctx context.Context, serialNumber, message, signature string) error {
	if serialNumber == v.pubicKeyVerifier.keyID {
		return v.pubicKeyVerifier.Verify(ctx, serialNumber, message, signature)
	} else {
		return v.certVerifier.Verify(ctx, serialNumber, message, signature)
	}
}

func NewSHA256WithRSADuoVerifier(getter core.CertificateGetter, keyID string, publicKey rsa.PublicKey) *SHA256WithRSADuoVerifier {
	return &SHA256WithRSADuoVerifier{
		*NewSHA256WithRSAPubkeyVerifier(keyID, publicKey),
		*NewSHA256WithRSAVerifier(getter),
	}
}
