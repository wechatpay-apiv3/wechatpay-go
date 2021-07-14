// Copyright 2021 Tencent Inc. All rights reserved.

package core_test

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cipher"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
)

func ExampleNewClient_default() {
	// 示例参数，实际使用时请自行初始化
	var (
		mchID                      string
		mchCertificateSerialNumber string
		mchPrivateKey              *rsa.PrivateKey
		wechatPayCertList          []*x509.Certificate
		customHTTPClient           *http.Client
	)

	client, err := core.NewClient(
		context.Background(),
		// 一次性设置 签名/验签/敏感字段加解密，并注册 平台证书下载器，自动定时获取最新的平台证书
		option.WithWechatPayAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, wechatPayCertList),
		// 设置自定义 HTTPClient 实例，不设置时默认使用 http.Client{}，并设置超时时间为 30s
		option.WithHTTPClient(customHTTPClient),
	)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err.Error())
		return
	}
	// 接下来使用 client 进行请求发送
	_ = client
}

func ExampleNewClient_auto_update_certificate() {
	// 示例参数，实际使用时请自行初始化
	var (
		mchID                      string
		mchCertificateSerialNumber string
		mchPrivateKey              *rsa.PrivateKey
		mchAPIv3Key                string
	)

	client, err := core.NewClient(
		context.Background(),
		// 一次性设置 签名/验签/敏感字段加解密，并注册 平台证书下载器，自动定时获取最新的平台证书
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err.Error())
		return
	}
	// 接下来使用 client 进行请求发送
	_ = client
}

func ExampleNewClient_fully_customized() {
	var (
		signer           auth.Signer      // 自定义实现 auth.Signer 接口的实例
		verifier         auth.Verifier    // 自定义实现 auth.Verifier 接口的实例
		encryptor        cipher.Encryptor // 自定义实现 auth.Encryptor 接口的实例
		decryptor        cipher.Decryptor // 自定义实现 cipher.Decryptor 接口的实例
		customHTTPClient *http.Client     // 自定义 HTTPClient
	)

	client, err := core.NewClient(
		context.Background(),
		// 使用自定义 Signer 初始化 微信支付签名器
		option.WithSigner(signer),
		// 使用自定义 Verifier 初始化 微信支付应答验证器
		option.WithVerifier(verifier),
		// 使用自定义 Encryptor/Decryptor 初始化 微信支付敏感字段加解密器
		option.WithWechatPayCipher(encryptor, decryptor),
		// 使用自定义 HTTPClient
		option.WithHTTPClient(customHTTPClient),
	)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err.Error())
		return
	}
	// 接下来使用 client 进行请求发送
	_ = client
}

func ExampleCreateFormField() {
	var w multipart.Writer

	meta := map[string]string{
		"filename": "sample.jpg",
		"sha256":   "5944758444f0af3bc843e39b611a6b0c8c38cca44af653cd461b5765b71dc3f8",
	}

	metaBytes, err := json.Marshal(meta)
	if err != nil {
		// TODO: 处理错误
		return
	}

	err = core.CreateFormField(&w, "meta", consts.ApplicationJSON, metaBytes)
	if err != nil {
		// TODO: 处理错误
	}
}

func ExampleCreateFormFile() {
	var w multipart.Writer

	var fileContent []byte

	err := core.CreateFormFile(&w, "sample.jpg", consts.ImageJPG, fileContent)
	if err != nil {
		// TODO: 处理错误
	}
}
