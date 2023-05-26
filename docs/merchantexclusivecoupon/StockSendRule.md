# StockSendRule

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**MaxAmount** | **int64** | 总预算金额，单位分 注：该字段暂未开放 | [可选] 
**MaxCoupons** | **int64** | 批次最大可发放个数限制 特殊规则：取值范围 1 ≤ value ≤ 1000000000  | [可选] 
**MaxCouponsPerUser** | **int64** | 用户可领个数，每个用户最多100张券。 | 
**MaxAmountByDay** | **int64** | 单天发放上限金额 注：该字段暂未开放 | [可选] 
**MaxCouponsByDay** | **int64** | 单天发放上限个数（stock\\_type为DISCOUNT或EXCHANGE时可传入此字段控制单天发放上限）。 特殊规则：取值范围 1 ≤ value ≤ 1000000000 | [可选] 
**NaturalPersonLimit** | **bool** | 不填默认否，枚举值： true：是 false：否 注：自然人防刷即同证件号下的所有账户合并计算的限领次数（限领次数指的是参数字段“用户最大领取个数”填写的值） | [可选] 
**PreventApiAbuse** | **bool** | true-是；false-否，不填默认否 | [可选] 
**Transferable** | **bool** | 不填默认否，枚举值： true：是 false：否 该字段暂未开放 | [可选] 
**Shareable** | **bool** | 不填默认否，枚举值： true：是 false：否 该字段暂未开放 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


