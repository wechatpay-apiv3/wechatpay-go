package builder

import (
	"crypto/rsa"

	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/credentials"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/signers"
)

func BuildMerchantCredential(
	mchID, certificateSerialNo string, privateKey *rsa.PrivateKey,
) *credentials.WechatPayCredentials {
	return &credentials.WechatPayCredentials{
		Signer: &signers.SHA256WithRSASigner{PrivateKey: privateKey, CertificateSerialNo: certificateSerialNo},
		MchID:  mchID,
	}
}
