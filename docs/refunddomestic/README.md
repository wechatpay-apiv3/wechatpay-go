# 微信支付 API v3 Go SDK - refunddomestic

境内普通商户退款功能涉及的API文档


## 总览
本 SDK 由 WechatPay APIv3 SDK 生成器生成。生成器基于 [OpenAPI Generator](https://openapi-generator.tech) 构建。

- API 版本: 1.1.1

## 接口列表

所有URI均基于微信支付 API 地址：*https://api.mch.weixin.qq.com*

服务名 | 方法名 | HTTP 请求 | 描述
------------ | ------------- | ------------- | -------------
*RefundsApi* | [**Create**](RefundsApi.md#create) | **Post** /v3/refund/domestic/refunds | 退款申请
*RefundsApi* | [**QueryByOutRefundNo**](RefundsApi.md#querybyoutrefundno) | **Get** /v3/refund/domestic/refunds/{out_refund_no} | 查询单笔退款（通过商户退款单号）


## 类型列表

 - [Account](Account.md)
 - [Amount](Amount.md)
 - [AmountReq](AmountReq.md)
 - [Channel](Channel.md)
 - [CreateRequest](CreateRequest.md)
 - [FundsAccount](FundsAccount.md)
 - [FundsFromItem](FundsFromItem.md)
 - [GoodsDetail](GoodsDetail.md)
 - [Promotion](Promotion.md)
 - [QueryByOutRefundNoRequest](QueryByOutRefundNoRequest.md)
 - [Refund](Refund.md)
 - [ReqFundsAccount](ReqFundsAccount.md)
 - [Scope](Scope.md)
 - [Status](Status.md)
 - [Type](Type.md)

