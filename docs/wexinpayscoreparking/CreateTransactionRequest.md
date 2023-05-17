# CreateTransactionRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Appid** | **string** | appid是商户在微信申请公众号或移动应用成功后分配的帐号ID，登录平台为mp.weixin.qq.com或open.weixin.qq.com | 
**SubAppid** | **string** | 子公众账号id，服务商模式下选传，用于扣费信息的商户信息展示， | [可选] 
**SubMchid** | **string** | 微信支付分配的子商户号，服务商模式下必传 | [可选] 
**Description** | **string** | 商户自定义字段，用于交易账单中对扣费服务的描述。 | 
**Attach** | **string** | 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用 | [可选] 
**OutTradeNo** | **string** | 商户系统内部订单号，只能是数字、大小写字母_-*且在同一个商户号下唯一 | 
**TradeScene** | **string** | 交易场景值，目前支持 ：PARKING：车场停车场景  | 
**GoodsTag** | **string** | 代金券或立减优惠功能的参数，说明详见[代金券或立减优惠](https://pay.weixin.qq.com/wiki/doc/apiv3/open/pay/chapter5_1_2.shtml) | [可选] 
**NotifyUrl** | **string** | 接受扣款结果异步回调通知的url，注意回调url只接受https | 
**ProfitSharing** | **string** | Y：是，需要分账 N：否，不分账 字母要求大写，不传默认不分账，分账详细说明见[直连分账API](https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter&#x3D;26_1)、[服务商分账API文档](https://pay.weixin.qq.com/wiki/doc/api/allocation_sl.php?chapter&#x3D;24_1&amp;index&#x3D;1) | [可选] 
**Amount** | [**OrderAmount**](OrderAmount.md) | 订单金额信息 | 
**ParkingInfo** | [**ParkingTradeScene**](ParkingTradeScene.md) | 当交易场景为PARKING时，需要在该字段添加停车场景信息 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


