# 微信支付 API v3 Go SDK - wexinpayscoreparking

微信支付分停车服务 扣费API

## 总览
本 SDK 由 WechatPay APIv3 SDK 生成器生成。生成器基于 [OpenAPI Generator](https://openapi-generator.tech) 构建。

- API 版本: 1.2.1

## 接口列表

所有URI均基于微信支付 API 地址：*https://api.mch.weixin.qq.com*

服务名 | 方法名 | HTTP 请求 | 描述
------------ | ------------- | ------------- | -------------
*ParkingsApi* | [**CreateParking**](ParkingsApi.md#createparking) | **Post** /v3/vehicle/parking/parkings | 创建停车入场
*ServicesApi* | [**QueryPlateService**](ServicesApi.md#queryplateservice) | **Get** /v3/vehicle/parking/services/find | 查询车牌服务开通信息
*TransactionsApi* | [**CreateTransaction**](TransactionsApi.md#createtransaction) | **Post** /v3/vehicle/transactions/parking | 扣费受理
*TransactionsApi* | [**QueryTransaction**](TransactionsApi.md#querytransaction) | **Get** /v3/vehicle/transactions/out-trade-no/{out_trade_no} | 查询订单


## 类型列表

 - [CreateParkingRequest](CreateParkingRequest.md)
 - [CreateTransactionRequest](CreateTransactionRequest.md)
 - [OrderAmount](OrderAmount.md)
 - [Parking](Parking.md)
 - [ParkingTradeScene](ParkingTradeScene.md)
 - [Payer](Payer.md)
 - [PlateColor](PlateColor.md)
 - [PlateService](PlateService.md)
 - [PromotionDetail](PromotionDetail.md)
 - [QueryOrderAmount](QueryOrderAmount.md)
 - [QueryPlateServiceRequest](QueryPlateServiceRequest.md)
 - [QueryTransactionRequest](QueryTransactionRequest.md)
 - [Transaction](Transaction.md)

