# ModifyStockInfoRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**StockId** | **string** | 批次号 | 
**CustomEntrance** | [**ModifyCustomEntrance**](ModifyCustomEntrance.md) | 卡详情页面，可选择多种入口引导用户 | [可选] 
**StockName** | **string** | 批次名称，字数上限为21个，一个中文汉字/英文字母/数字均占用一个字数。 注：该字段暂不支持修改 | [可选] 
**Comment** | **string** | 仅配置商户可见，用于自定义信息。字数上限为20个，一个中文汉字/英文字母/数字均占用一个字数。 | [可选] 
**GoodsName** | **string** | 用来描述批次在哪些商品可用，会显示在微信卡包中。字数上限为15个，一个中文汉字/英文字母/数字均占用一个字数。 | [可选] 
**OutRequestNo** | **string** | 商户修改批次凭据号（格式：商户ID+日期+流水号），商户侧需保持唯一性。 | 
**DisplayPatternInfo** | [**DisplayPatternInfo**](DisplayPatternInfo.md) |  | [可选] 
**CouponUseRule** | [**ModifyCouponUseRule**](ModifyCouponUseRule.md) | 券核销相关规则 | [可选] 
**StockSendRule** | [**ModifyStockSendRule**](ModifyStockSendRule.md) | 券发放相关规则 | [可选] 
**NotifyConfig** | [**NotifyConfig**](NotifyConfig.md) | 事件回调通知商户的配置 | [可选] 
**Subsidy** | **bool** | 该批次发放的券是否允许进行补差。只允许从false改为true，不支持从true改为false 注：该字段暂未开放 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


