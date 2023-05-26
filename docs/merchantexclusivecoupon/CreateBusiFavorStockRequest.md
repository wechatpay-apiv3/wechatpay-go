# CreateBusiFavorStockRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**StockName** | **string** | 批次名称，字数上限为21个，一个中文汉字/英文字母/数字均占用一个字数。 | 
**BelongMerchant** | **string** | 批次归属于哪个商户。  注： 普通直连模式，该参数为直连商户号； 服务商模式，该参数为子商户号； 间连模式，该参数为子商户号。 | 
**Comment** | **string** | 仅配置商户可见，用于自定义信息。字数上限为20个，一个中文汉字/英文字母/数字均占用一个字数。 | [可选] 
**GoodsName** | **string** | 用来描述批次在哪些商品可用，会显示在微信卡包中。字数上限为15个，一个中文汉字/英文字母/数字均占用一个字数。 | 
**StockType** | [**BusiFavorStockType**](BusiFavorStockType.md) | 批次类型 | 
**CouponUseRule** | [**CouponUseRule**](CouponUseRule.md) | 券核销相关规则 | 
**StockSendRule** | [**StockSendRule**](StockSendRule.md) | 券发放相关规则 | 
**OutRequestNo** | **string** | 商户创建批次凭据号（格式：商户ID+日期+流水号），商户侧需保持唯一性 | 
**CustomEntrance** | [**CustomEntrance**](CustomEntrance.md) | 卡详情页面，可选择多种入口引导用户 | [可选] 
**DisplayPatternInfo** | [**DisplayPatternInfo**](DisplayPatternInfo.md) | 创建批次时的样式信息。 | [可选] 
**CouponCodeMode** | [**CouponCodeMode**](CouponCodeMode.md) | 特殊规则： 1、券code模式为WECHATPAY\\_MODE时，是微信自动分配券code，商户不需要预存code；适用于多种场景 2、券code模式为MERCHANT\\_API时，无需调用上传预存code接口，调用发券接口时需指定券code；更多用在商家自有流量场景（例如：商家自有小程序、H5网页等） 3、券code模式为MERCHANT\\_UPLOAD，需要调用上传预存code接口上传code，调用发券接口时无需指定code；更多适用在微信支付平台流量场景（例如：支付有礼、支付有优惠等） | 
**NotifyConfig** | [**NotifyConfig**](NotifyConfig.md) | 事件回调通知商户的配置 | [可选] 
**Subsidy** | **bool** | 该批次发放的券是否允许进行补差，默认为false 注：该字段暂未开放 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


