# AddReceiverResponse

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Account** | **string** | 参考请求参数 | 
**CustomRelation** | **string** | 参考请求参数 | [可选] 
**Name** | **string** | 参考请求参数。使用APIv3敏感信息加密方式 | [可选] 
**RelationType** | [**ReceiverRelationType**](ReceiverRelationType.md) | 参考请求参数  * &#x60;SERVICE_PROVIDER&#x60; - 服务商，  * &#x60;STORE&#x60; - 门店，  * &#x60;STAFF&#x60; - 员工，  * &#x60;STORE_OWNER&#x60; - 店主，  * &#x60;PARTNER&#x60; - 合作伙伴，  * &#x60;HEADQUARTER&#x60; - 总部，  * &#x60;BRAND&#x60; - 品牌方，  * &#x60;DISTRIBUTOR&#x60; - 分销商，  * &#x60;USER&#x60; - 用户，  * &#x60;SUPPLIER&#x60; - 供应商，  * &#x60;CUSTOM&#x60; - 自定义， | 
**SubMchid** | **string** | 参考请求参数 | [可选] 
**Type** | [**ReceiverType**](ReceiverType.md) | 参考请求参数  * &#x60;MERCHANT_ID&#x60; - 商户号，  * &#x60;PERSONAL_OPENID&#x60; - 个人openid（由父商户APPID转换得到），  * &#x60;PERSONAL_SUB_OPENID&#x60; - 个人sub_openid（由子商户APPID转换得到）（直连商户不需要，服务商需要）， | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


