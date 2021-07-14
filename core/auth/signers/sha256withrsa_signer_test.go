// Copyright 2021 Tencent Inc. All rights reserved.

package signers

import (
	"context"
	"crypto/rsa"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

const (
	testPrivateKeyStr = `-----BEGIN TESTING KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDZUJN33V+dSfvd
fL0Mu+39XrZNXFFMQSy1V15FpncHeV47SmV0TzTqZc7hHB0ddqAdDi8Z5k3TKqb7
6sOwYr5TcAfuR6PIPaleyE0/0KrljBum2Isa2Nyq7Dgc3ElBQ6YN4l/a+DpvKaz1
FSKmKrhLNskqokWVSlu4g8OlKlbPXQ9ibII14MZRQrrkTmHYHzfi7GXXM0thAKuR
0HNvyhTHBh4/lrYM3GaMvmWwkwvsMavnOex6+eioZHBOb1/EIZ/LzC6zuHArPpyW
3daGaZ1rtQB1vVzTyERAVVFsXXgBHvfFud3w3ShsJYk8JvMwK2RpJ5/gV0QSARcm
LDRUAlPzAgMBAAECggEBAMc7rDeUaXiWv6bMGbZ3BTXpg1FhdddnWUnYE8HfX/km
OFI7XtBHXcgYFpcjYz4D5787pcsk7ezPidAj58zqenuclmjKnUmT3pfbI5eCA2v4
C9HnbYDrmUPK1ZcADtka4D6ScDccpNYNa1g2TFHzkIrEa6H+q7S3O2fqxY/DRVtN
0JIXalBb8daaqL5QVzSmM2BMVnHy+YITJWIkP2a3pKs9C0W65JGDsnG0wVrHinHF
+cnhFZIbaPEI//DAFMc9NkrWOKVRTEgcCUxCFaHOZVNxDWZD7A2ZfJB2rK6eg//y
gEiFDR2h6mTaDowMB4YF2n2dsIO4/dCG8vPHI20jn4ECgYEA/ZGu6lEMlO0XZnam
AZGtiNgLcCfM/C2ZERZE7QTRPZH1WdK92Al9ndldsswFw4baJrJLCmghjF/iG4zi
hhBvLnOLksnZUfjdumxoHDWXo2QBWbI5QsWIE7AuTiWgWj1I7X4fCXSQf6i+M/y2
6TogQ7d0ANpZFyOkTNMn/tiJvLECgYEA22XqlamG/yfAGWery5KNH2DGlTIyd6xJ
WtJ9j3jU99lZ0bCQ5xhiBbU9ImxCi3zgTsoqLWgA/p00HhNFNoUcTl9ofc0G3zwT
D1y0ZzcnVKxGJdZ6ohW52V0hJStAigtjYAsUgjm7//FH7PiQDBDP1Wa6xSRkDQU/
aSbQxvEE8+MCgYEA3bb8krW7opyM0XL9RHH0oqsFlVO30Oit5lrqebS0oHl3Zsr2
ZGgoBlWBsEzk3UqUhTFwm/DhJLTSJ/TQPRkxnhQ5/mewNhS9C7yua7wQkzVmWN+V
YeUGTvDGDF6qDz12/vJAgSwDDRym8x4NcXD5tTw7mmNRcwIfL22SkysThIECgYAV
BgccoEoXWS/HP2/u6fQr9ZIR6eV8Ij5FPbZacTG3LlS1Cz5XZra95UgebFFUHHtC
EY1JHJY7z8SWvTH8r3Su7eWNaIAoFBGffzqqSVazfm6aYZsOvRY6BfqPHT3p/H1h
Tq6AbBffxrcltgvXnCTORjHPglU0CjSxVs7awW3AEQKBgB5WtaC8VLROM7rkfVIq
+RXqE5vtJfa3e3N7W3RqxKp4zHFAPfr82FK5CX2bppEaxY7SEZVvVInKDc5gKdG/
jWNRBmvvftZhY59PILHO2X5vO4FXh7suEjy6VIh0gsnK36mmRboYIBGsNuDHjXLe
BDa+8mDLkWu5nHEhOxy2JJZl
-----END TESTING KEY-----`
	testCertificateSerial = `F5765756002FDD77`
	testExpectedSignature = "BKyAfU4iMCuvXMXS0Wzam3V/cnxZ+JaqigPM5OhljS2iOT95OO6Fsuml2JkFANJU9K6q9bLlDhPXuoVz+pp4hAm6" +
		"pHU4ld815U4jsKu1RkyaII+1CYBUYC8TK0XtJ8FwUXXz8vZHh58rrAVN1XwNyvD1vfpxrMT4SL536GLwvpUHlCqIMzoZUguLli/K8V29QiOh" +
		"uH6IEqLNJn8e9b3nwNcQ7be3CzYGpDAKBfDGPCqCv8Rw5zndhlffk2FEA70G4hvMwe51qMN/RAJbknXG23bSlObuTCN7Ndj1aJGH6/L+hdwf" +
		"LpUtJm4QYVazzW7DFD27EpSQEqA8bX9+8m1rLg=="
	testMessage = "source"
	testMchID   = "1234567890"
)

func testingKey(s string) string { return strings.ReplaceAll(s, "TESTING KEY", "PRIVATE KEY") }

func TestSha256WithRSASigner_Sign(t *testing.T) {
	privateKey, err := utils.LoadPrivateKey(testingKey(testPrivateKeyStr))
	require.NoError(t, err)

	type args struct {
		mchID      string
		certSerial string
		privateKey *rsa.PrivateKey

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
			name: "Sha256WithRSASigner_Sign success",
			args: args{
				mchID:      testMchID,
				certSerial: testCertificateSerial,
				privateKey: privateKey,

				message: testMessage,
			},
			want: &auth.SignatureResult{
				MchID:               testMchID,
				CertificateSerialNo: testCertificateSerial,
				Signature:           testExpectedSignature,
			},
			wantErr: false,
		},
		{
			name: "Sha256WithRSASigner_Sign err when unset privateKey",
			args: args{
				mchID:      testMchID,
				certSerial: testCertificateSerial,
				privateKey: nil,

				message: testMessage,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SHA256WithRSASigner{
				MchID:               tt.args.mchID,
				CertificateSerialNo: tt.args.certSerial,
				PrivateKey:          tt.args.privateKey,
			}
			got, err := s.Sign(context.Background(), tt.args.message)
			require.Equal(t, tt.wantErr, err != nil)

			if err == nil {
				require.NotNil(t, got)

				assert.Equal(t, tt.want.MchID, got.MchID)
				assert.Equal(t, tt.want.CertificateSerialNo, got.CertificateSerialNo)
				assert.Equal(t, tt.want.Signature, got.Signature)
			}
		})
	}
}
