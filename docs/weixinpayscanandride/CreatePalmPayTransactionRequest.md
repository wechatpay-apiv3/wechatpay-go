# CreatePalmPayTransactionRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Appid** | **string** | 商户在微信申请公众号或移动应用成功后分配的账号ID，登录平台为mp.weixin.qq.com或open.weixin.qq.com | 
**Description** | **string** | 商户自定义字段，用于交易账单中对扣费服务的描述。该字段长度限制为字节长度限制 | 
**Attach** | **string** | 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用，该字段长度限制为字节长度限制 | [可选] 
**OutTradeNo** | **string** | 商户系统内部订单号，只能是数字、大小写字母，且在同一个商户号下唯一。该字段长度限制为字节长度限制 | 
**GoodsTag** | **string** | 代金券或立减优惠功能的参数，说明详见代金券或立减优惠 | [可选] 
**PlanId** | **string** | 该模板ID是商户在向微信支付提权限申请时由微信支付生成 | 
**Openid** | **string** | 用户在商户AppID下的唯一标识。  | 
**NotifyUrl** | **string** | 接受扣款结果异步回调通知的URL，注意回调URL只接受HTTPS | 
**TradeScene** | [**PalmPayTradeScene**](PalmPayTradeScene.md) | 交易场景值，目前支持 ：PALM_METRO—刷掌地铁； | 
**MetroInfo** | [**PalmMetroSceneInfo**](PalmMetroSceneInfo.md) | 请求受理扣费 tradescene为PALM_METRO时，传入该参数 | [可选] 
**Amount** | [**OrderAmount**](OrderAmount.md) | 订单金额 | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


