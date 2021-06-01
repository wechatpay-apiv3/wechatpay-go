package ciphers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cipher/decryptors"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cipher/encryptors"
)

type Student struct {
	Name      string `encryption:"EM_APIV3"`
	Age       int
	Addresses []Address
	Parents   *[]Parent
}

type Address struct {
	// No Tag
	Country *string
	// Not EM_APIV3 encryption Tag
	Province *string `encryption:"EM_APIV2"`
	// EM_APIV3 encryption Tag
	City   *string `encryption:"EM_APIV3"`
	Street *string `encryption:"EM_APIV3"`
}

type Parent struct {
	Name        string  `encryption:"EM_APIV3"`
	PhoneNumber *string `encryption:"EM_APIV3"`
}

func TestWechatPayCipher_Encrypt_Decrypt(t *testing.T) {
	s := Student{
		Name: "小可",
		Age:  8,
		Addresses: []Address{
			{
				Country:  core.String("中国"),
				Province: core.String("四川"),
				City:     core.String("成都"),
				Street:   core.String("春熙路"),
			},
			{
				Country:  core.String("USA"),
				Province: core.String("California"),
				City:     core.String("LA"),
				Street:   core.String("Nowhere"),
			},
		},
		Parents: &[]Parent{
			{
				Name:        "爸",
				PhoneNumber: core.String("13000000000"),
			},
			{
				Name:        "妈",
				PhoneNumber: nil,
			},
		},
	}

	c := WechatPayCipher{
		encryptor: &encryptors.MockEncryptor{
			Serial: "Mock Serial",
		},
		decryptor: &decryptors.MockDecryptor{},
	}

	serial, err := c.Encrypt(context.Background(), &s)
	assert.Equal(t, "Mock Serial", serial)
	assert.Nil(t, err)
	assert.Equal(t, "Encrypted小可", s.Name)
	assert.Equal(t, 8, s.Age)
	assert.Equal(t, "中国", *(s.Addresses[0].Country))
	assert.Equal(t, "四川", *(s.Addresses[0].Province))
	assert.Equal(t, "Encrypted成都", *(s.Addresses[0].City))
	assert.Equal(t, "Encrypted春熙路", *(s.Addresses[0].Street))
	assert.Equal(t, "USA", *(s.Addresses[1].Country))
	assert.Equal(t, "California", *(s.Addresses[1].Province))
	assert.Equal(t, "EncryptedLA", *(s.Addresses[1].City))
	assert.Equal(t, "EncryptedNowhere", *(s.Addresses[1].Street))
	assert.Equal(t, "Encrypted爸", (*s.Parents)[0].Name)
	assert.Equal(t, "Encrypted13000000000", *((*s.Parents)[0].PhoneNumber))
	assert.Equal(t, "Encrypted妈", (*s.Parents)[1].Name)
	assert.Equal(t, (*string)(nil), (*s.Parents)[1].PhoneNumber)

	err = c.Decrypt(context.Background(), &s)
	assert.Nil(t, err)
	assert.Equal(t, "小可", s.Name)
	assert.Equal(t, 8, s.Age)
	assert.Equal(t, "中国", *(s.Addresses[0].Country))
	assert.Equal(t, "四川", *(s.Addresses[0].Province))
	assert.Equal(t, "成都", *(s.Addresses[0].City))
	assert.Equal(t, "春熙路", *(s.Addresses[0].Street))
	assert.Equal(t, "USA", *(s.Addresses[1].Country))
	assert.Equal(t, "California", *(s.Addresses[1].Province))
	assert.Equal(t, "LA", *(s.Addresses[1].City))
	assert.Equal(t, "Nowhere", *(s.Addresses[1].Street))
	assert.Equal(t, "爸", (*s.Parents)[0].Name)
	assert.Equal(t, "13000000000", *((*s.Parents)[0].PhoneNumber))
	assert.Equal(t, "妈", (*s.Parents)[1].Name)
	assert.Equal(t, (*string)(nil), (*s.Parents)[1].PhoneNumber)
}

func TestWechatPayCipher_CipherValue(t *testing.T) {
	s := Student{
		Name: "小可",
		Age:  8,
	}

	c := WechatPayCipher{
		encryptor: &encryptors.MockEncryptor{
			Serial: "Mock Serial",
		},
		decryptor: &decryptors.MockDecryptor{},
	}

	serial, err := c.Encrypt(context.Background(), s)
	assert.Equal(t, "Mock Serial", serial)
	assert.NotNil(t, err)
	assert.Equal(t, "in-place cipher requires settable input, ptr for example", err.Error())
}
