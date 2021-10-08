# TransferDetailEntity

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Mchid** | **string** | 微信支付分配的商户号  | 
**OutBatchNo** | **string** | 商户系统内部的商家批次单号，在商户系统内部唯一  | 
**BatchId** | **string** | 微信批次单号，微信商家转账系统返回的唯一标识  | 
**Appid** | **string** | 申请商户号的appid或商户号绑定的appid（企业号corpid即为此appid）  | 
**OutDetailNo** | **string** | 商户系统内部区分转账批次单下不同转账明细单的唯一标识  | 
**DetailId** | **string** | 微信支付系统内部区分转账批次单下不同转账明细单的唯一标识  | 
**DetailStatus** | **string** | PROCESSING:转账中。正在处理中，转账结果尚未明确   SUCCESS:转账成功   FAIL:转账失败。需要确认失败原因后，再决定是否重新发起对该笔明细单的转账（并非整个转账批次单）  | 
**TransferAmount** | **int64** | 转账金额单位为“分”  | 
**TransferRemark** | **string** | 单条转账备注（微信用户会收到该备注），UTF8编码，最多允许32个字符  | 
**FailReason** | [**FailReasonType**](FailReasonType.md) | 如果转账失败则有失败原因 * &#x60;ACCOUNT_FROZEN&#x60; - 账户冻结 * &#x60;REAL_NAME_CHECK_FAIL&#x60; - 用户未实名 * &#x60;NAME_NOT_CORRECT&#x60; - 用户姓名校验失败 * &#x60;OPENID_INVALID&#x60; - Openid校验失败 * &#x60;TRANSFER_QUOTA_EXCEED&#x60; - 超过用户单笔收款额度 * &#x60;DAY_RECEIVED_QUOTA_EXCEED&#x60; - 超过用户单日收款额度 * &#x60;MONTH_RECEIVED_QUOTA_EXCEED&#x60; - 超过用户单月收款额度 * &#x60;DAY_RECEIVED_COUNT_EXCEED&#x60; - 超过用户单日收款次数 * &#x60;PRODUCT_AUTH_CHECK_FAIL&#x60; - 产品权限校验失败 * &#x60;OVERDUE_CLOSE&#x60; - 转账关闭 * &#x60;ID_CARD_NOT_CORRECT&#x60; - 用户身份证校验失败 * &#x60;ACCOUNT_NOT_EXIST&#x60; - 用户账户不存在 * &#x60;TRANSFER_RISK&#x60; - 转账存在风险  | [可选] 
**Openid** | **string** | 商户appid下，某用户的openid  | 
**UserName** | ***os.File** | 收款方姓名。采用标准RSA算法，公钥由微信侧提供  | 
**InitiateTime** | **time.Time** | 转账发起的时间，按照使用rfc3339所定义的格式，格式为YYYY-MM-DDThh:mm:ss+TIMEZONE  | 
**UpdateTime** | **time.Time** | 明细最后一次状态变更的时间，按照使用rfc3339所定义的格式，格式为YYYY-MM-DDThh:mm:ss+TIMEZONE  | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


