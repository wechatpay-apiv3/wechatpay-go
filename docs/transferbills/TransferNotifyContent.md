# TransferNotifyContent

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Mchid** | **string** | 微信支付分配的商户号 | 
**OutBillNo** | **string** | 商户系统内部的商家订单号 | 
**TransferBillNo** | **string** | 微信转账单号，微信商家转账系统返回的唯一标识 | 
**Appid** | **string** | 申请商户号的appid或商户号绑定的appid | 
**State** | [**TransferState**](TransferState.md) | 转账状态 | 
**TransferAmount** | **int64** | 转账金额，单位为"分" | 
**TransferRemark** | **string** | 转账备注 | 
**FailReason** | **string** | 如果转账失败则有失败原因 | [可选] 
**Openid** | **string** | 用户在直连商户appid下的唯一标识 | 
**UserName** | **string** | 收款方真实姓名 | [可选] 
**CreateTime** | **string** | 转账创建时间，遵循rfc3339标准格式 | 
**UpdateTime** | **string** | 转账更新时间，遵循rfc3339标准格式 | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)
