# CreateTransactionRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Appid** | **string** | 商户在微信申请公众号或移动应用成功后分配的账号ID，登录平台为mp.weixin.qq.com或open.weixin.qq.com | 
**SubAppid** | **string** | 子公众账号ID，服务商模式下选传，用于扣费信息的商户信息展示， | [可选] 
**SubMchid** | **string** | 微信支付分配的子商户号，服务商模式下必传 | [可选] 
**Description** | **string** | 商户自定义字段，用于交易账单中对扣费服务的描述。该字段长度限制为字节长度限制 | 
**Attach** | **string** | 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用，该字段长度限制为字节长度限制 | [可选] 
**OutTradeNo** | **string** | 商户系统内部订单号，只能是数字、大小写字母，且在同一个商户号下唯一。该字段长度限制为字节长度限制 | 
**TradeScene** | [**TradeScene**](TradeScene.md) | 交易场景值，现在支持公交和地铁的场景值 | 
**GoodsTag** | **string** | 代金券或立减优惠功能的参数，说明详见代金券或立减优惠 | [可选] 
**ContractId** | **string** | 签约成功后，微信返回代扣签约ID 查用户是否能继续使用乘车卡时必传，此时，才会返回签约信息，如contract_state | 
**NotifyUrl** | **string** | 接受扣款结果异步回调通知的URL，注意回调URL只接受HTTPS | 
**Amount** | [**OrderAmount**](OrderAmount.md) | 订单金额信息 | 
**BusInfo** | [**BusSceneInfo**](BusSceneInfo.md) | 请求受理扣费 tradescene为BUS时，传入该参数 | [可选] 
**MetroInfo** | [**MetroSceneInfo**](MetroSceneInfo.md) | 请求受理扣费 tradescene为METRO时，传入该参数 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


