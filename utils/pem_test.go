package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testPemUtilCertificateStrWithoutTags = `MIIDVzCCAj+gAwIBAgIJANfOWdH1ItcBMA0GCSqGSIb3DQEBCwUAMEIxCzAJBgNV
BAYTAlhYMRUwEwYDVQQHDAxEZWZhdWx0IENpdHkxHDAaBgNVBAoME0RlZmF1bHQg
Q29tcGFueSBMdGQwHhcNMjEwNDI3MDg1NTIzWhcNMzEwNDI1MDg1NTIzWjBCMQsw
CQYDVQQGEwJYWDEVMBMGA1UEBwwMRGVmYXVsdCBDaXR5MRwwGgYDVQQKDBNEZWZh
dWx0IENvbXBhbnkgTHRkMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA
2VCTd91fnUn73Xy9DLvt/V62TVxRTEEstVdeRaZ3B3leO0pldE806mXO4RwdHXag
HQ4vGeZN0yqm++rDsGK+U3AH7kejyD2pXshNP9Cq5YwbptiLGtjcquw4HNxJQUOm
DeJf2vg6byms9RUipiq4SzbJKqJFlUpbuIPDpSpWz10PYmyCNeDGUUK65E5h2B83
4uxl1zNLYQCrkdBzb8oUxwYeP5a2DNxmjL5lsJML7DGr5znsevnoqGRwTm9fxCGf
y8wus7hwKz6clt3Whmmda7UAdb1c08hEQFVRbF14AR73xbnd8N0obCWJPCbzMCtk
aSef4FdEEgEXJiw0VAJT8wIDAQABo1AwTjAdBgNVHQ4EFgQUT1c7nd/SUO76HSoZ
umNUJv1R5PwwHwYDVR0jBBgwFoAUT1c7nd/SUO76HSoZumNUJv1R5PwwDAYDVR0T
BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAfTjxKRQMzNB/U6ZoCUS+BSNfa2Oh
0plMN6ZuzwiVVZwg1jywvv5yv04koS7Pd4i9E4gt9ZBUQXlpq+A3oOCEEHNRR6b2
kyazGRM7s0OP5X21WrbpSmKmU6K7hkfx30yYs08LVs/Q8DIhvaj1FCFeJzUCzYn/
fHMq4tsbKO0dKAeydPM/nrUZBmaYQVKMVOORGLFjFKVO7JV6Kq/R86ouhjEPgJOe
2xulNBUcjicqtZlBdEh/PWCYP2SpGVDclKm8jeo175T3EVAkdKzzmfpxtMmnMlmq
cTJOU9TxuGvNASMtjj7pYIerTx+xgZDXEVBWFW9PjJ0TV06tCRsgSHItgg==`
	testPemUtilCertificateStr = `-----BEGIN CERTIFICATE-----
MIIDVzCCAj+gAwIBAgIJANfOWdH1ItcBMA0GCSqGSIb3DQEBCwUAMEIxCzAJBgNV
BAYTAlhYMRUwEwYDVQQHDAxEZWZhdWx0IENpdHkxHDAaBgNVBAoME0RlZmF1bHQg
Q29tcGFueSBMdGQwHhcNMjEwNDI3MDg1NTIzWhcNMzEwNDI1MDg1NTIzWjBCMQsw
CQYDVQQGEwJYWDEVMBMGA1UEBwwMRGVmYXVsdCBDaXR5MRwwGgYDVQQKDBNEZWZh
dWx0IENvbXBhbnkgTHRkMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA
2VCTd91fnUn73Xy9DLvt/V62TVxRTEEstVdeRaZ3B3leO0pldE806mXO4RwdHXag
HQ4vGeZN0yqm++rDsGK+U3AH7kejyD2pXshNP9Cq5YwbptiLGtjcquw4HNxJQUOm
DeJf2vg6byms9RUipiq4SzbJKqJFlUpbuIPDpSpWz10PYmyCNeDGUUK65E5h2B83
4uxl1zNLYQCrkdBzb8oUxwYeP5a2DNxmjL5lsJML7DGr5znsevnoqGRwTm9fxCGf
y8wus7hwKz6clt3Whmmda7UAdb1c08hEQFVRbF14AR73xbnd8N0obCWJPCbzMCtk
aSef4FdEEgEXJiw0VAJT8wIDAQABo1AwTjAdBgNVHQ4EFgQUT1c7nd/SUO76HSoZ
umNUJv1R5PwwHwYDVR0jBBgwFoAUT1c7nd/SUO76HSoZumNUJv1R5PwwDAYDVR0T
BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAfTjxKRQMzNB/U6ZoCUS+BSNfa2Oh
0plMN6ZuzwiVVZwg1jywvv5yv04koS7Pd4i9E4gt9ZBUQXlpq+A3oOCEEHNRR6b2
kyazGRM7s0OP5X21WrbpSmKmU6K7hkfx30yYs08LVs/Q8DIhvaj1FCFeJzUCzYn/
fHMq4tsbKO0dKAeydPM/nrUZBmaYQVKMVOORGLFjFKVO7JV6Kq/R86ouhjEPgJOe
2xulNBUcjicqtZlBdEh/PWCYP2SpGVDclKm8jeo175T3EVAkdKzzmfpxtMmnMlmq
cTJOU9TxuGvNASMtjj7pYIerTx+xgZDXEVBWFW9PjJ0TV06tCRsgSHItgg==
-----END CERTIFICATE-----`
	testPemUtilCertificateSerial = `D7CE59D1F522D701`

	testPemUtilPrivateKeyStrWithoutTags = `MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDZUJN33V+dSfvd
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
BDa+8mDLkWu5nHEhOxy2JJZl`
	testPemUtilPrivateKeyStr = `-----BEGIN PRIVATE KEY-----
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
-----END PRIVATE KEY-----`

	testPemUtilPublicKeyStrWithoutTags = `MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2VCTd91fnUn73Xy9DLvt
/V62TVxRTEEstVdeRaZ3B3leO0pldE806mXO4RwdHXagHQ4vGeZN0yqm++rDsGK+
U3AH7kejyD2pXshNP9Cq5YwbptiLGtjcquw4HNxJQUOmDeJf2vg6byms9RUipiq4
SzbJKqJFlUpbuIPDpSpWz10PYmyCNeDGUUK65E5h2B834uxl1zNLYQCrkdBzb8oU
xwYeP5a2DNxmjL5lsJML7DGr5znsevnoqGRwTm9fxCGfy8wus7hwKz6clt3Whmmd
a7UAdb1c08hEQFVRbF14AR73xbnd8N0obCWJPCbzMCtkaSef4FdEEgEXJiw0VAJT
8wIDAQAB`
	testPemUtilPublicKeyStr = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2VCTd91fnUn73Xy9DLvt
/V62TVxRTEEstVdeRaZ3B3leO0pldE806mXO4RwdHXagHQ4vGeZN0yqm++rDsGK+
U3AH7kejyD2pXshNP9Cq5YwbptiLGtjcquw4HNxJQUOmDeJf2vg6byms9RUipiq4
SzbJKqJFlUpbuIPDpSpWz10PYmyCNeDGUUK65E5h2B834uxl1zNLYQCrkdBzb8oU
xwYeP5a2DNxmjL5lsJML7DGr5znsevnoqGRwTm9fxCGfy8wus7hwKz6clt3Whmmd
a7UAdb1c08hEQFVRbF14AR73xbnd8N0obCWJPCbzMCtkaSef4FdEEgEXJiw0VAJT
8wIDAQAB
-----END PUBLIC KEY-----`
)

func TestLoadCertificate(t *testing.T) {
	type args struct {
		certificateStr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "load certificate from str success",
			args:    args{certificateStr: testPemUtilCertificateStr},
			wantErr: false,
		},
		{
			name:    "error loading certificate without tags",
			args:    args{certificateStr: testPemUtilCertificateStrWithoutTags},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := LoadCertificate(tt.args.certificateStr)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestLoadPrivateKey(t *testing.T) {
	type args struct {
		privateKeyStr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "load private key from str success",
			args:    args{privateKeyStr: testPemUtilPrivateKeyStr},
			wantErr: false,
		},
		{
			name:    "error loading private key without tags",
			args:    args{privateKeyStr: testPemUtilPrivateKeyStrWithoutTags},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := LoadPrivateKey(tt.args.privateKeyStr)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestLoadPublicKey(t *testing.T) {
	type args struct {
		publicKeyStr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "load public key from str success",
			args:    args{publicKeyStr: testPemUtilPublicKeyStr},
			wantErr: false,
		},
		{
			name:    "error loading public key without tags",
			args:    args{publicKeyStr: testPemUtilPublicKeyStrWithoutTags},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := LoadPublicKey(tt.args.publicKeyStr)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestGetCertificateSerialNumber(t *testing.T) {
	certificate, err := LoadCertificate(testPemUtilCertificateStr)
	require.NoError(t, err)

	serial := GetCertificateSerialNumber(*certificate)

	assert.Equal(t, testPemUtilCertificateSerial, serial)
}
