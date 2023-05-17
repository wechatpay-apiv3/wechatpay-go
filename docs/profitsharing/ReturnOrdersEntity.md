# ReturnOrdersEntity

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**SubMchid** | **string** | 参考请求参数 | [可选] 
**OrderId** | **string** | 微信分账单号，微信系统返回的唯一标识 | 
**OutOrderNo** | **string** | 商户系统内部的分账单号，在商户系统内部唯一，同一分账单号多次请求等同一次 | 
**OutReturnNo** | **string** | 调用回退接口提供的商户系统内部的回退单号 | 
**ReturnId** | **string** | 微信分账回退单号，微信系统返回的唯一标识 | 
**ReturnMchid** | **string** | 只能对原分账请求中成功分给商户接收方进行回退 | 
**Amount** | **int64** | 需要从分账接收方回退的金额，单位为分，只能为整数 | 
**Description** | **string** | 分账回退的原因描述 | 
**Result** | [**ReturnOrderStatus**](ReturnOrderStatus.md) | 如果请求返回为处理中，则商户可以通过调用回退结果查询接口获取请求的最终处理结果。如果查询到回退结果在处理中，请勿变更商户回退单号，使用相同的参数再次发起分账回退，否则会出现资金风险。在处理中状态的回退单如果5天没有成功，会因为超时被设置为已失败 | 
**FailReason** | [**ReturnOrderFailReason**](ReturnOrderFailReason.md) | 失败原因 | [可选] 
**CreateTime** | **time.Time** | 分账回退创建时间，遵循RFC3339标准格式 | 
**FinishTime** | **time.Time** | 分账回退完成时间，遵循RFC3339标准格式 | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


