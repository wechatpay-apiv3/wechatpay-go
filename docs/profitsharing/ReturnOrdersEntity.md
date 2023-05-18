# ReturnOrdersEntity

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | 需要从分账接收方回退的金额，单位为分，只能为整数 | 
**CreateTime** | **time.Time** | 分账回退创建时间，遵循RFC3339标准格式 | 
**Description** | **string** | 分账回退的原因描述 | 
**FailReason** | [**ReturnOrderFailReason**](ReturnOrderFailReason.md) | 失败原因。包含以下枚举值： ACCOUNT_ABNORMAL : 分账接收方账户异常 TIME_OUT_CLOSED : 超时关单  * &#x60;ACCOUNT_ABNORMAL&#x60; - 分账接收方账户异常，  * &#x60;BALANCE_NOT_ENOUGH&#x60; - 余额不足，  * &#x60;TIME_OUT_CLOSED&#x60; - 超时关单， | [可选] 
**FinishTime** | **time.Time** | 分账回退完成时间，遵循RFC3339标准格式 | 
**OrderId** | **string** | 参考请求参数 | 
**OutOrderNo** | **string** | 参考请求参数 | 
**OutReturnNo** | **string** | 参考请求参数 | 
**Result** | [**ReturnOrderStatus**](ReturnOrderStatus.md) | 如果请求返回为处理中，则商户可以通过调用回退结果查询接口获取请求的最终处理结果。如果查询到回退结果在处理中，请勿变更商户回退单号，使用相同的参数再次发起分账回退，否则会出现资金风险。在处理中状态的回退单如果5天没有成功，会因为超时被设置为已失败。 枚举值： PROCESSING：处理中 SUCCESS：已成功 FAILED：已失败   * &#x60;PROCESSING&#x60; - 处理中，  * &#x60;SUCCESS&#x60; - 已成功，  * &#x60;FAILED&#x60; - 已失败， | 
**ReturnId** | **string** | 微信分账回退单号，微信系统返回的唯一标识 | 
**ReturnMchid** | **string** | 只能对原分账请求中成功分给商户接收方进行回退 | 
**SubMchid** | **string** | 参考请求参数 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


