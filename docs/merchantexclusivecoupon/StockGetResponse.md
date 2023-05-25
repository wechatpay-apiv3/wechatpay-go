# StockGetResponse

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**StockName** | **string** | 批次名称，字数上限为21个，一个中文汉字/英文字母/数字均占用一个字数。 | 
**BelongMerchant** | **string** | 批次是归属于哪个商户 注： 普通直连模式，该参数为直连商户号； 服务商模式，该参数为子商户号； 间连模式，该参数为子商户号。 | 
**Comment** | **string** | 仅配置商户可见，用于自定义信息。字数上限为20个，一个中文汉字/英文字母/数字均占用一个字数。 | [可选] 
**GoodsName** | **string** | 用来描述批次在哪些商品可用，会显示在微信卡包中。字数上限为15个，一个中文汉字/英文字母/数字均占用一个字数。 | 
**StockType** | [**BusiFavorStockType**](BusiFavorStockType.md) | 批次类型 | 
**CouponUseRule** | [**CouponUseRule**](CouponUseRule.md) | 核销规则 | 
**StockSendRule** | [**StockSendRule**](StockSendRule.md) | 发放规则 | 
**CustomEntrance** | [**CustomEntrance**](CustomEntrance.md) | 卡详情页面，可选择多种入口引导用户 | [可选] 
**DisplayPatternInfo** | [**DisplayPatternInfo**](DisplayPatternInfo.md) | 创建批次时的样式信息 | [可选] 
**StockState** | [**StockStatus**](StockStatus.md) | 批次状态 | 
**CouponCodeMode** | [**CouponCodeMode**](CouponCodeMode.md) | 券code模式 | 
**StockId** | **string** | 批次唯一标识 | 
**CouponCodeCount** | [**CouponCodeCount**](CouponCodeCount.md) | 当且仅当coupon_code_mode(券code模式)为MERCHANT_UPLOAD(商户上传自定义code)模式时，返回该字段；返回内容为商户上传code的数量信息 | [可选] 
**NotifyConfig** | [**NotifyConfig**](NotifyConfig.md) | 事件回调通知商户的配置 | [可选] 
**SendCountInformation** | [**SendCount**](SendCount.md) | 批次发放情况 | [可选] 
**Subsidy** | **bool** | 该批次发放的券是否允许进行补差 注：该字段暂未开放 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


