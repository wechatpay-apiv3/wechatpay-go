// Copyright 2021 Tencent Inc. All rights reserved.
//
// 微信支付营销系统开放API
//
// 新增立减金api
//
// API version: 3.4.0

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package cashcoupons_test

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/cashcoupons"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

func ExamplePlatformFavorApiService_CreatePlatformFavor() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
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

	svc := cashcoupons.PlatformFavorApiService{Client: client}
	resp, result, err := svc.CreatePlatformFavor(ctx,
		cashcoupons.CreatePlatformFavorRequest{
			BelongMerchant: core.String("98568865"),
			SendRule: &cashcoupons.PlatformFavorSendRule{
				MaxCoupons:          core.Int64(9000),
				MaxAmount:           core.Int64(10000),
				MaxCouponsByDay:     core.Int64(100),
				MaxCouponsPerUser:   core.Int64(3),
				NaturalPersonLimit:  core.Bool(true),
				PreventApiAbuse:     core.Bool(false),
				DeductBalanceMethod: cashcoupons.DEDUCTBALANCEMETHOD_BATCH_DEDUCT.Ptr(),
			},
			UseRule: &cashcoupons.PlatformFavorUseRule{
				AvailableTime: &cashcoupons.PlatformFavorAvailableTime{
					AvailableBeginTime: core.String("2015-05-20T13:29:35.120+08:00"),
					AvailableEndTime:   core.String("2015-05-20T13:29:35.120+08:00"),
				},
				CouponAmount:       core.Int64(15),
				TransactionMinimum: core.Int64(100),
			},
			DisplayPattern: &cashcoupons.PlatformFavorDisplayPattern{
				Description: core.String("使用微信支付消费时自动抵扣。单笔消费满xx元以上（含xx元）可一次性全额使用立减金，不找零，不重复使用。生活缴费、理财等少数特定商户不可使用，有问题请致电95017"),
			},
			OutRequestNo: core.String("mch_create_202003080999"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call CreatePlatformFavor err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExamplePlatformFavorApiService_GetPlatformFavor() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
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

	svc := cashcoupons.PlatformFavorApiService{Client: client}
	resp, result, err := svc.GetPlatformFavor(ctx,
		cashcoupons.GetPlatformFavorRequest{
			StockId:           core.String("1263456"),
			StockCreatorMchid: core.String("987654"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call GetPlatformFavor err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}