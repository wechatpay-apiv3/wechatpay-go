package validators

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"

	"github.com/agiledragon/gomonkey"
)

func TestWechatPayResponseValidator_Validate(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()
	validateParametersOutputs := []gomonkey.OutputCell{
		{Values: gomonkey.Params{nil}, Times: 4},
		{Values: gomonkey.Params{fmt.Errorf("validate parameters err")}, Times: 1},
	}
	buildMessageOutputs := []gomonkey.OutputCell{
		{Values: gomonkey.Params{"", nil}, Times: 3},
		{Values: gomonkey.Params{"", fmt.Errorf("build message err")}, Times: 1},
	}
	verifyOutputs := []gomonkey.OutputCell{
		{Values: gomonkey.Params{nil}, Times: 1},
		{Values: gomonkey.Params{fmt.Errorf("verift err")}, Times: 1},
	}

	validator := NewWechatPayResponseValidator(&verifiers.SHA256WithRSAVerifier{})
	validHeader := map[string][]string{
		consts.WechatPaySignature: {base64.StdEncoding.EncodeToString([]byte("1"))},
		consts.WechatPaySerial:    {"1"},
		consts.RequestID:          {"1"},
	}
	inValidHeader := map[string][]string{
		consts.WechatPaySignature: {"1"},
		consts.WechatPaySerial:    {"1"},
		consts.RequestID:          {"1"},
	}
	patches.ApplyFuncSeq(checkParameters, validateParametersOutputs)
	patches.ApplyFuncSeq(buildMessage, buildMessageOutputs)
	patches.ApplyMethodSeq(reflect.TypeOf(validator.verifier), "Verify", verifyOutputs)
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
				ctx:      context.Background(),
				response: &http.Response{Header: validHeader},
			},
		},
		{
			name: "response validate verify err",
			args: args{
				ctx:      context.Background(),
				response: &http.Response{Header: validHeader},
			},
			wantErr: true,
		},
		{
			name: "response validate decode signature err",
			args: args{
				ctx:      context.Background(),
				response: &http.Response{Header: inValidHeader},
			},
			wantErr: true,
		},
		{
			name: "response validate decode build message err",
			args: args{
				ctx:      context.Background(),
				response: &http.Response{Header: inValidHeader},
			},
			wantErr: true,
		},
		{
			name: "response validate decode check parameters err",
			args: args{
				ctx:      context.Background(),
				response: &http.Response{Header: inValidHeader},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validator.Validate(tt.args.ctx, tt.args.response); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
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
			args: args{response: &http.Response{
				Header: map[string][]string{
					consts.RequestID:          {"1"},
					consts.WechatPaySerial:    {"1"},
					consts.WechatPaySignature: {"1"},
					consts.WechatPayTimestamp: {strconv.FormatInt(time.Now().Unix(), 10)},
					consts.WechatPayNonce:     {"1"},
				},
			}},
			wantErr: false,
		},
		{
			name: "time is expire",
			args: args{response: &http.Response{
				Header: map[string][]string{
					consts.RequestID:          {"1"},
					consts.WechatPaySerial:    {"1"},
					consts.WechatPaySignature: {"1"},
					consts.WechatPayTimestamp: {strconv.FormatInt(0, 10)},
					consts.WechatPayNonce:     {"1"},
				},
			}},
			wantErr: true,
		},
		{
			name: "nonce is empty",
			args: args{response: &http.Response{
				Header: map[string][]string{
					consts.RequestID:          {"1"},
					consts.WechatPaySerial:    {"1"},
					consts.WechatPaySignature: {"1"},
					consts.WechatPayTimestamp: {strconv.FormatInt(0, 10)},
				},
			}},
			wantErr: true,
		},
		{
			name: "timestamp is empty",
			args: args{response: &http.Response{
				Header: map[string][]string{
					consts.RequestID:          {"1"},
					consts.WechatPaySerial:    {"1"},
					consts.WechatPaySignature: {"1"},
				},
			}},
			wantErr: true,
		},
		{
			name: "signature is empty",
			args: args{response: &http.Response{
				Header: map[string][]string{
					consts.RequestID:          {"1"},
					consts.WechatPaySerial:    {"1"},
					consts.WechatPaySignature: {"1"},
				},
			}},
			wantErr: true,
		},
		{
			name: "serial number is empty",
			args: args{response: &http.Response{
				Header: map[string][]string{
					consts.RequestID: {"1"},
				},
			}},
			wantErr: true,
		},
		{
			name: "request id is empty",
			args: args{response: &http.Response{
				Header: map[string][]string{},
			}},
			wantErr: true,
		},
	}
	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
		})
	}
}
