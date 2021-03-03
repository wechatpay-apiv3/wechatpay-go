package setting

import (
	"net/http"
	"testing"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/credentials"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/validators"
)

func TestDialSettings_Validate(t *testing.T) {
	type fields struct {
		HTTPClient *http.Client
		UserAgent  string
		Credential auth.Credential
		Validator  auth.Validator
		Timeout    time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid",
			fields: fields{
				Credential: &credentials.WechatPayCredentials{},
				Validator:  &validators.WechatPayValidator{}},
			wantErr: false,
		},
		{
			name: "without validator invalid",
			fields: fields{
				Credential: &credentials.WechatPayCredentials{},
			},
			wantErr: true,
		},
		{
			name: "without credential invalid",
			fields: fields{
				Validator: &validators.WechatPayValidator{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &DialSettings{
				HTTPClient: tt.fields.HTTPClient,
				UserAgent:  tt.fields.UserAgent,
				Credential: tt.fields.Credential,
				Validator:  tt.fields.Validator,
				Timeout:    tt.fields.Timeout,
			}
			if err := ds.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
