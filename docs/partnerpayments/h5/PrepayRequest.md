# PrepayRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**SpAppid** | **string** | 服务商申请的公众号appid  | 
**SpMchid** | **string** | 服务商户号，由微信支付生成并下发  | 
**SubAppid** | **string** | 子商户申请的公众号appid  | [可选] 
**SubMchid** | **string** | 子商户的商户号，由微信支付生成并下发  | 
**Description** | **string** | 商品描述  | 
**OutTradeNo** | **string** | 商户订单号  | 
**TimeExpire** | **time.Time** | 订单失效时间，格式为rfc3339格式 | [可选] 
**Attach** | **string** | 附加数据  | [可选] 
**NotifyUrl** | **string** | 有效性：1. HTTPS；2. 不允许携带查询串。  | 
**GoodsTag** | **string** | 商品标记，代金券或立减优惠功能的参数。  | [可选] 
**LimitPay** | **[]string** | 指定支付方式  | [可选] 
**SupportFapiao** | **bool** | 传入true时，支付成功消息和支付详情页将出现开票入口。需要在微信支付商户平台或微信公众平台开通电子发票功能，传此字段才可生效。  | [可选] 
**Amount** | [**Amount**](Amount.md) |  | 
**Detail** | [**Detail**](Detail.md) |  | [可选] 
**SceneInfo** | [**SceneInfo**](SceneInfo.md) |  | 
**SettleInfo** | [**SettleInfo**](SettleInfo.md) |  | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


