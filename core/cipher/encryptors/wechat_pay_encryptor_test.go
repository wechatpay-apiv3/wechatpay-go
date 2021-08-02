// Copyright 2021 Tencent Inc. All rights reserved.

package encryptors

import (
	"context"
	"crypto/x509"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

var testCertStrList = [2]string{
	// Serial=D7CE59D1F522D701 NotBefore=Apr 27 08:55:23 2021 GMT
	`-----BEGIN CERTIFICATE-----
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
-----END CERTIFICATE-----`,
	// Serial=F5765756002FDD77 NotBefore=Apr 27 08:40:32 2021 GMT
	`-----BEGIN CERTIFICATE-----
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
-----END CERTIFICATE-----`,
}

const privateKeyStr = `-----BEGIN TESTING KEY-----
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

func testingKey(s string) string { return strings.ReplaceAll(s, "TESTING KEY", "PRIVATE KEY") }

func initWechatPayEncryptor() (*WechatPayEncryptor, error) {
	l := make([]*x509.Certificate, 0, 2)
	for _, certStr := range testCertStrList {
		cert, err := utils.LoadCertificate(certStr)
		if err != nil {
			return nil, err
		}
		l = append(l, cert)
	}

	return NewWechatPayEncryptor(core.NewCertificateMapWithList(l)), nil
}

func TestWechatPayEncryptor_SelectCertificate(t *testing.T) {
	e, err := initWechatPayEncryptor()
	require.NoError(t, err)

	serial, err := e.SelectCertificate(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "D7CE59D1F522D701", serial)
}

func TestWechatPayEncryptor_Encrypt(t *testing.T) {
	e, err := initWechatPayEncryptor()
	require.NoError(t, err)

	const serial = "F5765756002FDD77"
	const plaintext = "hello world"

	ciphertext, err := e.Encrypt(context.Background(), serial, plaintext)
	require.NoError(t, err)

	privateKey, err := utils.LoadPrivateKey(testingKey(privateKeyStr))
	require.NoError(t, err)

	newPlainText, err := utils.DecryptOAEP(ciphertext, privateKey)
	require.NoError(t, err)
	assert.Equal(t, newPlainText, plaintext)
}

func TestWechatPayEncryptor_EncryptEmpty(t *testing.T) {
	e, err := initWechatPayEncryptor()
	require.NoError(t, err)

	const serial = "F5765756002FDD77"
	const plaintext = ""

	ciphertext, err := e.Encrypt(context.Background(), serial, plaintext)
	require.NoError(t, err)
	assert.Equal(t, "", ciphertext)
}

func TestWechatPayEncryptor_EncryptWithWrongSerial(t *testing.T) {
	e, err := initWechatPayEncryptor()
	require.NoError(t, err)

	const serial = "unknown serial"
	const plaintext = ""

	_, err = e.Encrypt(context.Background(), serial, plaintext)
	require.Error(t, err)
}

func TestMockEncryptor_Encrypt(t *testing.T) {
	e := MockEncryptor{Serial: "F5765756002FDD77"}

	cipertext, err := e.Encrypt(context.Background(), "F5765756002FDD77", "hehe")
	require.NoError(t, err)
	assert.Equal(t, "Encryptedhehe", cipertext)
}

func TestMockEncryptor_EncryptWithWrontSerial(t *testing.T) {
	e := MockEncryptor{Serial: "F5765756002FDD77"}

	_, err := e.Encrypt(context.Background(), "wrong serial", "hehe")
	require.Error(t, err)
}
