# OrderReceiverDetail

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Account** | **string** | 1、类型是MERCHANT_ID时，是商户号 2、类型是PERSONAL_OPENID时，是个人openid 3、类型是PERSONAL_SUB_OPENID时，是个人sub_openid | 
**Amount** | **int64** | 分账金额，单位为分，只能为整数，不能超过原订单支付金额及最大分账比例金额 | 
**CreateTime** | **time.Time** | 分账创建时间，遵循RFC3339标准格式 | 
**Description** | **string** | 分账的原因描述，分账账单中需要体现 | 
**DetailId** | **string** | 微信分账明细单号，每笔分账业务执行的明细单号，可与资金账单对账使用 | 
**FailReason** | [**DetailFailReason**](DetailFailReason.md) | 分账失败原因。包含以下枚举值： 1、ACCOUNT_ABNORMAL : 分账接收账户异常 2、NO_RELATION : 分账关系已解除 3、RECEIVER_HIGH_RISK : 高风险接收方 4、RECEIVER_REAL_NAME_NOT_VERIFIED : 接收方未实名 5、NO_AUTH : 分账权限已解除  * &#x60;ACCOUNT_ABNORMAL&#x60; - 分账接收账户异常，  * &#x60;NO_RELATION&#x60; - 分账关系已解除，  * &#x60;RECEIVER_HIGH_RISK&#x60; - 高风险接收方，  * &#x60;RECEIVER_REAL_NAME_NOT_VERIFIED&#x60; - 接收方未实名，  * &#x60;NO_AUTH&#x60; - 分账权限已解除， | [可选] 
**FinishTime** | **time.Time** | 分账完成时间，遵循RFC3339标准格式 | 
**Result** | [**DetailStatus**](DetailStatus.md) | 枚举值： 1、PENDING：待分账 2、SUCCESS：分账成功 3、CLOSED：已关闭  * &#x60;PENDING&#x60; - 待分账，  * &#x60;SUCCESS&#x60; - 分账成功，  * &#x60;CLOSED&#x60; - 已关闭， | 
**Type** | [**ReceiverType**](ReceiverType.md) | 1、MERCHANT_ID：商户号 2、PERSONAL_OPENID：个人openid（由父商户APPID转换得到） 3、PERSONAL_SUB_OPENID: 个人sub_openid（由子商户APPID转换得到）  * &#x60;MERCHANT_ID&#x60; - 商户号，  * &#x60;PERSONAL_OPENID&#x60; - 个人openid（由父商户APPID转换得到），  * &#x60;PERSONAL_SUB_OPENID&#x60; - 个人sub_openid（由子商户APPID转换得到）（直连商户不需要，服务商需要）， | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


