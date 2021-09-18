# 升级指南

## v0.2.2

版本 `v0.2.2` 中我们包含了一个重大更新内容：将现有接口SDK中的整型参数统一为`int64`。
这一行为的目的是规范SDK中对整型参数的实现，避免因为未来可能的`int32 -> int64`的字段升级导致大规模兼容性问题。

此次升级会导致 `payments` 下4个API接口的SDK的兼容性问题，建议开发者以如下方式对自己的代码进行更新。

### 1. 升级依赖
1. 在你的模块目录下执行 `go get -u github.com/wechatpay-apiv3/wechatpay-go@v0.2.7` 升级依赖。
2. （正常情况下该步骤会自动完成）修改模块 `go.mod` 文件中依赖的 `github.com/wechatpay-apiv3/wechatpay-go` 至 `v0.2.7`版本。

### 2. 定位需要修改的代码
在项目根目录下执行`go build ./...`可以递归检查代码中的编译错误，即可快速定位到需要修改的代码。

### 3. 对请求构建代码进行更新
对于请求 `payments` 接口的数据，可以在设置参数时使用`int64(xxx)`进行类型转换。当然也可以将请求链路上的类型从`int32`更新为`int64`。
```go
req := jsapi.PrepayRequest{}
// 升级前
req.Amount = &jsapi.Amount{
	Currency: core.String("CNY"),
	Total:    &totalInt32,
}
// 升级后
totalInt64 := int64(totalInt32)
req.Amount = &jsapi.Amount{
	Currency: core.String("CNY"),
	Total:    &totalInt64,
}
```

### 4. 对应答处理代码进行更新
对于应答结果的处理，我们不建议将返回结果中的`int64`强制类型转换为`int32`，而是建议将后续处理链路中的类型从`int32`更新为`int64`。
这样变更可能会更复杂，但是安全性更好，避免因为数据溢出导致错误。
```go
// 升级前
func GetTransactionTotal(resp *payments.Transaction) int32 {
	return *resp.Amount.Total
}
// 升级后
func GetTransactionTotal(resp *payments.Transaction) int64 {
	return *resp.Amount.Total
}
```

### 5. 更新你的测试用例代码并测试
如果你有针对 `payments` 编写测试用例，你可能需要对测试用例代码进行更新，并重新测试确保一切正常。
