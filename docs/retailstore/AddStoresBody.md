# AddStoresBody

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**OutRequestNo** | **string** | 商户新增小店活动门店凭据号，商户侧需保持唯一性。可包含英文字母，数字，\\\\|，\\\\_，\\\\*，\\\\-等内容，不允许出现其他不合法符号。 | 
**AddTime** | **string** | 添加时间，遵循rfc3339标准格式，格式为yyyy-MM-DDTHH:mm:ss+TIMEZONE，yyyy-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒。 | 
**Stores** | [**[]RetailStoreInfo**](RetailStoreInfo.md) | 待新增的小店活动门店列表 | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


