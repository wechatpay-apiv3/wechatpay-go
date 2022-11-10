# CreateTransferBatchRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**SubMchid** | **string** | 特约商户号 | 
**SubAppid** | **string** | 微信分配的特约商户公众账号ID，特约商户授权类型为INFORMATION_AUTHORIZATION_TYPE和INFORMATION_AND_FUND_AUTHORIZATION_TYPE时 需要填写 | [可选] 
**AuthorizationType** | [**AuthType**](AuthType.md) | 特约商户授权类型 | 
**OutBatchNo** | **string** | 商户系统内部的商家批次单号，在商户系统内部唯一 | 
**BatchName** | **string** | 该笔批量转账的名称 | 
**BatchRemark** | **string** | 转账说明，UTF8编码，最多允许32个字符 | 
**TotalAmount** | **int64** | 转账金额单位为“分”。转账总金额必须与批次内所有明细转账金额之和保持一致，否则无法发起转账操作 | 
**TotalNum** | **int64** | 一个转账批次单最多发起三千笔转账。转账总笔数必须与批次内所有明细之和保持一致，否则无法发起转账操作 | 
**TransferDetailList** | [**[]TransferDetailInput**](TransferDetailInput.md) | 发起批量转账的明细列表，最多三千笔 | 
**SpAppid** | **string** | 微信分配的服务商商户公众账号ID，特约商户授权类型为FUND_AUTHORIZATION_TYPE时 需要填写 | [可选] 
**EmploymentType** | [**EmploymentType**](EmploymentType.md) | 微工卡服务仅支持用于与商户有用工关系的用户，需明确用工类型；参考值：长期用工：LONG_TERM_EMPLOYMENT，短期用工：SHORT_TERM_EMPLOYMENT，合作关系：COOPERATION_EMPLOYMENT | [可选] 
**EmploymentScene** | [**EmploymentScene**](EmploymentScene.md) | 用工场景；参考值：LOGISTICS：物流；MANUFACTURING：制造业；HOTEL：酒店；CATERING：餐饮业；EVENT：活动促销；RETAIL：零售；OTHERS：其他 | [可选] 
**BusinessType** | [**BusinessType**](BusinessType.md) | 传入业务ID后必填，区分不同任务来源于哪个业务系统 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


