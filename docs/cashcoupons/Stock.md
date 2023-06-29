# Stock

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**StockId** | **string** | 批次id | 
**StockCreatorMchid** | **string** | 微信为创建方商户分配的商户号 | 
**StockName** | **string** | 批次名称 | 
**Status** | **string** | 批次状态： unactivated-未激活；audit-审核中；running-运行中；stoped-已停止；paused-暂停发放 | 
**CreateTime** | **string** | 批次创建时间 | 
**Description** | **string** | 批次描述信息 | 
**StockUseRule** | [**StockUseRule**](StockUseRule.md) | 普通发券批次特定信息 | [可选] 
**AvailableBeginTime** | **string** | 可用开始时间 | 
**AvailableEndTime** | **string** | 可用结束时间 | 
**DistributedCoupons** | **int64** | 已发券数量 | 
**NoCash** | **bool** | 是否无资金流，ture-是；false-否 | 
**StartTime** | **string** | 批次激活开启时间 | [可选] 
**StopTime** | **string** | 批次永久停止时间 | [可选] 
**CutToMessage** | [**CutTypeMsg**](CutTypeMsg.md) | 单品优惠特定信息 | [可选] 
**Singleitem** | **bool** | true-是；false-否 | 
**StockType** | **string** | NORMAL-代金券批次；DISCOUNT_CUT-立减与折扣；OTHER-其他 | 
**CardId** | **string** | 微信卡包ID | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


