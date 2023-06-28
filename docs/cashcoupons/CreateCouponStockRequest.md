# CreateCouponStockRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**StockName** | **string** | 批次名称 | 
**Comment** | **string** | 仅配置商户可见，用于自定义信息 | [可选] 
**BelongMerchant** | **string** | 批次归属商户号 | 
**AvailableBeginTime** | **string** | 批次开始时间 | 
**AvailableEndTime** | **string** | 批次结束时间 | 
**StockUseRule** | [**StockRule**](StockRule.md) | 批次使用规则 | 
**PatternInfo** | [**PatternInfo**](PatternInfo.md) | 代金券详情页 | [可选] 
**CouponUseRule** | [**CouponRule**](CouponRule.md) |  | 
**NoCash** | **bool** | 是否无资金流，true-是；false-否 | 
**StockType** | **string** | 批次类型，NORMAL-固定面额满减券批次；DISCOUNT-折扣券批次；EXCHAHGE-换购券批次；RANDOM-千人千面券批次 | 
**OutRequestNo** | **string** | 商户创建批次凭据号（格式：商户id+日期+流水号），商户侧需保持唯一性 | 
**ExtInfo** | **string** | 扩展属性字段，按json格式，暂时无需填写 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


