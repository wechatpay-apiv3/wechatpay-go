# Refund4Create

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**RefundId** | **string** | 微信支付退款订单号 | 
**OutRefundNo** | **string** | 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。 | 
**CreateTime** | **time.Time** | 退款受理时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日13点29分35秒。 | 
**Amount** | [**RefundAmount**](RefundAmount.md) | 订单退款金额信息 | 
**PromotionDetail** | [**[]PromotionDetail**](PromotionDetail.md) | 优惠退款功能信息 | [可选] 
**RefundAccount** | **string** | REFUND_SOURCE_PARTNER_ADVANCE : 电商平台垫付 REFUND_SOURCE_SUB_MERCHANT : 二级商户，默认值 注意：申请退款时可能会出现暂未扣除退款资金的情况，此时不返回本字段，需查单确认。 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


