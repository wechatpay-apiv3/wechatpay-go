package downloader_test

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey"
	"github.com/stretchr/testify/require"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

const (
	mockWechatPayPrivateKeyStr = `-----BEGIN PRIVATE KEY-----
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
	mockWechatPayCertificateStr = `-----BEGIN CERTIFICATE-----
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
	mockAPIv3Key             = "mockAPIv3Key1234"
	mockMchID                = "1234567890"
	mockMchCertificateSerial = "BE6DCDA7A5931FA0"
	mockMchPrivateKey        = `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC/KuDQhHMw9Rkv
IUvXrBgCNTqhWAxdu1O4pdSzeaJJYeUQPPP3KLlm6jjvjZg8nOS2zv0hZRGHNviS
Am70+HKsRlZThbB6Kz08c/jPN01CoLJ1bLGFEX4tvcmalM5RGyRZrhXYSWtO/aDr
ezC1HRTRrKCvGq0utVb8wJz3NykX2U5PN3g7VliADFJvAjcLH9ctNwl/DU6+uNVd
3ch/REZIhM/8K3UKYULTBpg3Hx2oKmFdRfay5Fy7Q1snrz3ziIciNKfynu6dWQ0R
f2pvkulTLASDDbQD2k1z/oStV29O/VcnYTuwVhQmOM0v8v6jWNhnqxzlMG8QY4uR
K2UrmvIpAgMBAAECggEBAJmmWgne20MXTMWfynivnqBvrg8uWMohnZoE15/lfLXk
lBroEuKt/c6lJVwNf7YAGKiCV9w2gs8eNM4OXKZS5sBmzE3XX0/iXxau0+WhOlz6
ElXlJg2hULxtFZREVEvGOAJItNAhurlDi8qZOj3oAOrCCGiVVWr8X81I9yCQWlwK
a46UoWixU8E0W4jFFt4AFLbrMPa+7TsKr9gilYcwcNHvKDqYhuqKdGRmO+QtKlPw
4AdlfhwUh1L+TCUcDVb6Zcby+HWUpcx7hBcKRd8sLxPeWcS8imeR8OlVO/8bRJYD
WhFlBRqvIuJqycW5/v01ka3zRe1yO/bl+sxJ6QHPrA0CgYEA+/7+97lmNeKIfFYS
4JD8jiAZeCl9C6CZFj3gCpXE3ChT9wGsHdkw2QZRZkKYDFMhF7KVtiijvexSeWaB
uXelk+3yHhmVUC8DH1aXh2CXGgJhhxE08YzBtFp7oqn8dd2SXQAf37FGjjlhCcnw
JqgE212qdFzY1nrTCjS6ymTdv5sCgYEAwjR1pAXW5mR/PXMH/8i+Wr3ZeZawHByN
/H4wtw5F/mR7xzJup1y6XA0q7BEH+zp5ts0nKzOwIr6uZsBJl2xUuntD/pZCwsN+
6P+mC1JGWKYhXTZ8PASVU834wqUjaCDBxXZqFqkIZAnjQZhpC7meSJ1ANAt6f5zq
/DhoyI8My4sCgYAkxkW3KRs9ad25J2aB1ybEJvMQkh1pgPpWQJldchXUex4lwdy4
TmXOhhmC7tz5j3gY2Tr35l5e7QnsJYVw54EGYYcf1WPw26t8+0oJu5LRfN92spAj
YAM0qq+4QU5SdQ9S+x2rq1c8kisTHqEpQwPSb4zchmAujKYXqzJHLwkdWQKBgCV9
EPR/uBMzdSh8ix/CNZS4r0F8aDSVAoeqKGc91x8bcJVsU9X03XilhhKZ2wuRJyu1
gIkjai3a1zm4hrw3SkfRQbfkc7C7IkWCDCCFWlUEhM5ElUjUrarGyO1yCVqxcBZZ
HHORX7BIBFmGPUjpJPfpexpQ2O3HcckMbpXAn3yvAoGAPxQSNSa8JKGgdD8pZQIC
eQU9Br/ZPXtI1EZdxfMwdwFov+6DtPaq5kt01wACK9ozTHONLFH3diW+HIl+49r6
0q+pf03FucCa+chixggyeyvV5Zi2KrV5CRH+tRVDy75Erf6YgjTzoXrPfZZ1FIzM
wvAZNrxhvtBilUC8adsqhSY=
-----END PRIVATE KEY-----`

	mockNonce = "mockNonce1234"
	data      = "{\"data\":[{\"effective_time\":\"2021-04-27T16:55:23+08:00\",\"encrypt_certificate\":{" +
		"\"algorithm\":\"AEAD_AES_256_GCM\"," +
		"\"associated_data\":\"certificate\"," +
		"\"ciphertext\":\"QHBP2NGCmcEhGWa1f+t5AeImEdkS+XyUx1nB50LLGTgfDdb+EELcl9/GRJR0m9SFcTSlsS67QKemf2DuvPGTFEn" +
		"EIhZ/I26zym7Ift1YJt7ftiVqoU4lv9aopCWD7IrzQBLT7XeYKh0CFHj0fUeqnNrGC+3EMj+NHbDY6zGfW4SfhvzLxO6tvpQ3xbFyT2C" +
		"gxHmiKiVTxkewiSDuaimy3X/SN/Nq901lhEtqBVMFyFOyT1MvTL0RUWpU+l4Wdyz9xGTrrv0pd7regqfXwklx+Y94ULUd7lhSyb77h9j" +
		"8golVvxYUwrUrxWj5ri6APPd8qedWuxkBjOPUAho/q1ETpjYJqK0eXuVWf5EPBS63jvqEE4IfopXOwltPLEME++6u7H4qjhAfzdUvfLQ" +
		"awUe/zPPwWS5GOWRxov7qnuK3aV0+ybZM33hmPwP5LsPNQlDdaQxcNgqbnmpVWjxCPdiiZ52z19AxcB4Ry9LN5IBAEbTnm9WnCp4cOqN" +
		"OXBr2/l+FRLSqZnNTIFbOzXWKw5AngHbciLEvTYEsFNpT/u1SBCU2xQ3Y+Q8OHIn6kmE3AO/gUmXjciLv05ZGo+DCAnqGOVdr2yCLv2j" +
		"maBSIcrComs1WzliuNfFJ5+On7i+rOAbKMaZLdTcB//Lpa4mfkhULhslQYUT5H4XbQqsi5IPGplYvvNh2+ktahI6lYWnCprFeFthyEfy" +
		"LNc9MGtsO1rdsZqoI6ed+KLzMJDOGAhgEQqSJIOzz546/pyn5DnINlxjMSpq6+Zdan+iuiz3Un2idaLf6iPSx1FtdNPi94EjlI+bQSEi" +
		"hKKm4Em2xVgauajO3mK6+JTPNQMVrbtEV9wIGHvPLVm6Uw3OAuLWL0UCn3P6wuxF7+4X+S8s1EKN/3q/w++U93NKdfnpOKe7Vxg7Pg3f" +
		"96WQzJTvRJ7C1Xm4W7GvS09hGjCNGCvbcGXla14X/Y35o6Uf37NlnmnDe2rdcnmYsXE5qooS0ThLdOKdKNybg+ih9iFPcxdTJ67WfvVP" +
		"vaOhLrRm8cDgTYjfVYAj7lxztRYnZE65PTeWUE/mLZTd52g3WEE8ty79KLcqnzARGq7sjdptcZQ4Vw11hHA3PbrxQdyvuyXbONZzjrMR" +
		"0RsOmJPfyQTYPQuh6xdbcLP8bQiibREl+iOop0PLYRu9GtQ+r2uEI4VIuuNW1DwFzJRkWdxPm8kjC+g7XLDWqBYImFfep9Dfxj6Jam/w" +
		"ILKWVg3JjfuZe6nTwK7xqwDsY4Ylj8rnjnf+Cw3XQh94IU2E3uq3wqhoR1NCM7Qky+a6JIMGCALiqlj2DMnUgPKWEwFJ7bigRCdC+0s+" +
		"QQaeIEQ04Rm8jeuHaNvEsfloqLNrWx0D3ZHeIKB6hEegNnpSllfPJ/lF1U3ZnhMYd6oUef5HuV4wcu4oPf0nCEBi/CGlyRhmRKaondms" +
		"puRv5C4/zqQrqaVfJrkL0XO6EnhaxI0Yh6t5piA6vw73LtNH5fI55d1S2KWu0zUVbfyijzFohSfVZ5Zryo8uSggAgE96O+jiYzXfuS0d" +
		"vCVhXW/lw0ekQnKMx0Xh6U7XIckIvLgb4QQ6Oqv4/GZBo4dV7s78pNj/KvcCI6ya6qfrWgy1+pWmOLO+wcTmfzYOae4IyZnmuDXwyPng=\"," +
		"\"nonce\":\"3a584b49ed9b\"}," +
		"\"expire_time\":\"2031-04-25T16:55:23+08:00\",\"serial_no\":\"D7CE59D1F522D701\"}]}"
)

var mockWechatPayPrivateKey *rsa.PrivateKey
var mockWechatPayCertificate *x509.Certificate

func init() {
	var err error

	if mockWechatPayPrivateKey, err = utils.LoadPrivateKey(mockWechatPayPrivateKeyStr); err != nil {
		panic("mockWechatPayPrivateKeyStr is invalid")
	}
	if mockWechatPayCertificate, err = utils.LoadCertificate(mockWechatPayCertificateStr); err != nil {
		panic("mockWechatPayCertificateStr is invalid")
	}
}

func mockDownloadServer(t *testing.T) *gomonkey.Patches {
	patches := gomonkey.NewPatches()
	patches.ApplyMethod(
		reflect.TypeOf(&http.Client{}), "Do", func(_ *http.Client, req *http.Request) (*http.Response, error) {
			resp := http.Response{
				Status:        "200 OK",
				StatusCode:    200,
				Proto:         "HTTP/1.1",
				ProtoMajor:    1,
				ProtoMinor:    1,
				Header:        http.Header{},
				Body:          ioutil.NopCloser(bytes.NewBufferString(data)),
				ContentLength: int64(len(data)),
				Request:       req,
			}

			resp.Header.Set(consts.ContentLength, strconv.Itoa(len(data)))
			resp.Header.Set(consts.ContentType, "application/json; charset=utf-8")

			resp.Header.Set(consts.RequestID, "mock-request-id")
			resp.Header.Set(consts.WechatPaySerial, utils.GetCertificateSerialNumber(*mockWechatPayCertificate))
			resp.Header.Set(consts.WechatPayNonce, mockNonce)

			timestamp := strconv.FormatInt(time.Now().Unix(), 10)
			resp.Header.Set(consts.WechatPayTimestamp, timestamp)

			signature, err := utils.SignSHA256WithRSA(
				fmt.Sprintf("%s\n%s\n%s\n", timestamp, mockNonce, data), mockWechatPayPrivateKey,
			)
			require.NoError(t, err)

			resp.Header.Set(consts.WechatPaySignature, signature)

			return &resp, nil
		},
	)
	return patches
}
