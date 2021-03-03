package utils

import (
	"crypto/x509"
	"reflect"
	"testing"
)

var (
	testPemUtilCertificatePath = ""
	testPemUtilCertificateStr  = ""
	testPemUtilPrivateKeyPath  = ""
	testPemUtilPrivateKeyStr   = ""
	testPemUtilPublicKeyPath   = ""
	testPemUtilPublicKeyStr    = ""
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
			name:    "load certificate from str to x509.certificate success",
			args:    args{certificateStr: testPemUtilCertificateStr},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := LoadCertificate(tt.args.certificateStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadCertificate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestLoadCertificateWithPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name            string
		args            args
		wantCertificate *x509.Certificate
		wantErr         bool
	}{
		{
			name:    "load certificate with from path to x509.certificate success",
			args:    args{path: testPemUtilCertificatePath},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCertificate, err := LoadCertificateWithPath(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadCertificateWithPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCertificate, tt.wantCertificate) {
				t.Errorf("LoadCertificateWithPath() gotCertificate = %v, want %v", gotCertificate, tt.wantCertificate)
			}
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
			name:    "load private key from str to rsa.PrivateKey success",
			args:    args{privateKeyStr: testPemUtilPrivateKeyStr},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := LoadPrivateKey(tt.args.privateKeyStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadPrivateKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestLoadPrivateKeyWithPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args

		wantErr bool
	}{
		{
			name:    "load private key from path to *rsa.PrivateKey success",
			args:    args{path: testPemUtilPrivateKeyPath},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := LoadPrivateKeyWithPath(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadPrivateKeyWithPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
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
			name:    "load private key from str to rsa.PrivateKey success",
			args:    args{publicKeyStr: testPemUtilPublicKeyStr},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := LoadPublicKey(tt.args.publicKeyStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadPublicKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestLoadPublicKeyWithPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "load public key from  path to *rsa.PublicKey success",
			args:    args{path: testPemUtilPublicKeyPath},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := LoadPublicKeyWithPath(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadPublicKeyWithPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
