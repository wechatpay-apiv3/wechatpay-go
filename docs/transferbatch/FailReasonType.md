# FailReasonType

  ACCOUNT_FROZEN - 账户冻结，该用户账户被冻结  REAL_NAME_CHECK_FAIL - 用户未实名，收款人未实名认证，需要用户完成微信实名认证  NAME_NOT_CORRECT - 用户姓名校验失败，收款人姓名校验不通过，请核实信息  OPENID_INVALID - Openid校验失败，Openid格式错误或者不属于商家公众账号  TRANSFER_QUOTA_EXCEED - 超过用户单笔收款额度，超过用户单笔收款额度，核实产品设置是否准确  DAY_RECEIVED_QUOTA_EXCEED - 超过用户单日收款额度，超过用户单日收款额度，核实产品设置是否准确  MONTH_RECEIVED_QUOTA_EXCEED - 超过用户单月收款额度，超过用户单月收款额度，核实产品设置是否准确  DAY_RECEIVED_COUNT_EXCEED - 超过用户单日收款次数，超过用户单日收款次数，核实产品设置是否准确  PRODUCT_AUTH_CHECK_FAIL - 产品权限校验失败，未开通该权限或权限被冻结，请核实产品权限状态  OVERDUE_CLOSE - 转账关闭，超过系统重试期，系统自动关闭  ID_CARD_NOT_CORRECT - 用户身份证校验失败，收款人身份证校验不通过，请核实信息  ACCOUNT_NOT_EXIST - 用户账户不存在，该用户账户不存在  TRANSFER_RISK - 转账存在风险，该笔转账可能存在风险，已被微信拦截  REALNAME_ACCOUNT_RECEIVED_QUOTA_EXCEED - 用户账户收款受限，请引导用户在微信支付查看详情

## 枚举


* `ACCOUNT_FROZEN` (value: `"ACCOUNT_FROZEN"`)

* `REAL_NAME_CHECK_FAIL` (value: `"REAL_NAME_CHECK_FAIL"`)

* `NAME_NOT_CORRECT` (value: `"NAME_NOT_CORRECT"`)

* `OPENID_INVALID` (value: `"OPENID_INVALID"`)

* `TRANSFER_QUOTA_EXCEED` (value: `"TRANSFER_QUOTA_EXCEED"`)

* `DAY_RECEIVED_QUOTA_EXCEED` (value: `"DAY_RECEIVED_QUOTA_EXCEED"`)

* `MONTH_RECEIVED_QUOTA_EXCEED` (value: `"MONTH_RECEIVED_QUOTA_EXCEED"`)

* `DAY_RECEIVED_COUNT_EXCEED` (value: `"DAY_RECEIVED_COUNT_EXCEED"`)

* `PRODUCT_AUTH_CHECK_FAIL` (value: `"PRODUCT_AUTH_CHECK_FAIL"`)

* `OVERDUE_CLOSE` (value: `"OVERDUE_CLOSE"`)

* `ID_CARD_NOT_CORRECT` (value: `"ID_CARD_NOT_CORRECT"`)

* `ACCOUNT_NOT_EXIST` (value: `"ACCOUNT_NOT_EXIST"`)

* `TRANSFER_RISK` (value: `"TRANSFER_RISK"`)

* `REALNAME_ACCOUNT_RECEIVED_QUOTA_EXCEED` (value: `"REALNAME_ACCOUNT_RECEIVED_QUOTA_EXCEED"`)


[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


