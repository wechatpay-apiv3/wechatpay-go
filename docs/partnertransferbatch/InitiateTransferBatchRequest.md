# InitiateTransferBatchRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**SubMchid** | **string** | 特约商户号  | 
**SubAppid** | **string** | 微信分配的特约商户公众账号ID，特约商户授权类型为INFORMATION_AUTHORIZATION_TYPE和INFORMATION_AND_FUND_AUTHORIZATION_TYPE时 需要填写  | [可选] 
**AuthorizationType** | [**AuthType**](AuthType.md) | 特约商户授权类型 * &#x60;INFORMATION_AUTHORIZATION_TYPE&#x60; - 特约商户信息授权类型 * &#x60;FUND_AUTHORIZATION_TYPE&#x60; - 特约商户资金授权类型 * &#x60;INFORMATION_AND_FUND_AUTHORIZATION_TYPE&#x60; - 特约商户信息和资金授权类型  | 
**OutBatchNo** | **string** | 商户系统内部的商家批次单号，在商户系统内部唯一  | 
**BatchName** | **string** | 该笔批量转账的名称  | 
**BatchRemark** | **string** | 转账说明，UTF8编码，最多允许32个字符  | 
**TotalAmount** | **int64** | 转账金额单位为“分”。转账总金额必须与批次内所有明细转账金额之和保持一致，否则无法发起转账操作  | 
**TotalNum** | **int64** | 一个转账批次单最多发起三千笔转账。转账总笔数必须与批次内所有明细之和保持一致，否则无法发起转账操作  | 
**TransferDetailList** | [**[]TransferDetailInput**](TransferDetailInput.md) | 发起批量转账的明细列表，最多三千笔  | [可选] 
**SpAppid** | **string** | 微信分配的服务商商户公众账号ID，特约商户授权类型为FUND_AUTHORIZATION_TYPE时 需要填写  | [可选] 
**TransferPurpose** | [**TransferUseType**](TransferUseType.md) | 批量转账用途 * &#x60;GOODSPAYMENT&#x60; - 货款 * &#x60;COMMISSION&#x60; - 佣金 * &#x60;REFUND&#x60; - 退款 * &#x60;REIMBURSEMENT&#x60; - 报销 * &#x60;FREIGHT&#x60; - 运费 * &#x60;OTHERS&#x60; - 其他  | [可选] 
**TransferScene** | [**TransferScene**](TransferScene.md) | 商户的转账场景 * &#x60;ORDINARY_TRANSFER&#x60; - 普通转账 * &#x60;PAYROLL_CARD_TRANSFER&#x60; - 薪工卡转账  | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


