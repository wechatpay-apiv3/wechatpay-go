# Coupon

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**StockCreatorMchid** | **string** | 微信为创建方商户分配的商户号 | 
**StockId** | **string** | 批次id | 
**CutToMessage** | [**CutTypeMsg**](CutTypeMsg.md) | 单品优惠特定信息 | [可选] 
**CouponName** | **string** | 代金券名称 | 
**Status** | **string** | 代金券状态：SENDED-可用，USED-已实扣，EXPIRED-已过期 | 
**Description** | **string** | 代金券描述说明字段 | 
**CreateTime** | **string** | 领券时间 | 
**CouponType** | **string** | NORMAL-满减券；CUT_TO-减至券 | 
**NoCash** | **bool** | true-是；false-否 | 
**AvailableBeginTime** | **string** | 可用开始时间 | 
**AvailableEndTime** | **string** | 可用结束时间 | 
**Singleitem** | **bool** | TRUE-是；FALSE-否 | 
**NormalCouponInformation** | [**FixedValueStockMsg**](FixedValueStockMsg.md) | 普通满减券面额、门槛信息 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


