# 微信支付 API v3 Go SDK - partnerpayments/native

Native支付API

## 总览
本 SDK 由 WechatPay APIv3 SDK 生成器生成。生成器基于 [OpenAPI Generator](https://openapi-generator.tech) 构建。

- API 版本: 1.2.3

想获取更多信息，请访问 [https://pay.weixin.qq.com/wiki/doc/apiv3/index.shtml](https://pay.weixin.qq.com/wiki/doc/apiv3/index.shtml)

## 接口列表

所有URI均基于微信支付 API 地址：*https://api.mch.weixin.qq.com*

服务名 | 方法名 | HTTP 请求 | 描述
------------ | ------------- | ------------- | -------------
*NativeApi* | [**CloseOrder**](NativeApi.md#closeorder) | **Post** /v3/pay/partner/transactions/out-trade-no/{out_trade_no}/close | 关闭订单
*NativeApi* | [**Prepay**](NativeApi.md#prepay) | **Post** /v3/pay/partner/transactions/native | Native支付预下单
*NativeApi* | [**QueryOrderById**](NativeApi.md#queryorderbyid) | **Get** /v3/pay/partner/transactions/id/{transaction_id} | 微信支付订单号查询订单
*NativeApi* | [**QueryOrderByOutTradeNo**](NativeApi.md#queryorderbyouttradeno) | **Get** /v3/pay/partner/transactions/out-trade-no/{out_trade_no} | 商户订单号查询订单


## 类型列表

 - [Amount](Amount.md)
 - [CloseOrderRequest](CloseOrderRequest.md)
 - [CloseRequest](CloseRequest.md)
 - [Detail](Detail.md)
 - [GoodsDetail](GoodsDetail.md)
 - [PrepayRequest](PrepayRequest.md)
 - [PrepayResponse](PrepayResponse.md)
 - [QueryOrderByIdRequest](QueryOrderByIdRequest.md)
 - [QueryOrderByOutTradeNoRequest](QueryOrderByOutTradeNoRequest.md)
 - [SceneInfo](SceneInfo.md)
 - [SettleInfo](SettleInfo.md)
 - [StoreInfo](StoreInfo.md)

