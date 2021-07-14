// Copyright 2021 Tencent Inc. All rights reserved.

package validators

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
)

type mockVerifier struct {
}

func (v *mockVerifier) Verify(ctx context.Context, serialNumber string, message string, signature string) error {
	if "["+serialNumber+"-"+message+"]" == signature {
		return nil
	}

	return fmt.Errorf("verification failed")
}

func TestWechatPayResponseValidator_Validate(t *testing.T) {
	mockTimestamp := time.Now().Unix()
	mockTimestampStr := fmt.Sprintf("%d", mockTimestamp)

	validator := NewWechatPayResponseValidator(&mockVerifier{})

	type args struct {
		ctx      context.Context
		response *http.Response
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "response validate success",
			args: args{
				ctx: context.Background(),
				response: &http.Response{
					Header: http.Header{
						consts.WechatPaySignature: {
							"[SERIAL1234567890-" + mockTimestampStr + "\nNONCE1234567890\nBODY\n]",
						},
						consts.WechatPaySerial:    {"SERIAL1234567890"},
						consts.WechatPayTimestamp: {mockTimestampStr},
						consts.WechatPayNonce:     {"NONCE1234567890"},
						consts.RequestID:          {"any-request-id"},
					},
					Body: ioutil.NopCloser(bytes.NewBuffer([]byte("BODY"))),
				},
			},
			wantErr: false,
		},
		{
			name: "response validate success without body",
			args: args{
				ctx: context.Background(),
				response: &http.Response{
					Header: http.Header{
						consts.WechatPaySignature: {
							"[SERIAL1234567890-" + mockTimestampStr + "\nNONCE1234567890\n\n]",
						},
						consts.WechatPaySerial:    {"SERIAL1234567890"},
						consts.WechatPayTimestamp: {mockTimestampStr},
						consts.WechatPayNonce:     {"NONCE1234567890"},
						consts.RequestID:          {"any-request-id"},
					},
					Body: ioutil.NopCloser(bytes.NewBuffer([]byte(""))),
				},
			},
			wantErr: false,
		},
		{
			name: "response validate success without RequestID",
			args: args{
				ctx: context.Background(),
				response: &http.Response{
					Header: http.Header{
						consts.WechatPaySignature: {
							"[SERIAL1234567890-" + mockTimestampStr + "\nNONCE1234567890\n\n]",
						},
						consts.WechatPaySerial:    {"SERIAL1234567890"},
						consts.WechatPayTimestamp: {mockTimestampStr},
						consts.WechatPayNonce:     {"NONCE1234567890"},
					},
					Body: ioutil.NopCloser(bytes.NewBuffer([]byte(""))),
				},
			},
			wantErr: false,
		},
		{
			name: "response validate error with error signature",
			args: args{
				ctx: context.Background(),
				response: &http.Response{
					Header: http.Header{
						consts.WechatPaySignature: {"[SERIAL1234567890-" + mockTimestampStr + "\nNONCE1234567\n\n]"},
						consts.WechatPaySerial:    {"SERIAL1234567890"},
						consts.WechatPayTimestamp: {mockTimestampStr},
						consts.WechatPayNonce:     {"NONCE1234567890"},
						consts.RequestID:          {"any-request-id"},
					},
					Body: ioutil.NopCloser(bytes.NewBuffer([]byte(""))),
				},
			},
			wantErr: true,
		},
		{
			name: "response validate error missing signature",
			args: args{
				ctx: context.Background(),
				response: &http.Response{
					Header: http.Header{
						consts.WechatPaySerial:    {"SERIAL1234567890"},
						consts.WechatPayTimestamp: {mockTimestampStr},
						consts.WechatPayNonce:     {"NONCE1234567890"},
						consts.RequestID:          {"any-request-id"},
					},
					Body: ioutil.NopCloser(bytes.NewBuffer([]byte(""))),
				},
			},
			wantErr: true,
		},
		{
			name: "response validate error missing serial",
			args: args{
				ctx: context.Background(),
				response: &http.Response{
					Header: http.Header{
						consts.WechatPaySignature: {"[SERIAL1234567890-" + mockTimestampStr + "\nNONCE1234567890\n]"},
						consts.WechatPayTimestamp: {mockTimestampStr},
						consts.WechatPayNonce:     {"NONCE1234567890"},
						consts.RequestID:          {"any-request-id"},
					},
					Body: ioutil.NopCloser(bytes.NewBuffer([]byte(""))),
				},
			},
			wantErr: true,
		},
		{
			name: "response validate error missing timestamp",
			args: args{
				ctx: context.Background(),
				response: &http.Response{
					Header: http.Header{
						consts.WechatPaySignature: {"[SERIAL1234567890-" + mockTimestampStr + "\nNONCE1234567890\n]"},
						consts.WechatPaySerial:    {"SERIAL1234567890"},
						consts.WechatPayNonce:     {"NONCE1234567890"},
						consts.RequestID:          {"any-request-id"},
					},
					Body: ioutil.NopCloser(bytes.NewBuffer([]byte(""))),
				},
			},
			wantErr: true,
		},
		{
			name: "response validate error invalid timestamp",
			args: args{
				ctx: context.Background(),
				response: &http.Response{
					Header: http.Header{
						consts.WechatPaySignature: {"[SERIAL1234567890-" + mockTimestampStr + "\nNONCE1234567890\n]"},
						consts.WechatPaySerial:    {"SERIAL1234567890"},
						consts.WechatPayTimestamp: {"invalid timestamp"},
						consts.WechatPayNonce:     {"NONCE1234567890"},
						consts.RequestID:          {"any-request-id"},
					},
					Body: ioutil.NopCloser(bytes.NewBuffer([]byte(""))),
				},
			},
			wantErr: true,
		},
		{
			name: "response validate error missing nonce",
			args: args{
				ctx: context.Background(),
				response: &http.Response{
					Header: http.Header{
						consts.WechatPaySignature: {"[SERIAL1234567890-" + mockTimestampStr + "\nNONCE1234567890\n]"},
						consts.WechatPaySerial:    {"SERIAL1234567890"},
						consts.WechatPayTimestamp: {mockTimestampStr},
						consts.RequestID:          {"any-request-id"},
					},
					Body: ioutil.NopCloser(bytes.NewBuffer([]byte(""))),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if err := validator.Validate(tt.args.ctx, tt.args.response); (err != nil) != tt.wantErr {
					t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				}
			},
		)
	}
}

func TestWechatPayResponseValidator_WithoutVerifierShouldFail(t *testing.T) {
	mockTimestamp := time.Now().Unix()
	mockTimestampStr := fmt.Sprintf("%d", mockTimestamp)

	invalidValidator := NewWechatPayResponseValidator(nil)

	response := &http.Response{
		Header: http.Header{
			consts.WechatPaySignature: {
				"[SERIAL1234567890-" + mockTimestampStr + "\nNONCE1234567890\nBODY\n]",
			},
			consts.WechatPaySerial:    {"SERIAL1234567890"},
			consts.WechatPayTimestamp: {mockTimestampStr},
			consts.WechatPayNonce:     {"NONCE1234567890"},
			consts.RequestID:          {"any-request-id"},
		},
		Body: ioutil.NopCloser(bytes.NewBuffer([]byte("BODY"))),
	}

	err := invalidValidator.Validate(context.Background(), response)
	assert.Error(t, err)
}

func TestNullValidator_Validate(t *testing.T) {
	nullValidator := NullValidator{}

	assert.NoError(t, nullValidator.Validate(context.Background(), &http.Response{}))
	assert.NoError(t, nullValidator.Validate(context.Background(), nil))
}
