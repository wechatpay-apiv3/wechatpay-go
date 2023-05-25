# UploadCouponCodeResponse

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**StockId** | **string** | 商家券批次号 | 
**TotalCount** | **int64** | 本次上传操作，去重后实际上传的code数目 | 
**SuccessCount** | **int64** | 本次上传操作上传成功个数 | 
**SuccessCodes** | **[]string** | 本次新增上传成功的code信息。 特殊规则：单个券code长度为【1，32】，条目个数限制为【1，200】。 | [可选] 
**SuccessTime** | **string** | 上传操作完成时间，遵循[rfc3339](https://datatracker.ietf.org/doc/html/rfc3339)标准格式，格式为yyyy-MM-DDTHH:mm:ss+TIMEZONE，yyyy-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒。 | 
**FailCount** | **int64** | 本次上传操作上传失败的code数 | [可选] 
**FailCodes** | [**[]UploadCouponCodeFailReason**](UploadCouponCodeFailReason.md) | 本次导入失败的code信息，请参照错误信息，修改后重试 | [可选] 
**ExistCodes** | **[]string** | 历史已存在的code列表，本次不会重复导入。 特殊规则：单个券code长度为【1，32】，条目个数限制为【1，200】。 | [可选] 
**DuplicateCodes** | **[]string** | 本次重复导入的code会被自动过滤，仅保留一个做导入，如满足要求则成功；如不满足要求，则失败；请参照报错提示修改重试。 特殊规则：单个券code长度为【1，32】，条目个数限制为【1，200】。 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


