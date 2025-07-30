# TransferState

转账状态枚举值

## 可能的值

* `ACCEPTED` - 转账已受理
* `PROCESSING` - 转账锁定资金中。如果一直停留在该状态，建议检查账户余额是否足够，如余额不足，可充值后再原单重试
* `WAIT_USER_CONFIRM` - 待收款用户确认，可拉起微信收款确认页面进行收款确认
* `TRANSFERING` - 转账中，可拉起微信收款确认页面再次重试确认收款
* `SUCCESS` - 转账成功
* `FAIL` - 转账失败
* `CANCELING` - 商户撤销请求受理成功，该笔转账正在撤销中
* `CANCELLED` - 转账撤销完成

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)

