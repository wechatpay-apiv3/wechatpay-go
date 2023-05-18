# AddReceiverRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Account** | **string** | 类型是MERCHANT_ID时，是商户号 类型是PERSONAL_OPENID时，是个人openid 类型是PERSONAL_SUB_OPENID时，是个人sub_openid | 
**Appid** | **string** | 微信分配的公众账号ID | 
**CustomRelation** | **string** | 子商户与接收方具体的关系，本字段最多10个字。  当字段relation_type的值为CUSTOM时，本字段必填 当字段relation_type的值不为CUSTOM时，本字段无需填写 | [可选] 
**Name** | **string** | 分账接收方类型是MERCHANT_ID时，是商户全称（必传），当商户是小微商户或个体户时，是开户人姓名 分账接收方类型是PERSONAL_OPENID时，是个人姓名（选传，传则校验） 分账接收方类型是PERSONAL_SUB_OPENID时，是个人姓名（选传，传则校验） 1、此字段需要加密，的加密方法详见：敏感信息加密说明 2、使用微信支付平台证书中的公钥 3、使用RSAES-OAEP算法进行加密 4、将请求中HTTP头部的Wechatpay-Serial设置为证书序列号 | [可选] 
**RelationType** | [**ReceiverRelationType**](ReceiverRelationType.md) | 子商户与接收方的关系。 本字段值为枚举： SERVICE_PROVIDER：服务商 STORE：门店  STAFF：员工 STORE_OWNER：店主 PARTNER：合作伙伴 HEADQUARTER：总部 BRAND：品牌方 DISTRIBUTOR：分销商 USER：用户 SUPPLIER：供应商 CUSTOM：自定义  * &#x60;SERVICE_PROVIDER&#x60; - 服务商，  * &#x60;STORE&#x60; - 门店，  * &#x60;STAFF&#x60; - 员工，  * &#x60;STORE_OWNER&#x60; - 店主，  * &#x60;PARTNER&#x60; - 合作伙伴，  * &#x60;HEADQUARTER&#x60; - 总部，  * &#x60;BRAND&#x60; - 品牌方，  * &#x60;DISTRIBUTOR&#x60; - 分销商，  * &#x60;USER&#x60; - 用户，  * &#x60;SUPPLIER&#x60; - 供应商，  * &#x60;CUSTOM&#x60; - 自定义， | 
**SubAppid** | **string** | 子商户的公众账号ID，分账接收方类型包含PERSONAL_SUB_OPENID时必填。（直连商户不需要，服务商需要） | [可选] 
**SubMchid** | **string** | 微信支付分配的子商户号，即分账的出资商户号。（直连商户不需要，服务商需要） | [可选] 
**Type** | [**ReceiverType**](ReceiverType.md) | 枚举值： MERCHANT_ID：商户ID PERSONAL_OPENID：个人openid（由父商户APPID转换得到） PERSONAL_SUB_OPENID：个人sub_openid（由子商户APPID转换得到）  * &#x60;MERCHANT_ID&#x60; - 商户号，  * &#x60;PERSONAL_OPENID&#x60; - 个人openid（由父商户APPID转换得到），  * &#x60;PERSONAL_SUB_OPENID&#x60; - 个人sub_openid（由子商户APPID转换得到）（直连商户不需要，服务商需要）， | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


