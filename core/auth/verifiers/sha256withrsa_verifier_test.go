// Copyright 2021 Tencent Inc. All rights reserved.

package verifiers

import (
	"context"
	"crypto/x509"
	"testing"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

const (
	testWechatPayVerifierPlatformSerialNumber = "F5765756002FDD77"
	testWechatPayVerifierPlatformCertificate  = `-----BEGIN CERTIFICATE-----
MIIDVzCCAj+gAwIBAgIJAPV2V1YAL913MA0GCSqGSIb3DQEBCwUAMEIxCzAJBgNV
BAYTAlhYMRUwEwYDVQQHDAxEZWZhdWx0IENpdHkxHDAaBgNVBAoME0RlZmF1bHQg
Q29tcGFueSBMdGQwHhcNMjEwNDI3MDg0MDMyWhcNMzEwNDI1MDg0MDMyWjBCMQsw
CQYDVQQGEwJYWDEVMBMGA1UEBwwMRGVmYXVsdCBDaXR5MRwwGgYDVQQKDBNEZWZh
dWx0IENvbXBhbnkgTHRkMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA
2VCTd91fnUn73Xy9DLvt/V62TVxRTEEstVdeRaZ3B3leO0pldE806mXO4RwdHXag
HQ4vGeZN0yqm++rDsGK+U3AH7kejyD2pXshNP9Cq5YwbptiLGtjcquw4HNxJQUOm
DeJf2vg6byms9RUipiq4SzbJKqJFlUpbuIPDpSpWz10PYmyCNeDGUUK65E5h2B83
4uxl1zNLYQCrkdBzb8oUxwYeP5a2DNxmjL5lsJML7DGr5znsevnoqGRwTm9fxCGf
y8wus7hwKz6clt3Whmmda7UAdb1c08hEQFVRbF14AR73xbnd8N0obCWJPCbzMCtk
aSef4FdEEgEXJiw0VAJT8wIDAQABo1AwTjAdBgNVHQ4EFgQUT1c7nd/SUO76HSoZ
umNUJv1R5PwwHwYDVR0jBBgwFoAUT1c7nd/SUO76HSoZumNUJv1R5PwwDAYDVR0T
BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAM+tslqBxYwqL9fdvGG6hfy69sjfX
UhBtBLWYugKKQCOWWLeq5dDWm3i5Cx2Rgiy9uc7RfmJNxQfIKlcoCNP85BjDoG1B
YnVc6znlcrT9uHgseha3987WwZsFAQbcy8TLUYHzVB8gmDgq8O08xdIe0eczatI8
t3Rg8WXO6Gs66JJ4JR+rD01o3FiSOQCRWhn19NSyDydsgPlOR2t9B9L+MkJwlsMG
Krn85TnwL3qcInzRnU8X86faXXJrI0IJi44tECKw8ftngCl6vyNwNNKPDwdkcuuV
8y3iBixO5IuKxEKEp2wGPV/4W1AXO73Z3Gb7z/1oxdgeO0hVqz1hBasTCQ==
-----END CERTIFICATE-----`
	testExpectedSignature = "BKyAfU4iMCuvXMXS0Wzam3V/cnxZ+JaqigPM5OhljS2iOT95OO6Fsuml2JkFANJU9" +
		"K6q9bLlDhPXuoVz+pp4hAm6pHU4ld815U4jsKu1RkyaII+1CYBUYC8TK0XtJ8FwUXXz8vZHh58rrAVN1XwNyv" +
		"D1vfpxrMT4SL536GLwvpUHlCqIMzoZUguLli/K8V29QiOhuH6IEqLNJn8e9b3nwNcQ7be3CzYGpDAKBfDGPCq" +
		"Cv8Rw5zndhlffk2FEA70G4hvMwe51qMN/RAJbknXG23bSlObuTCN7Ndj1aJGH6/L+hdwfLpUtJm4QYVazzW7D" +
		"FD27EpSQEqA8bX9+8m1rLg=="
)

var (
	certificate *x509.Certificate
)

func init() {
	certificate, _ = utils.LoadCertificate(testWechatPayVerifierPlatformCertificate)
}

func TestWechatPayVerifier_Verify(t *testing.T) {
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
			name: "verify success",
			fields: fields{
				Certificates: map[string]*x509.Certificate{testWechatPayVerifierPlatformSerialNumber: certificate},
			},
			args: args{
				ctx:          context.Background(),
				serialNumber: testWechatPayVerifierPlatformSerialNumber,
				signature:    testExpectedSignature,
				message:      "source",
			},
			wantErr: false,
		},
		{
			name: "verify failed",
			fields: fields{
				Certificates: map[string]*x509.Certificate{testWechatPayVerifierPlatformSerialNumber: certificate},
			},
			args: args{
				ctx:          context.Background(),
				serialNumber: testWechatPayVerifierPlatformSerialNumber,
				signature:    testExpectedSignature,
				message:      "wrong source",
			},
			wantErr: true,
		},
		{
			name:   "verify failed with no corresponding certificate",
			fields: fields{Certificates: map[string]*x509.Certificate{}},
			args: args{
				ctx:          context.Background(),
				serialNumber: testWechatPayVerifierPlatformSerialNumber,
				signature:    testExpectedSignature,
				message:      "source",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			verifier := NewSHA256WithRSAVerifier(core.NewCertificateMap(tt.fields.Certificates))
			if err := verifier.Verify(tt.args.ctx, tt.args.serialNumber, tt.args.message,
				tt.args.signature); (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
