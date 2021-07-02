package validators

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
)

type mockVerifier struct {
}

func (v *mockVerifier) Verify(ctx context.Context, serialNumber string, message string, signature string) error {
	if serialNumber+"-"+message == signature {
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
						consts.WechatPaySignature: {"SERIAL1234567890-" + mockTimestampStr + "\nNONCE1234567890\nBODY\n"},
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
						consts.WechatPaySignature: {"SERIAL1234567890-" + mockTimestampStr + "\nNONCE1234567890\n\n"},
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
			name: "response validate verify err",
			args: args{
				ctx: context.Background(),
				response: &http.Response{
					Header: http.Header{
						consts.WechatPaySignature: {"SERIAL1234567890-" + mockTimestampStr + "\nNONCE1234567890\n"},
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
			name: "response validate decode check parameters err",
			args: args{
				ctx: context.Background(),
				response: &http.Response{
					Header: http.Header{
						consts.WechatPaySignature: {"SERIAL1234567890-" + mockTimestampStr + "\nNONCE1234567890\n"},
						consts.WechatPaySerial:    {"SERIAL1234567890"},
						consts.WechatPayTimestamp: {mockTimestampStr},
						consts.WechatPayNonce:     {"NONCE1234567890"},
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

func Test_checkParameters(t *testing.T) {
	type args struct {
		response *http.Response
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "parameter is valid",
			args: args{
				response: &http.Response{
					Header: map[string][]string{
						consts.RequestID:          {"1"},
						consts.WechatPaySerial:    {"1"},
						consts.WechatPaySignature: {"1"},
						consts.WechatPayTimestamp: {strconv.FormatInt(time.Now().Unix(), 10)},
						consts.WechatPayNonce:     {"1"},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "time is expire",
			args: args{
				response: &http.Response{
					Header: map[string][]string{
						consts.RequestID:          {"1"},
						consts.WechatPaySerial:    {"1"},
						consts.WechatPaySignature: {"1"},
						consts.WechatPayTimestamp: {"0"},
						consts.WechatPayNonce:     {"1"},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "nonce is empty",
			args: args{
				response: &http.Response{
					Header: map[string][]string{
						consts.RequestID:          {"1"},
						consts.WechatPaySerial:    {"1"},
						consts.WechatPaySignature: {"1"},
						consts.WechatPayTimestamp: {strconv.FormatInt(0, 10)},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "timestamp is empty",
			args: args{
				response: &http.Response{
					Header: map[string][]string{
						consts.RequestID:          {"1"},
						consts.WechatPaySerial:    {"1"},
						consts.WechatPaySignature: {"1"},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "signature is empty",
			args: args{
				response: &http.Response{
					Header: map[string][]string{
						consts.RequestID:          {"1"},
						consts.WechatPaySerial:    {"1"},
						consts.WechatPaySignature: {"1"},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "serial number is empty",
			args: args{
				response: &http.Response{
					Header: map[string][]string{
						consts.RequestID: {"1"},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "request id is empty",
			args: args{
				response: &http.Response{
					Header: map[string][]string{},
				},
			},
			wantErr: true,
		},
	}
	ctx := context.Background()
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				var (
					body []byte
					err  error
				)
				if tt.args.response.Body != nil {
					body, err = ioutil.ReadAll(tt.args.response.Body)
					require.NoError(t, err)
				}

				if err := checkParameters(ctx, tt.args.response.Header, body); (err != nil) != tt.wantErr {
					t.Errorf("validateParameters() error = %v, wantErr %v", err, tt.wantErr)
				}
			},
		)
	}
}

func TestNullValidator_Validate(t *testing.T) {
	nullValidator := NullValidator{}

	assert.NoError(t, nullValidator.Validate(context.Background(), &http.Response{}))
	assert.NoError(t, nullValidator.Validate(context.Background(), nil))
}
