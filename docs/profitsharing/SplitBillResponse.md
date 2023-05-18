# SplitBillResponse

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**DownloadUrl** | **string** | 供下一步请求账单文件的下载地址，该地址30s内有效 | 
**HashType** | [**SplitBillHashType**](SplitBillHashType.md) | 原始账单（gzip需要解压缩）的摘要算法，用于校验文件的完整性  * &#x60;SHA1&#x60; - SHA1，Secure Hash Algorithm 1 | 
**HashValue** | **string** | 原始账单（gzip需要解压缩）的摘要值，用于校验文件的完整性 | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


