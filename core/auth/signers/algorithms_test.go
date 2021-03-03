package signers

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"testing"
)

const (
	testAlgorithmPrivateKeyStr = `-----BEGIN PRIVATE KEY-----
-----END PRIVATE KEY-----`
	testAlgorithmExpectSignature = ""
)

var testAlgorithmPrivateKey *rsa.PrivateKey
var ok bool

func init() {
	block, _ := pem.Decode([]byte(testAlgorithmPrivateKeyStr))
	if block == nil {
		panic("decode private key err")
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		panic(fmt.Errorf("parse private key err:%s", err.Error()))
	}
	testAlgorithmPrivateKey, ok = key.(*rsa.PrivateKey)
	if !ok {
		panic(fmt.Errorf("%s is not rsa private key", testAlgorithmPrivateKeyStr))
	}
}

func TestSha256WithRsa1(t *testing.T) {
	type args struct {
		source     string
		privateKey *rsa.PrivateKey
	}
	tests := []struct {
		name          string
		args          args
		wantSignature string
		wantErr       bool
	}{
		{
			name: "sign message in sha256withrsa success",
			args: args{
				source:     "source",
				privateKey: testAlgorithmPrivateKey,
			},
			wantSignature: testAlgorithmExpectSignature,
			wantErr:       false,
		},
		{
			name: "sign message in sha256withrsa err",
			args: args{
				source:     "source",
				privateKey: nil,
			},
			wantSignature: "",
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSignature, err := Sha256WithRsa(tt.args.source, tt.args.privateKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sha256WithRsa() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSignature != tt.wantSignature {
				t.Errorf("Sha256WithRsa() gotSignature = %v, want %v", gotSignature, tt.wantSignature)
			}
		})
	}
}
