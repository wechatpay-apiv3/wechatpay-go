# TransferDetailEntity

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**SpMchid** | **string** | 微信支付分配的商户号，此处为服务商商户号  | 
**OutBatchNo** | **string** | 商户系统内部的商家批次单号，在商户系统内部唯一  | 
**BatchId** | **string** | 微信支付批次单号，微信商家转账系统返回的唯一标识  | 
**Appid** | **string** | 微信分配的商户公众账号ID。特约商户授权类型为INFORMATION_AUTHORIZATION_TYPE和INFORMATION_AND_FUND_AUTHORIZATION_TYPE时对应的是特约商户的appid，特约商户授权类型为FUND_AUTHORIZATION_TYPE时为服务商的appid  | [可选] 
**OutDetailNo** | **string** | 商户系统内部区分转账批次单下不同转账明细单的唯一标识  | 
**DetailId** | **string** | 微信支付系统内部区分转账批次单下不同转账明细单的唯一标识  | 
**DetailStatus** | **string** | PROCESSING:转账中。正在处理中，转账结果尚未明确   SUCCESS:转账成功   FAIL:转账失败。需要确认失败原因后，再决定是否重新发起对该笔明细单的转账（并非整个转账批次单）  | 
**TransferAmount** | **int64** | 转账金额单位为“分”  | 
**TransferRemark** | **string** | 单条转账备注（微信用户会收到该备注），UTF8编码，最多允许32个字符  | 
**FailReason** | [**FailReasonType**](FailReasonType.md) | 如果转账失败则有失败原因 * &#x60;ACCOUNT_FROZEN&#x60; - 账户冻结 * &#x60;REAL_NAME_CHECK_FAIL&#x60; - 用户未实名 * &#x60;NAME_NOT_CORRECT&#x60; - 用户姓名校验失败 * &#x60;OPENID_INVALID&#x60; - Openid校验失败 * &#x60;TRANSFER_QUOTA_EXCEED&#x60; - 超过用户单笔收款额度 * &#x60;DAY_RECEIVED_QUOTA_EXCEED&#x60; - 超过用户单日收款额度 * &#x60;MONTH_RECEIVED_QUOTA_EXCEED&#x60; - 超过用户单月收款额度 * &#x60;DAY_RECEIVED_COUNT_EXCEED&#x60; - 超过用户单日收款次数 * &#x60;PRODUCT_AUTH_CHECK_FAIL&#x60; - 产品权限校验失败 * &#x60;OVERDUE_CLOSE&#x60; - 转账关闭 * &#x60;ID_CARD_NOT_CORRECT&#x60; - 用户身份证校验失败 * &#x60;ACCOUNT_NOT_EXIST&#x60; - 用户账户不存在 * &#x60;TRANSFER_RISK&#x60; - 转账存在风险 * &#x60;PAY_ROLL_CARD_ALREADY_LOGOUT&#x60; - 用户薪工卡已注销 * &#x60;PAY_ROLL_CARD_ALREADY_FROZEN&#x60; - 用户薪工卡已冻结 * &#x60;PAY_ROLL_CARD_UNAUTHORIZED&#x60; - 用户薪工卡未授权该商户 * &#x60;PAY_ROLL_CARD_USER_NOT_OPEN&#x60; - 用户未开通薪工卡  | [可选] 
**Openid** | **string** | 收款用户openid。如果转账特约商户授权类型是INFORMATION_AUTHORIZATION_TYPE，对应的是特约商户公众号下的openid;如果转账特约商户授权类型是FUND_AUTHORIZATION_TYPE，对应的是服务商商户公众号下的openid。  | 
**Username** | ***os.File** | 收款方姓名。采用标准RSA算法，公钥由微信侧提供  | 
**InitiateTime** | **time.Time** | 转账发起的时间，按照使用rfc3339所定义的格式，格式为YYYY-MM-DDThh:mm:ss+TIMEZONE  | 
**UpdateTime** | **time.Time** | 明细最后一次状态变更的时间，按照使用rfc3339所定义的格式，格式为YYYY-MM-DDThh:mm:ss+TIMEZONE  | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


