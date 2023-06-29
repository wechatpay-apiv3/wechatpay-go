# ListStocksRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Offset** | **int64** | 页码从0开始，默认第0页 | 
**Limit** | **int64** | 分页大小，最大10 | 
**StockCreatorMchid** | **string** | 批次创建方商户号 | 
**CreateStartTime** | **string** | 起始创建时间 | [可选] 
**CreateEndTime** | **string** | 终止创建时间 | [可选] 
**Status** | **string** | 批次状态： unactivated-未激活；audit-审核中；running-运行中；stoped-已停止；paused-暂停发放 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


