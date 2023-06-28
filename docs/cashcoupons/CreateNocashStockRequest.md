# CreateNocashStockRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**StockName** | **string** | 批次名称 | 
**AvailableMerchants** | **[]string** | 可核销商户 | [可选] 
**AvailableBeginTime** | **string** | 批次开始时间 | 
**AvailableEndTime** | **string** | 批次结束时间 | 
**NaturalPersonLimit** | **bool** | true-是；false-否，默认否 | 
**PreventApiAbuse** | **bool** | true-是；false-否，默认否 | 
**FavorAvailableTime** | [**FavorAvailableTime**](FavorAvailableTime.md) | 优惠生效时间（包括立减、券） | 
**Description** | **string** | 代金券说明 | [可选] 
**StockUseRule** | [**StockUseRule**](StockUseRule.md) | 批次使用规则 | 
**StockType** | **string** | NORMAL-代金券批次；DISCOUNT_CUT-立减与折扣； | 
**DiscountCutMessage** | [**DiscountCutMsg**](DiscountCutMsg.md) | 立减与折扣批次特定信息 | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


