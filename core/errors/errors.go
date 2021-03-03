// Package errors 微信支付api v3 go http-client 错误类
package errors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

//Error 自定义的错误类型
type Error struct {
	StatusCode int           `json:"status_code"` // http 回包的状态码信息
	Code       string        `json:"code"`        // 将回包中body解析后的错误码信息，仅回包不符合预期时存在
	Message    string        `json:"message"`     // 将回包中body解析后的message信息，仅回包不符合预期时存在
	Details    []interface{} `json:"details"`     // 将回包中body解析出的detail信息，仅回包不符合预期时存在
	Body       string        // http回包中的body信息
	Header     http.Header   // http回包中的header信息
}

//Error 返回自定义错误类型的字符串内容
func (e *Error) Error() string {
	var buf bytes.Buffer
	_, _ = fmt.Fprintf(&buf, "wechat pay sdk: error HTTP response\nStatusCode: %d \nCode: %s ", e.StatusCode, e.Code)
	if e.Message != "" {
		_, _ = fmt.Fprintf(&buf, "\nMessage: %s", e.Message)
	}
	if len(e.Details) > 0 {
		var detailBuf bytes.Buffer
		enc := json.NewEncoder(&detailBuf)
		enc.SetIndent("", "  ")
		if err := enc.Encode(e.Details); err == nil {
			_, _ = fmt.Fprint(&buf, "\nDetails:")
			_, _ = fmt.Fprintf(&buf, "\n%s", detailBuf.String())
		}
	}
	return buf.String()
}
