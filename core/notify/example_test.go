package notify_test

import (
	"context"
	"fmt"
	"net/http"

	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
)

func ExampleHandler_ParseNotifyRequest_transaction() {
	var handler notify.Handler
	var request *http.Request

	content := new(payments.Transaction)
	notifyReq, err := handler.ParseNotifyRequest(context.Background(), request, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 处理通知内容
	fmt.Println(notifyReq.Summary)
	fmt.Println(content)
}

func ExampleHandler_ParseNotifyRequest_no_model() {
	var handler notify.Handler
	var request *http.Request

	content := make(map[string]interface{})
	notifyReq, err := handler.ParseNotifyRequest(context.Background(), request, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 处理通知内容
	fmt.Println(notifyReq.Summary)
	fmt.Println(content)
}
