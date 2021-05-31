package signers

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"testing"

	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
)

const (
	testSha256WithRSASignerPrivateKeyStr = `-----BEGIN PRIVATE KEY-----
-----END PRIVATE KEY-----`
	testSha256WithRSASignerMchCertificateSerialNo = ``
	testSha256WithRSASignerExpectSignature        = ``
)

var testSha256WithRSASignerPrivateKey *rsa.PrivateKey

func init() {
	block, _ := pem.Decode([]byte(testSha256WithRSASignerPrivateKeyStr))
	if block == nil {
		fmt.Printf("decode private key err")
		return
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		fmt.Printf("parse private key err:%s", err.Error())
	}
	var ok bool
	testSha256WithRSASignerPrivateKey, ok = key.(*rsa.PrivateKey)
	if !ok {
		fmt.Printf("%s is not rsa private key", testSha256WithRSASignerPrivateKeyStr)
	}
}

func TestSha256WithRSASigner_Sign(t *testing.T) {
	type args struct {
		ctx     context.Context
		message string
	}
	tests := []struct {
		name    string
		args    args
		want    *auth.SignatureResult
		wantErr bool
		mock    func()
	}{
		{
			name: "TestSha256WithRSASigner_Sign success",
			args: args{
				message: "source",
			},
			want: &auth.SignatureResult{
				CertificateSerialNo: testSha256WithRSASignerMchCertificateSerialNo,
				Signature:           testSha256WithRSASignerExpectSignature,
			},
			wantErr: false,
		},
		{
			name: "TestSha256WithRSASigner_Sign err should set privateKey",
			args: args{
				message: "source",
			},
			want:    &auth.SignatureResult{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SHA256WithRSASigner{
				CertificateSerialNo: testSha256WithRSASignerMchCertificateSerialNo,
				PrivateKey:          testSha256WithRSASignerPrivateKey,
			}
			got, err := s.Sign(tt.args.ctx, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				if got != tt.want {
					t.Errorf("sign() got nil but want not")
					return
				}
			} else {
				if got.Signature != tt.want.Signature || got.CertificateSerialNo != tt.want.Signature {
					t.Errorf("sign() got = %v, want %v", *got, *tt.want)
					return
				}
			}
		})
	}
}
