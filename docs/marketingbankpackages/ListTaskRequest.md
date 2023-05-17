# ListTaskRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**PackageId** | **string** | 号码包唯一标识符。可在微信支付商户平台创建号码包后获得。 | 
**Filename** | **string** | 任务上传的文件名 | [可选] 
**Offset** | **int64** | 分页页码，页码从0开始 | [可选] 
**Limit** | **int64** | 分页大小，单次最多查询20条记录 | [可选] 
**Status** | [**TaskStatus**](TaskStatus.md) | 上传任务的状态：PROCESSING - 处理中，FINISHED - 已完成 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


