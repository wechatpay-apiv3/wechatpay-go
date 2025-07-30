// Copyright 2021 Tencent Inc. All rights reserved.

package transferbills_test

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/transferbills"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

func ExampleTransferBillsApiService_CreateTransferBill() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
		mchPrivateKeyPath          string = "path/to/merchant/apiclient_key.pem"       // 商户私钥文件路径
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(mchPrivateKeyPath)
	if err != nil {
		log.Printf("load merchant private key error:%s", err)
		return
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	svc := transferbills.TransferBillsApiService{Client: client}
	resp, result, err := svc.CreateTransferBill(ctx,
		transferbills.CreateTransferBillRequest{
			Appid:           core.String("wxf636efh567hg4356"),
			OutBillNo:       core.String("plfk2020042013"),
			TransferSceneId: core.String("1001"),
			Openid:          core.String("o-MYE42l80oelYMDE34nYD456Xoy"),
			TransferAmount:  core.Int64(100),
			TransferRemark:  core.String("理财返佣"),
			TransferSceneReportInfos: []transferbills.TransferSceneReportInfo{
				{
					InfoType:    core.String("订单信息"),
					InfoContent: core.String("理财产品收益"),
				},
			},
			UserName:           core.String("金三"),
			NotifyUrl:          core.String("https://www.example.com/notify"),
			UserRecvPerception: core.String("理财收益到账"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call CreateTransferBill err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleTransferBillsApiService_GetTransferBillByOutBillNo() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
		mchPrivateKeyPath          string = "path/to/merchant/apiclient_key.pem"       // 商户私钥文件路径
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(mchPrivateKeyPath)
	if err != nil {
		log.Printf("load merchant private key error:%s", err)
		return
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	svc := transferbills.TransferBillsApiService{Client: client}
	resp, result, err := svc.GetTransferBillByOutBillNo(ctx,
		transferbills.GetTransferBillByOutBillNoRequest{
			OutBillNo: core.String("plfk2020042013"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call GetTransferBillByOutBillNo err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleTransferBillsApiService_CancelTransferBill() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
		mchPrivateKeyPath          string = "path/to/merchant/apiclient_key.pem"       // 商户私钥文件路径
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(mchPrivateKeyPath)
	if err != nil {
		log.Printf("load merchant private key error:%s", err)
		return
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	svc := transferbills.TransferBillsApiService{Client: client}
	resp, result, err := svc.CancelTransferBill(ctx,
		transferbills.CancelTransferBillRequest{
			OutBillNo: core.String("plfk2020042013"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call CancelTransferBill err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleTransferBillsApiService_GetTransferBillByTransferBillNo() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
		mchPrivateKeyPath          string = "path/to/merchant/apiclient_key.pem"       // 商户私钥文件路径
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(mchPrivateKeyPath)
	if err != nil {
		log.Printf("load merchant private key error:%s", err)
		return
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	svc := transferbills.TransferBillsApiService{Client: client}
	resp, result, err := svc.GetTransferBillByTransferBillNo(ctx,
		transferbills.GetTransferBillByTransferBillNoRequest{
			TransferBillNo: core.String("1000000000001"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call GetTransferBillByTransferBillNo err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}
