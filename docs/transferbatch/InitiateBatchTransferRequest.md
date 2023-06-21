# InitiateBatchTransferRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Appid** | **string** | 申请商户号的appid或商户号绑定的appid（企业号corpid即为此appid） | 
**OutBatchNo** | **string** | 商户系统内部的商家批次单号，要求此参数只能由数字、大小写字母组成，在商户系统内部唯一 | 
**BatchName** | **string** | 该笔批量转账的名称 | 
**BatchRemark** | **string** | 转账说明，UTF8编码，最多允许32个字符 | 
**TotalAmount** | **int64** | 转账金额单位为“分”。转账总金额必须与批次内所有明细转账金额之和保持一致，否则无法发起转账操作 | 
**TotalNum** | **int64** | 一个转账批次单最多发起一千笔转账。转账总笔数必须与批次内所有明细之和保持一致，否则无法发起转账操作 | 
**TransferDetailList** | [**[]TransferDetailInput**](TransferDetailInput.md) | 发起批量转账的明细列表，最多一千笔 | 
**TransferSceneId** | **string** | 该批次转账使用的转账场景，如不填写则使用商家的默认场景，如无默认场景可为空，可前往“商家转账到零钱-前往功能”中申请。 如：1001-现金营销 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


