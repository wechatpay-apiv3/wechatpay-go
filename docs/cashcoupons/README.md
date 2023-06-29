# 微信支付 API v3 Go SDK - cashcoupons

新增立减金api

## 总览
本 SDK 由 WechatPay APIv3 SDK 生成器生成。生成器基于 [OpenAPI Generator](https://openapi-generator.tech) 构建。

- API 版本: 3.4.0

## 接口列表

所有URI均基于微信支付 API 地址：*https://api.mch.weixin.qq.com*

服务名 | 方法名 | HTTP 请求 | 描述
------------ | ------------- | ------------- | -------------
*CallBackUrlApi* | [**QueryCallback**](CallBackUrlApi.md#querycallback) | **Get** /v3/marketing/favor/callbacks | 查询代金券消息通知地址
*CallBackUrlApi* | [**SetCallback**](CallBackUrlApi.md#setcallback) | **Post** /v3/marketing/favor/callbacks | 设置代金券消息通知地址
*CouponApi* | [**ListCouponsByFilter**](CouponApi.md#listcouponsbyfilter) | **Get** /v3/marketing/favor/users/{openid}/coupons | 根据过滤条件查询用户的券
*CouponApi* | [**QueryCoupon**](CouponApi.md#querycoupon) | **Get** /v3/marketing/favor/users/{openid}/coupons/{coupon_id} | 查询代金券详情
*CouponApi* | [**SendCoupon**](CouponApi.md#sendcoupon) | **Post** /v3/marketing/favor/users/{openid}/coupons | 发放指定批次的代金券
*StockApi* | [**CreateCouponStock**](StockApi.md#createcouponstock) | **Post** /v3/marketing/favor/coupon-stocks | 创建代金券批次
*StockApi* | [**ListAvailableMerchants**](StockApi.md#listavailablemerchants) | **Get** /v3/marketing/favor/stocks/{stock_id}/merchants | 查询代金券可用商户
*StockApi* | [**ListAvailableSingleitems**](StockApi.md#listavailablesingleitems) | **Get** /v3/marketing/favor/stocks/{stock_id}/items | 查询可核销商品编码
*StockApi* | [**ListStocks**](StockApi.md#liststocks) | **Get** /v3/marketing/favor/stocks | 条件查询批次列表
*StockApi* | [**PauseStock**](StockApi.md#pausestock) | **Post** /v3/marketing/favor/stocks/{stock_id}/pause | 暂停批次
*StockApi* | [**QueryStock**](StockApi.md#querystock) | **Get** /v3/marketing/favor/stocks/{stock_id} | 查询批次详情
*StockApi* | [**RefundFlow**](StockApi.md#refundflow) | **Get** /v3/marketing/favor/stocks/{stock_id}/refund-flow | 下载批次退款明细
*StockApi* | [**RestartStock**](StockApi.md#restartstock) | **Post** /v3/marketing/favor/stocks/{stock_id}/restart | 重启批次
*StockApi* | [**StartStock**](StockApi.md#startstock) | **Post** /v3/marketing/favor/stocks/{stock_id}/start | 激活开启批次
*StockApi* | [**StopStock**](StockApi.md#stopstock) | **Post** /v3/marketing/favor/stocks/{stock_id}/stop | 终止批次
*StockApi* | [**UseFlow**](StockApi.md#useflow) | **Get** /v3/marketing/favor/stocks/{stock_id}/use-flow | 下载批次核销明细


## 类型列表

 - [AvailableMerchantCollection](AvailableMerchantCollection.md)
 - [AvailableSingleitemCollection](AvailableSingleitemCollection.md)
 - [BackgroundColor](BackgroundColor.md)
 - [Callback](Callback.md)
 - [CardLimitation](CardLimitation.md)
 - [Coupon](Coupon.md)
 - [CouponCollection](CouponCollection.md)
 - [CouponRule](CouponRule.md)
 - [CreateCouponStockRequest](CreateCouponStockRequest.md)
 - [CreateCouponStockResponse](CreateCouponStockResponse.md)
 - [CutTypeMsg](CutTypeMsg.md)
 - [DeductBalanceMethod](DeductBalanceMethod.md)
 - [FavorAvailableTime](FavorAvailableTime.md)
 - [FixedAvailableTime](FixedAvailableTime.md)
 - [FixedValueStockMsg](FixedValueStockMsg.md)
 - [FormFile](FormFile.md)
 - [ImageMeta](ImageMeta.md)
 - [JumpTarget](JumpTarget.md)
 - [ListAvailableMerchantsRequest](ListAvailableMerchantsRequest.md)
 - [ListAvailableSingleitemsRequest](ListAvailableSingleitemsRequest.md)
 - [ListCouponsByFilterRequest](ListCouponsByFilterRequest.md)
 - [ListStocksRequest](ListStocksRequest.md)
 - [MediaImageRequest](MediaImageRequest.md)
 - [MediaImageResponse](MediaImageResponse.md)
 - [ModifyAvailableMerchantRequest](ModifyAvailableMerchantRequest.md)
 - [ModifyAvailableMerchantResponse](ModifyAvailableMerchantResponse.md)
 - [ModifyAvailableSingleitemRequest](ModifyAvailableSingleitemRequest.md)
 - [ModifyAvailableSingleitemResponse](ModifyAvailableSingleitemResponse.md)
 - [ModifyStockBudgetRequest](ModifyStockBudgetRequest.md)
 - [ModifyStockBudgetResponse](ModifyStockBudgetResponse.md)
 - [PatternInfo](PatternInfo.md)
 - [PauseStockBody](PauseStockBody.md)
 - [PauseStockRequest](PauseStockRequest.md)
 - [PauseStockResponse](PauseStockResponse.md)
 - [PlatformFavorStockStatus](PlatformFavorStockStatus.md)
 - [QueryCallbackRequest](QueryCallbackRequest.md)
 - [QueryCouponRequest](QueryCouponRequest.md)
 - [QueryStockRequest](QueryStockRequest.md)
 - [RefundFlowRequest](RefundFlowRequest.md)
 - [RefundFlowResponse](RefundFlowResponse.md)
 - [RestartStockBody](RestartStockBody.md)
 - [RestartStockRequest](RestartStockRequest.md)
 - [RestartStockResponse](RestartStockResponse.md)
 - [SendCouponBody](SendCouponBody.md)
 - [SendCouponRequest](SendCouponRequest.md)
 - [SendCouponResponse](SendCouponResponse.md)
 - [SetCallbackRequest](SetCallbackRequest.md)
 - [SetCallbackResponse](SetCallbackResponse.md)
 - [StartStockBody](StartStockBody.md)
 - [StartStockRequest](StartStockRequest.md)
 - [StartStockResponse](StartStockResponse.md)
 - [Stock](Stock.md)
 - [StockCollection](StockCollection.md)
 - [StockRule](StockRule.md)
 - [StockUseRule](StockUseRule.md)
 - [StopStockBody](StopStockBody.md)
 - [StopStockRequest](StopStockRequest.md)
 - [StopStockResponse](StopStockResponse.md)
 - [TradeType](TradeType.md)
 - [UseFlowRequest](UseFlowRequest.md)
 - [UseFlowResponse](UseFlowResponse.md)

