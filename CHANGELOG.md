# Changelog

## [0.2.1] - 2021-06-25

### Added

+ 平台证书下载器与自动下载管理器
+ 回调通知验签、解密器
+ 敏感字段自动加解密库
+ 提供拉起支付签名计算接口
+ 一键构建Client的签名器、验签器的复合Option
+ 平台证书下载命令行工具

### Changed

+ `core.WithXXX` 等方法移动至 `option.WithXXX`

### Removed

+ 移除 `core.Client` 自定义 Header 选项，如需设置 Header 可使用 `client.Request` 接口发起请求
+ 移除 `core.Client` 自定义 Timeout 选项，如需设置 Timeout 可设置自定义 HTTPClient

## [0.2.0] - 2021-05-31

### Added

+ 增加微信支付支付API的SDK
+ 增加文件上传API的SDK

## [0.1.0] - 2020-03-03

+ Initial version

