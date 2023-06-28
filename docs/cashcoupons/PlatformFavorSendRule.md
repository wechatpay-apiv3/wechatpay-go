# PlatformFavorSendRule

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**MaxCoupons** | **int64** | 批次可发放券的总张数 | 
**MaxAmount** | **int64** | 总消耗金额，单位分 | 
**MaxCouponsByDay** | **int64** | 单天发放上限个数 | [可选] 
**MaxCouponsPerUser** | **int64** | 用户最大可领个数 | 
**NaturalPersonLimit** | **bool** | true-是；false-否，默认否 | [可选] 
**PreventApiAbuse** | **bool** | true-是；false-否，默认否 | [可选] 
**DeductBalanceMethod** | [**DeductBalanceMethod**](DeductBalanceMethod.md) | 立减金批次预算的扣款方式 | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


