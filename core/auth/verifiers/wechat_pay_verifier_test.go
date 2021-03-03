package verifiers

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"testing"

	"github.com/wechatpay-apiv3/wechatpay-go/utils"

	"github.com/agiledragon/gomonkey"
)

const (
	testWechatPayVerifierPlatformSerialNumber = ""
	testWechatPayVerifierPlatformCertificate  = `-----BEGIN CERTIFICATE-----
-----END CERTIFICATE-----`
)

var (
	certificate *x509.Certificate
)

func init() {
	certificate, _ = utils.LoadCertificate(testWechatPayVerifierPlatformCertificate)
}

func TestWechatPayVerifier_Verify(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()
	verifyPKCS1v15Outputs := []gomonkey.OutputCell{
		{Values: gomonkey.Params{nil}, Times: 1},
		{Values: gomonkey.Params{fmt.Errorf("verify failed")}, Times: 1},
	}
	patches.ApplyFuncSeq(rsa.VerifyPKCS1v15, verifyPKCS1v15Outputs)
	type fields struct {
		Certificates map[string]*x509.Certificate
	}
	type args struct {
		ctx          context.Context
		serialNumber string
		message      string
		signature    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "verify success",
			fields: fields{Certificates: map[string]*x509.Certificate{testWechatPayVerifierPlatformSerialNumber: certificate}},
			args: args{
				ctx:          context.Background(),
				serialNumber: testWechatPayVerifierPlatformSerialNumber,
				signature:    "1",
				message:      "1",
			},
			wantErr: false,
		},
		{
			name:   "verify failed",
			fields: fields{Certificates: map[string]*x509.Certificate{testWechatPayVerifierPlatformSerialNumber: certificate}},
			args: args{
				ctx:          context.Background(),
				serialNumber: testWechatPayVerifierPlatformSerialNumber,
				signature:    "",
				message:      "",
			},
			wantErr: true,
		},
		{
			name:   "verify failed with no corresponding certificate",
			fields: fields{Certificates: map[string]*x509.Certificate{}},
			args: args{
				ctx:          context.Background(),
				serialNumber: testWechatPayVerifierPlatformSerialNumber,
				signature:    "1",
				message:      "1",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			verifier := &WechatPayVerifier{
				Certificates: tt.fields.Certificates,
			}
			if err := verifier.Verify(tt.args.ctx, tt.args.serialNumber, tt.args.message,
				tt.args.signature); (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
