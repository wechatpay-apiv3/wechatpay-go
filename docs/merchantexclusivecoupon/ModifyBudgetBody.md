# ModifyBudgetBody

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**TargetMaxCoupons** | **int64** | 批次最大发放个数，target\\_max\\_coupons与target\\_max\\_coupons\\_by\\_day只能修改一个 注：目标批次即修改后的批次 | [可选] 
**CurrentMaxCoupons** | **int64** | 当前批次最大发放个数，当传入target\\_max\\_coupons大于0时，current\\_max\\_coupons必传 注：当前批次即未修改的批次 | [可选] 
**TargetMaxCouponsByDay** | **int64** | 单天发放上限个数，target\\_max\\_coupons与target\\_max\\_coupons\\_by\\_day只能修改一个 注：目标批次即修改后的批次 | [可选] 
**CurrentMaxCouponsByDay** | **int64** | 当前单天发放上限个数 ，当传入target\\_max\\_coupons\\_by\\_day大于0时，current\\_max\\_coupons\\_by\\_day必填 注：当前批次即未修改的批次 | [可选] 
**ModifyBudgetRequestNo** | **string** | 修改预算请求单据号 | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


