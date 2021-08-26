# CreateOrderRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Appid** | **string** | 微信分配的服务商appid | 
**OutOrderNo** | **string** | 服务商系统内部的分账单号，在服务商系统内部唯一，同一分账单号多次请求等同一次。只能是数字、大小写字母_-|*@  | 
**Receivers** | [**[]CreateOrderReceiver**](CreateOrderReceiver.md) | 分账接收方列表，可以设置出资商户作为分账接受方，最多可有50个分账接收方 | [可选] 
**SubAppid** | **string** | 微信分配的子商户公众账号ID，分账接收方类型包含PERSONAL_SUB_OPENID时必填。（直连商户不需要，服务商需要） | [可选] 
**SubMchid** | **string** | 微信支付分配的子商户号，即分账的出资商户号。（直连商户不需要，服务商需要） | [可选] 
**TransactionId** | **string** | 微信支付订单号 | 
**UnfreezeUnsplit** | **bool** | 1、如果为true，该笔订单剩余未分账的金额会解冻回分账方商户； 2、如果为false，该笔订单剩余未分账的金额不会解冻回分账方商户，可以对该笔订单再次进行分账。 | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


