// Copyright 2021 Tencent Inc. All rights reserved.

package verifiers

import (
	"context"
	"crypto/rsa"
	"testing"

	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

const (
	testPubKeyID = "F5765756002FDD77"
	testPubKey   = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2VCTd91fnUn73Xy9DLvt
/V62TVxRTEEstVdeRaZ3B3leO0pldE806mXO4RwdHXagHQ4vGeZN0yqm++rDsGK+
U3AH7kejyD2pXshNP9Cq5YwbptiLGtjcquw4HNxJQUOmDeJf2vg6byms9RUipiq4
SzbJKqJFlUpbuIPDpSpWz10PYmyCNeDGUUK65E5h2B834uxl1zNLYQCrkdBzb8oU
xwYeP5a2DNxmjL5lsJML7DGr5znsevnoqGRwTm9fxCGfy8wus7hwKz6clt3Whmmd
a7UAdb1c08hEQFVRbF14AR73xbnd8N0obCWJPCbzMCtkaSef4FdEEgEXJiw0VAJT
8wIDAQAB
-----END PUBLIC KEY-----`
	// testExpectedSignature = "BKyAfU4iMCuvXMXS0Wzam3V/cnxZ+JaqigPM5OhljS2iOT95OO6Fsuml2JkFANJU9" +
	// 	"K6q9bLlDhPXuoVz+pp4hAm6pHU4ld815U4jsKu1RkyaII+1CYBUYC8TK0XtJ8FwUXXz8vZHh58rrAVN1XwNyv" +
	// 	"D1vfpxrMT4SL536GLwvpUHlCqIMzoZUguLli/K8V29QiOhuH6IEqLNJn8e9b3nwNcQ7be3CzYGpDAKBfDGPCq" +
	// 	"Cv8Rw5zndhlffk2FEA70G4hvMwe51qMN/RAJbknXG23bSlObuTCN7Ndj1aJGH6/L+hdwfLpUtJm4QYVazzW7D" +
	// 	"FD27EpSQEqA8bX9+8m1rLg=="
)

var (
	pubKey *rsa.PublicKey
)

func init() {
	var err error
	pubKey, err = utils.LoadPublicKey(testPubKey)
	if err != nil {
		panic(err)
	}
}

func TestWechatPayPubKeyVerifier(t *testing.T) {
	type args struct {
		ctx          context.Context
		serialNumber string
		message      string
		signature    string
	}
	tests := []struct {
		name    string
		fields  *rsa.PublicKey
		args    args
		wantErr bool
	}{
		{
			name:   "verify success",
			fields: pubKey,
			args: args{
				ctx:          context.Background(),
				serialNumber: testPubKeyID,
				signature:    testExpectedSignature,
				message:      "source",
			},
			wantErr: false,
		},
		{
			name:   "verify failed",
			fields: pubKey,
			args: args{
				ctx:          context.Background(),
				serialNumber: testPubKeyID,
				signature:    testExpectedSignature,
				message:      "wrong source",
			},
			wantErr: true,
		},
		{
			name:   "verify failed with null context",
			fields: pubKey,
			args: args{
				ctx:          nil,
				serialNumber: testWechatPayVerifierPlatformSerialNumber,
				signature:    testExpectedSignature,
				message:      "source",
			},
			wantErr: true,
		},
		{
			name:   "verify failed with empty keyId",
			fields: pubKey,
			args: args{
				ctx:          context.Background(),
				serialNumber: "",
				signature:    testExpectedSignature,
				message:      "source",
			},
			wantErr: true,
		},
		{
			name:   "verify failed with empty message",
			fields: pubKey,
			args: args{
				ctx:          context.Background(),
				serialNumber: testPubKeyID,
				signature:    testExpectedSignature,
				message:      "",
			},
			wantErr: true,
		},
		{
			name:   "verify failed with empty signature",
			fields: pubKey,
			args: args{
				ctx:          context.Background(),
				serialNumber: testPubKeyID,
				signature:    "",
				message:      "source",
			},
			wantErr: true,
		},
		{
			name:   "verify failed with non-base64 signature",
			fields: pubKey,
			args: args{
				ctx:          context.Background(),
				serialNumber: testPubKeyID,
				signature:    "invalid base64 signature",
				message:      "source",
			},
			wantErr: true,
		},
		{
			name:   "verify failed with no corresponding pubkey",
			fields: pubKey,
			args: args{
				ctx:          context.Background(),
				serialNumber: "invalid serial number",
				signature:    testExpectedSignature,
				message:      "source",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var verifier = NewSHA256WithRSAPubkeyVerifier(testPubKeyID, *tt.fields)
			if err := verifier.Verify(tt.args.ctx, tt.args.serialNumber, tt.args.message,
				tt.args.signature); (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
