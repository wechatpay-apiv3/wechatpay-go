# PlatformFavorStockEntity

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**StockId** | **string** | 立减金批次的主键，唯一定义此资源的标识 | 
**CreateTime** | **string** | 立减金批次创建成功时间 | 
**StockName** | **string** | 立减金批次展示名称，按照平台规则固定为x元立减金 | 
**BelongMerchant** | **string** | 批次归属商户id | 
**SendRule** | [**PlatformFavorSendRule**](PlatformFavorSendRule.md) |  | 
**UseRule** | [**PlatformFavorUseRule**](PlatformFavorUseRule.md) |  | 
**DisplayPattern** | [**PlatformFavorDisplayPattern**](PlatformFavorDisplayPattern.md) |  | [可选] 
**StockStatus** | [**PlatformFavorStockStatus**](PlatformFavorStockStatus.md) | 批次状态 | 
**ActiveTime** | **string** | 批次激活时间 | [可选] 
**PauseTime** | **string** | 批次暂停时间 | [可选] 
**StatisticalData** | [**StatisticalData**](StatisticalData.md) |  | [可选] 
**CardId** | **string** | 卡包id | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


