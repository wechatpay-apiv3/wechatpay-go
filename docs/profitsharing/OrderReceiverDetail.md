# OrderReceiverDetail

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | 分账金额，单位为分，只能为整数，不能超过原订单支付金额及最大分账比例金额 | 
**Description** | **string** | 分账的原因描述，分账账单中需要体现 | 
**Type** | [**ReceiverType**](ReceiverType.md) | 接收方类型 | 
**Account** | **string** | 1、类型是MERCHANT_ID时，是商户号 2、类型是PERSONAL_OPENID时，是个人openid 3、类型是PERSONAL_SUB_OPENID时，是个人sub_openid | 
**Result** | [**DetailStatus**](DetailStatus.md) | 分账结果 | 
**FailReason** | [**DetailFailReason**](DetailFailReason.md) | 分账失败原因，状态为CLOSED时返回 | [可选] 
**CreateTime** | **time.Time** | 分账创建时间，遵循RFC3339标准格式 | 
**FinishTime** | **time.Time** | 分账完成时间，遵循RFC3339标准格式 | 
**DetailId** | **string** | 微信分账明细单号，每笔分账业务执行的明细单号，可与资金账单对账使用 | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


