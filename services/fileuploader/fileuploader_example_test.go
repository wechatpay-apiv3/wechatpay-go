// Copyright 2021 Tencent Inc. All rights reserved.

package fileuploader_test

import (
	"context"
	"log"
	"os"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/fileuploader"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

func ExampleImageUploader_Upload() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Print("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
	}

	file, err := os.Open("picture.jpg")
	if err != nil {
		return
	}
	defer file.Close()

	svc := fileuploader.ImageUploader{Client: client}
	resp, result, err := svc.Upload(ctx, file, "picture.jpg", consts.ImageJPG)

	if err != nil {
		// 处理错误
		log.Printf("call ImageUploader err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d media_id=%s", result.Response.StatusCode, *resp.MediaId)
	}
}

func ExampleVideoUploader_Upload() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Print("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
	}

	file, err := os.Open("video.mp4")
	if err != nil {
		return
	}
	defer file.Close()

	svc := fileuploader.VideoUploader{Client: client}
	resp, result, err := svc.Upload(ctx, file, "video.mp4", consts.VideoMP4)

	if err != nil {
		// 处理错误
		log.Printf("call VideoUploader err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d media_id=%s", result.Response.StatusCode, *resp.MediaId)
	}
}

func ExampleMarketingImageUploader_Upload() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Print("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
	}

	file, err := os.Open("picture.jpg")
	if err != nil {
		return
	}
	defer file.Close()

	svc := fileuploader.MarketingImageUploader{Client: client}
	resp, result, err := svc.Upload(ctx, file, "picture.jpg", consts.ImageJPG)

	if err != nil {
		// 处理错误
		log.Printf("call MarketingImageUploader err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d media_url=%s", result.Response.StatusCode, *resp.MediaUrl)
	}
}

func ExampleMchBizUploader_Upload() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Print("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
	}

	file, err := os.Open("picture.jpg")
	if err != nil {
		return
	}
	defer file.Close()

	svc := fileuploader.MchBizUploader{Client: client}
	resp, result, err := svc.Upload(ctx, file, "picture.jpg", consts.ImageJPG)

	if err != nil {
		// 处理错误
		log.Printf("call MchBizUploader err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d media_id=%s", result.Response.StatusCode, *resp.MediaId)
	}
}
