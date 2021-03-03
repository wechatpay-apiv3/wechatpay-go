package credentials

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/signers"

	"github.com/agiledragon/gomonkey"
)

func TestWechatPayCredentials_GenerateAuthorizationHeader(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()
	signOutputs := []gomonkey.OutputCell{
		{Values: gomonkey.Params{"", nil}, Times: 1},
		{Values: gomonkey.Params{"", fmt.Errorf("sign err")}, Times: 1},
	}
	generateNonceStrOutputs := []gomonkey.OutputCell{
		{Values: gomonkey.Params{"123", nil}, Times: 2},
		{Values: gomonkey.Params{"", fmt.Errorf("generate nonce str err")}, Times: 1},
	}
	c := &WechatPayCredentials{Signer: &signers.Sha256WithRSASigner{}}
	patches.ApplyMethodSeq(reflect.TypeOf(c.Signer), "Sign", signOutputs)
	patches.ApplyFuncSeq(generateNonceStr, generateNonceStrOutputs)
	type fields struct {
		Signer auth.Signer
	}
	type args struct {
		ctx          context.Context
		method       string
		canonicalURL string
		signBody     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "get token success",
			wantErr: false,
		},
		{
			name:    "get token with sign err",
			wantErr: true,
		},
		{
			name:    "get token with generate nonce string err",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.GenerateAuthorizationHeader(tt.args.ctx, tt.args.method, tt.args.canonicalURL, tt.args.signBody)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestGenerateNonceStr(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "generate nonce str success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := generateNonceStr()
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateNonceStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
