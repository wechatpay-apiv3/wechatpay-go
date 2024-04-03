// Copyright 2021 Tencent Inc. All rights reserved.

package credentials

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"github.com/xhd2015/xgo/runtime/mock"
)

type mockSigner struct {
	MchID               string
	CertificateSerialNo string
}

func (s *mockSigner) Sign(_ context.Context, message string) (*auth.SignatureResult, error) {
	result := &auth.SignatureResult{
		MchID:               s.MchID,
		CertificateSerialNo: s.CertificateSerialNo,
		Signature:           "Sign:" + message,
	}
	return result, nil
}

func (s *mockSigner) Algorithm() string {
	return "Mock"
}

type mockErrorSigner struct {
}

func (s *mockErrorSigner) Sign(_ context.Context, message string) (*auth.SignatureResult, error) {
	return nil, fmt.Errorf("mock sign error")
}

func (s *mockErrorSigner) Algorithm() string {
	return "ErrorMock"
}

const (
	testMchID             = "1234567890"
	testCertificateSerial = "0123456789ABC"
	mockNonce             = "A1B2C3D4E5F6G7"
	mockTimestamp         = 1624523846
)

func TestWechatPayCredentials_GenerateAuthorizationHeader(t *testing.T) {
	mock.Patch(utils.GenerateNonce, func() (string, error) {
		return mockNonce, nil
	})

	mock.Patch(time.Now, func() time.Time {
		return time.Unix(mockTimestamp, 0)
	})

	signer := mockSigner{
		MchID:               testMchID,
		CertificateSerialNo: testCertificateSerial,
	}

	type args struct {
		signer auth.Signer

		ctx          context.Context
		method       string
		canonicalURL string
		signBody     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    string
	}{
		{
			name: "gen success without body",
			args: args{
				signer: &signer,

				ctx:          context.Background(),
				method:       "GET",
				canonicalURL: "/v3/certificates",
				signBody:     "",
			},
			wantErr: false,
			want: `WECHATPAY2-Mock mchid="1234567890",nonce_str="A1B2C3D4E5F6G7",timestamp="1624523846",` +
				`serial_no="0123456789ABC",signature=` +
				"\"Sign:GET\n/v3/certificates\n1624523846\nA1B2C3D4E5F6G7\n\n\"",
		},
		{
			name: "gen success with body",
			args: args{
				signer: &signer,

				ctx:          context.Background(),
				method:       "POST",
				canonicalURL: "/v3/certificates",
				signBody:     "Hello World!\n",
			},
			wantErr: false,
			want: `WECHATPAY2-Mock mchid="1234567890",nonce_str="A1B2C3D4E5F6G7",timestamp="1624523846",` +
				`serial_no="0123456789ABC",signature=` +
				"\"Sign:POST\n/v3/certificates\n1624523846\nA1B2C3D4E5F6G7\nHello World!\n\n\"",
		},
		{
			name: "gen error wihout signer",
			args: args{
				signer: nil,

				ctx:          context.Background(),
				method:       "post",
				canonicalURL: "/v3/certificates",
				signBody:     "Hello World!\n",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				credential := WechatPayCredentials{Signer: tt.args.signer}

				authorization, err := credential.GenerateAuthorizationHeader(
					tt.args.ctx, tt.args.method, tt.args.canonicalURL, tt.args.signBody,
				)
				require.Equal(t, tt.wantErr, err != nil)
				require.Equal(t, tt.want, authorization)
			},
		)
	}
}

func TestWechatPayCredentials_GenerateAuthorizationHeaderErrorGenerateNonce(t *testing.T) {
	mockGenerateNonceErr := fmt.Errorf("generate nonce error")

	mock.Patch(utils.GenerateNonce, func() (string, error) {
		return "", mockGenerateNonceErr
	})

	signer := mockSigner{
		MchID:               testMchID,
		CertificateSerialNo: testCertificateSerial,
	}
	credential := WechatPayCredentials{Signer: &signer}

	authorization, err := credential.GenerateAuthorizationHeader(context.Background(), "GET", "/v3/certificates", "")
	require.Error(t, err)
	assert.Empty(t, authorization)
}

func TestWechatPayCredentials_GenerateAuthorizationHeaderErrorSigner(t *testing.T) {
	mock.Patch(utils.GenerateNonce, func() (string, error) {
		return mockNonce, nil
	})
	mock.Patch(time.Now, func() time.Time {
		return time.Unix(mockTimestamp, 0)
	})

	signer := mockErrorSigner{}
	credential := WechatPayCredentials{Signer: &signer}
	authorization, err := credential.GenerateAuthorizationHeader(context.Background(), "GET", "/v3/certificates", "")
	require.Error(t, err)
	assert.Empty(t, authorization)
}
