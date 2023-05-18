# OrdersEntity

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**OrderId** | **string** | 微信分账单号，微信系统返回的唯一标识 | 
**OutOrderNo** | **string** | 商户系统内部的分账单号，在商户系统内部唯一，同一分账单号多次请求等同一次。只能是数字、大小写字母_-|*@  | 
**Receivers** | [**[]OrderReceiverDetail**](OrderReceiverDetail.md) | 分账接收方列表 | [可选] 
**State** | [**OrderStatus**](OrderStatus.md) | 分账单状态（每个接收方的分账结果请查看receivers中的result字段），枚举值： 1、PROCESSING：处理中 2、FINISHED：分账完成  * &#x60;PROCESSING&#x60; - 处理中，  * &#x60;FINISHED&#x60; - 分账完成， | 
**SubMchid** | **string** | 微信支付分配的子商户号，即分账的出资商户号。（直连商户不需要，服务商需要） | [可选] 
**TransactionId** | **string** | 微信支付订单号 | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


