// Copyright 2021 Tencent Inc. All rights reserved.

// Package consts 微信支付 API v3 Go SDK 常量
package consts

import "time"

// 微信支付 API 地址
const (
	WechatPayAPIServer       = "https://api.mch.weixin.qq.com"  // 微信支付 API 地址
	WechatPayAPIServerBackup = "https://api2.mch.weixin.qq.com" // 微信支付 API 备份地址
)

// SDK 相关信息
const (
	Version         = "0.2.5"                      // SDK 版本
	UserAgentFormat = "WechatPay-Go/%s (%s) GO/%s" // UserAgent中的信息
)

// HTTP 请求报文 Header 相关常量
const (
	Authorization = "Authorization"  // Header 中的 Authorization 字段
	Accept        = "Accept"         // Header 中的 Accept 字段
	ContentType   = "Content-Type"   // Header 中的 ContentType 字段
	ContentLength = "Content-Length" // Header 中的 ContentLength 字段
	UserAgent     = "User-Agent"     // Header 中的 UserAgent 字段
)

// 常用 ContentType
const (
	ApplicationJSON = "application/json"
	ImageJPG        = "image/jpg"
	ImagePNG        = "image/png"
	VideoMP4        = "video/mp4"
)

// 请求报文签名相关常量
const (
	SignatureMessageFormat = "%s\n%s\n%d\n%s\n%s\n" // 数字签名原文格式
	// HeaderAuthorizationFormat 请求头中的 Authorization 拼接格式
	HeaderAuthorizationFormat = "%s mchid=\"%s\",nonce_str=\"%s\",timestamp=\"%d\",serial_no=\"%s\",signature=\"%s\""
)

// HTTP 应答报文 Header 相关常量
const (
	WechatPayTimestamp = "Wechatpay-Timestamp" // 微信支付回包时间戳
	WechatPayNonce     = "Wechatpay-Nonce"     // 微信支付回包随机字符串
	WechatPaySignature = "Wechatpay-Signature" // 微信支付回包签名信息
	WechatPaySerial    = "Wechatpay-Serial"    // 微信支付回包平台序列号
	RequestID          = "Request-Id"          // 微信支付回包请求ID
)

// 时间相关常量
const (
	FiveMinute     = 5 * 60           // 回包校验最长时间（秒）
	DefaultTimeout = 30 * time.Second // HTTP 请求默认超时时间
)
