# 微信支付 API v3 Go SDK - profitsharing

微信支付分账API

## 总览
本 SDK 由 WechatPay APIv3 SDK 生成器生成。生成器基于 [OpenAPI Generator](https://openapi-generator.tech) 构建。

- API 版本: 0.0.3

## 接口列表

所有URI均基于微信支付 API 地址：*https://api.mch.weixin.qq.com*

服务名 | 方法名 | HTTP 请求 | 描述
------------ | ------------- | ------------- | -------------
*BillShipmentApi* | [**SplitBill**](BillShipmentApi.md#splitbill) | **Get** /v3/profitsharing/bills | 获取分账账单文件下载地址
*MerchantsApi* | [**QueryMerchantRatio**](MerchantsApi.md#querymerchantratio) | **Get** /v3/profitsharing/merchant-configs/{sub_mchid} | 查询最大分账比例API
*OrdersApi* | [**CreateOrder**](OrdersApi.md#createorder) | **Post** /v3/profitsharing/orders | 请求分账API
*OrdersApi* | [**QueryOrder**](OrdersApi.md#queryorder) | **Get** /v3/profitsharing/orders/{out_order_no} | 查询分账结果API
*OrdersApi* | [**UnfreezeOrder**](OrdersApi.md#unfreezeorder) | **Post** /v3/profitsharing/orders/unfreeze | 解冻剩余资金API
*ReceiversApi* | [**AddReceiver**](ReceiversApi.md#addreceiver) | **Post** /v3/profitsharing/receivers/add | 添加分账接收方API
*ReceiversApi* | [**DeleteReceiver**](ReceiversApi.md#deletereceiver) | **Post** /v3/profitsharing/receivers/delete | 删除分账接收方API
*ReturnOrdersApi* | [**CreateReturnOrder**](ReturnOrdersApi.md#createreturnorder) | **Post** /v3/profitsharing/return-orders | 请求分账回退API
*ReturnOrdersApi* | [**QueryReturnOrder**](ReturnOrdersApi.md#queryreturnorder) | **Get** /v3/profitsharing/return-orders/{out_return_no} | 查询分账回退结果API
*TransactionsApi* | [**QueryOrderAmount**](TransactionsApi.md#queryorderamount) | **Get** /v3/profitsharing/transactions/{transaction_id}/amounts | 查询剩余待分金额API


## 类型列表

 - [AddReceiverRequest](AddReceiverRequest.md)
 - [AddReceiverResponse](AddReceiverResponse.md)
 - [CreateOrderReceiver](CreateOrderReceiver.md)
 - [CreateOrderRequest](CreateOrderRequest.md)
 - [CreateReturnOrderRequest](CreateReturnOrderRequest.md)
 - [DeleteReceiverRequest](DeleteReceiverRequest.md)
 - [DeleteReceiverResponse](DeleteReceiverResponse.md)
 - [DetailFailReason](DetailFailReason.md)
 - [DetailStatus](DetailStatus.md)
 - [OrderReceiverDetail](OrderReceiverDetail.md)
 - [OrderStatus](OrderStatus.md)
 - [OrdersEntity](OrdersEntity.md)
 - [QueryMerchantRatioRequest](QueryMerchantRatioRequest.md)
 - [QueryMerchantRatioResponse](QueryMerchantRatioResponse.md)
 - [QueryOrderAmountRequest](QueryOrderAmountRequest.md)
 - [QueryOrderAmountResponse](QueryOrderAmountResponse.md)
 - [QueryOrderRequest](QueryOrderRequest.md)
 - [QueryReturnOrderRequest](QueryReturnOrderRequest.md)
 - [ReceiverRelationType](ReceiverRelationType.md)
 - [ReceiverType](ReceiverType.md)
 - [ReturnOrderFailReason](ReturnOrderFailReason.md)
 - [ReturnOrderStatus](ReturnOrderStatus.md)
 - [ReturnOrdersEntity](ReturnOrdersEntity.md)
 - [SplitBillHashType](SplitBillHashType.md)
 - [SplitBillRequest](SplitBillRequest.md)
 - [SplitBillResponse](SplitBillResponse.md)
 - [SplitBillTarType](SplitBillTarType.md)
 - [UnfreezeOrderRequest](UnfreezeOrderRequest.md)

