# 微信支付 API v3 Go SDK - ecommerce/refund

境内电商退款功能相关API文档

## 总览
本 SDK 由 WechatPay APIv3 SDK 生成器生成。生成器基于 [OpenAPI Generator](https://openapi-generator.tech) 构建。

- API 版本: 1.1.8

## 接口列表

所有URI均基于微信支付 API 地址：*https://api.mch.weixin.qq.com*

服务名 | 方法名 | HTTP 请求 | 描述
------------ | ------------- | ------------- | -------------
*RefundsApi* | [**CreateRefund**](RefundsApi.md#createrefund) | **Post** /v3/ecommerce/refunds/apply | 申请退款
*RefundsApi* | [**CreateReturnAdvance**](RefundsApi.md#createreturnadvance) | **Post** /v3/ecommerce/refunds/{refund_id}/return-advance | 垫付退款回补
*RefundsApi* | [**QueryRefund**](RefundsApi.md#queryrefund) | **Get** /v3/ecommerce/refunds/id/{refund_id} | 查询单笔退款（按微信支付退款单号）
*RefundsApi* | [**QueryRefundByOutRefundNo**](RefundsApi.md#queryrefundbyoutrefundno) | **Get** /v3/ecommerce/refunds/out-refund-no/{out_refund_no} | 查询单笔退款（按商户退款单号）
*RefundsApi* | [**QueryReturnAdvance**](RefundsApi.md#queryreturnadvance) | **Get** /v3/ecommerce/refunds/{refund_id}/return-advance | 查询垫付回补结果


## 类型列表

 - [CreateAbnormalRefundRequest](CreateAbnormalRefundRequest.md)
 - [CreateRefundRequest](CreateRefundRequest.md)
 - [CreateReturnAdvanceBody](CreateReturnAdvanceBody.md)
 - [CreateReturnAdvanceRequest](CreateReturnAdvanceRequest.md)
 - [FundsFromItem](FundsFromItem.md)
 - [PromotionDetail](PromotionDetail.md)
 - [QueryRefundByOutRefundNoRequest](QueryRefundByOutRefundNoRequest.md)
 - [QueryRefundRequest](QueryRefundRequest.md)
 - [QueryReturnAdvanceRequest](QueryReturnAdvanceRequest.md)
 - [Refund](Refund.md)
 - [Refund4Create](Refund4Create.md)
 - [RefundAmount](RefundAmount.md)
 - [RefundReqAmount](RefundReqAmount.md)
 - [ReturnAdvance](ReturnAdvance.md)
 - [Type](Type.md)

