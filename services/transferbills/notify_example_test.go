// Copyright 2021 Tencent Inc. All rights reserved.

package transferbills_test

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/transferbills"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

func ExampleTransferNotifyContent_handleTransferNotify() {
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
	_, err = core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	// 获取商户的 RSA 验签器
	certVisitor := downloader.MgrInstance().GetCertificateVisitor(mchID)
	verifier := verifiers.NewSHA256WithRSAVerifier(certVisitor)

	// 创建通知处理器
	handler, err := notify.NewRSANotifyHandler(mchAPIv3Key, verifier)
	if err != nil {
		log.Printf("create notify handler err:%s", err)
		return
	}

	// 模拟接收到转账通知的HTTP请求
	request := &http.Request{}

	// 解析转账通知
	content := new(transferbills.TransferNotifyContent)
	notifyReq, err := handler.ParseNotifyRequest(ctx, request, content)
	if err != nil {
		log.Printf("parse notify request err:%s", err)
		return
	}

	// 验证事件类型是否为转账通知
	if notifyReq.EventType != transferbills.TransferNotifyEventType {
		log.Printf("unexpected event type: %s", notifyReq.EventType)
		return
	}

	// 处理转账通知
	switch *content.State {
	case transferbills.TransferNotifyStateSuccess:
		log.Printf("转账成功通知: 商户订单号=%s, 微信转账单号=%s, 转账金额=%d分",
			*content.OutBillNo, *content.TransferBillNo, *content.TransferAmount)
		// 更新业务系统中的转账状态为成功
		// updateTransferStatus(content.OutBillNo, "SUCCESS")

	case transferbills.TransferNotifyStateFail:
		log.Printf("转账失败通知: 商户订单号=%s, 失败原因=%s",
			*content.OutBillNo, *content.FailReason)
		// 更新业务系统中的转账状态为失败
		// updateTransferStatus(content.OutBillNo, "FAIL")

	case transferbills.TransferNotifyStateCancelled:
		log.Printf("转账已撤销通知: 商户订单号=%s",
			*content.OutBillNo)
		// 更新业务系统中的转账状态为已撤销
		// updateTransferStatus(content.OutBillNo, "CANCELLED")

	default:
		log.Printf("未知转账状态: %s", *content.State)
	}

	// 通知处理完成后，需要返回成功状态给微信
	// 在实际的HTTP处理函数中，应该返回200状态码
	log.Printf("转账通知处理完成")
}

func ExampleTransferNotifyContent_handleWithGeneralContent() {
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
	_, err = core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	// 获取商户的 RSA 验签器
	certVisitor := downloader.MgrInstance().GetCertificateVisitor(mchID)
	verifier := verifiers.NewSHA256WithRSAVerifier(certVisitor)

	// 创建通知处理器
	handler, err := notify.NewRSANotifyHandler(mchAPIv3Key, verifier)
	if err != nil {
		log.Printf("create notify handler err:%s", err)
		return
	}

	// 模拟接收到转账通知的HTTP请求
	request := &http.Request{}

	// 使用通用的map解析通知内容（适用于SDK未明确支持的通知类型）
	content := make(map[string]interface{})
	notifyReq, err := handler.ParseNotifyRequest(ctx, request, &content)
	if err != nil {
		log.Printf("parse notify request err:%s", err)
		return
	}

	// 验证事件类型是否为转账通知
	if notifyReq.EventType == transferbills.TransferNotifyEventType {
		// 从通用map中提取转账相关信息
		outBillNo := content["out_bill_no"].(string)
		state := content["state"].(string)
		transferAmount := int64(content["transfer_amount"].(float64))

		log.Printf("转账通知: 商户订单号=%s, 状态=%s, 金额=%d分",
			outBillNo, state, transferAmount)

		// 根据状态处理业务逻辑
		switch state {
		case transferbills.TransferNotifyStateSuccess:
			fmt.Printf("处理转账成功逻辑")
		case transferbills.TransferNotifyStateFail:
			fmt.Printf("处理转账失败逻辑")
		case transferbills.TransferNotifyStateCancelled:
			fmt.Printf("处理转账撤销逻辑")
		}
	}

	log.Printf("通知处理完成")
}

func ExampleTransferNotifyContent_httpHandler() {
	var (
		mchAPIv3Key string = "2ab9****************************" // 商户APIv3密钥
	)

	// 获取验签器（实际使用中需要正确初始化）
	var verifier *verifiers.SHA256WithRSAVerifier

	// 创建通知处理器
	handler, err := notify.NewRSANotifyHandler(mchAPIv3Key, verifier)
	if err != nil {
		log.Printf("create notify handler err:%s", err)
		return
	}

	// 定义HTTP处理函数
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		// 解析转账通知
		content := new(transferbills.TransferNotifyContent)
		notifyReq, err := handler.ParseNotifyRequest(context.Background(), r, content)
		if err != nil {
			log.Printf("parse notify request failed: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"code":"FAIL","message":"验签失败"}`))
			return
		}

		// 验证事件类型
		if notifyReq.EventType != transferbills.TransferNotifyEventType {
			log.Printf("unexpected event type: %s", notifyReq.EventType)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"code":"FAIL","message":"事件类型不匹配"}`))
			return
		}

		// 处理业务逻辑（这里应该是幂等的操作）
		err = processTransferNotify(content)
		if err != nil {
			log.Printf("process transfer notify failed: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"code":"FAIL","message":"处理失败"}`))
			return
		}

		// 返回成功响应
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"code":"SUCCESS","message":"成功"}`))
	}

	// 注册HTTP路由
	http.HandleFunc("/notify/transfer", handleFunc)

	log.Printf("转账通知处理器已注册到 /notify/transfer")
}

// processTransferNotify 处理转账通知的业务逻辑
func processTransferNotify(content *transferbills.TransferNotifyContent) error {
	// 这里应该实现具体的业务处理逻辑
	// 注意：由于微信可能会重复发送通知，这里的处理必须是幂等的

	switch *content.State {
	case transferbills.TransferNotifyStateSuccess:
		// 处理转账成功的业务逻辑
		log.Printf("处理转账成功: %s", *content.OutBillNo)
		
	case transferbills.TransferNotifyStateFail:
		// 处理转账失败的业务逻辑
		log.Printf("处理转账失败: %s, 原因: %s", *content.OutBillNo, *content.FailReason)
		
	case transferbills.TransferNotifyStateCancelled:
		// 处理转账撤销的业务逻辑
		log.Printf("处理转账撤销: %s", *content.OutBillNo)
	}

	return nil
}