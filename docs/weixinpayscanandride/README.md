# 微信支付 API v3 Go SDK - weixinpayscanandride

公共出行平台代扣服务对外API

## 总览
本 SDK 由 WechatPay APIv3 SDK 生成器生成。生成器基于 [OpenAPI Generator](https://openapi-generator.tech) 构建。

- API 版本: 1.1.0

## 接口列表

所有URI均基于微信支付 API 地址：*https://api.mch.weixin.qq.com*

服务名 | 方法名 | HTTP 请求 | 描述
------------ | ------------- | ------------- | -------------
*PalmPayTransactionsApi* | [**CreatePalmPayTransaction**](PalmPayTransactionsApi.md#createpalmpaytransaction) | **Post** /v3/transit-palm-pay/transactions | 直连商户扣费受理—刷掌地铁
*PalmPayTransactionsApi* | [**QueryPalmPayTransaction**](PalmPayTransactionsApi.md#querypalmpaytransaction) | **Get** /v3/transit-palm-pay/transactions/{out_trade_no} | 直连商户查询订单—刷掌地铁
*TransactionsApi* | [**CreateTransaction**](TransactionsApi.md#createtransaction) | **Post** /v3/qrcode/transactions | 扣费受理
*TransactionsApi* | [**QueryTransaction**](TransactionsApi.md#querytransaction) | **Get** /v3/qrcode/transactions/out-trade-no/{out_trade_no} | 查询订单
*UserServicesApi* | [**QueryUserService**](UserServicesApi.md#queryuserservice) | **Get** /v3/qrcode/user-services/contract-id/{contract_id} | 查询用户服务可用信息


## 类型列表

 - [BlockReasonEnum](BlockReasonEnum.md)
 - [BusSceneInfo](BusSceneInfo.md)
 - [CreatePalmPayTransactionRequest](CreatePalmPayTransactionRequest.md)
 - [CreateTransactionRequest](CreateTransactionRequest.md)
 - [MetroSceneInfo](MetroSceneInfo.md)
 - [OrderAmount](OrderAmount.md)
 - [PalmMetroSceneInfo](PalmMetroSceneInfo.md)
 - [PalmPayTradeScene](PalmPayTradeScene.md)
 - [PalmPayTransaction](PalmPayTransaction.md)
 - [PromotionDetail](PromotionDetail.md)
 - [PromotionType](PromotionType.md)
 - [QueryOrderAmount](QueryOrderAmount.md)
 - [QueryPalmPayTransactionRequest](QueryPalmPayTransactionRequest.md)
 - [QueryTransactionRequest](QueryTransactionRequest.md)
 - [QueryUserServiceRequest](QueryUserServiceRequest.md)
 - [TradeScene](TradeScene.md)
 - [TradeState](TradeState.md)
 - [TransactionsEntity](TransactionsEntity.md)
 - [UserRepayState](UserRepayState.md)
 - [UserServiceEntity](UserServiceEntity.md)
 - [UserServiceStateEnum](UserServiceStateEnum.md)

