# SplitBillRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**SubMchid** | **string** | 若商户是直连商户： - 无需填写该字段  若商户是服务商： - 不填则默认返回服务商下的所有分账账单。 - 如需下载某个子商户下的分账账单，则填指定的子商户号。 | [可选] 
**BillDate** | **string** | 格式YYYY-MM-DD。仅支持三个月内的账单下载申请。 | 
**TarType** | [**SplitBillTarType**](SplitBillTarType.md) | 不填则以不压缩的方式返回数据流 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


