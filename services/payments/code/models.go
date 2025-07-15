package code

import (
	"encoding/json"
	"fmt"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"time"
)

// Amount
type Amount struct {
	// 订单总金额，单位为分
	Total *int64 `json:"total"`
	// CNY：人民币，境内商户号仅支持人民币。
	Currency *string `json:"currency,omitempty"`
}

func (o Amount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Total == nil {
		return nil, fmt.Errorf("field `Total` is required and must be specified in Amount")
	}
	toSerialize["total"] = o.Total

	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	return json.Marshal(toSerialize)
}

func (o Amount) String() string {
	var ret string
	if o.Total == nil {
		ret += "Total:<nil>, "
	} else {
		ret += fmt.Sprintf("Total:%v, ", *o.Total)
	}

	if o.Currency == nil {
		ret += "Currency:<nil>"
	} else {
		ret += fmt.Sprintf("Currency:%v", *o.Currency)
	}

	return fmt.Sprintf("Amount{%s}", ret)
}

func (o Amount) Clone() *Amount {
	ret := Amount{}

	if o.Total != nil {
		ret.Total = new(int64)
		*ret.Total = *o.Total
	}

	if o.Currency != nil {
		ret.Currency = new(string)
		*ret.Currency = *o.Currency
	}

	return &ret
}

// 【支付者】 支付者信息。
type AuthPayer struct {
	//【授权码】 付款码支付授权码，即用户打开微信钱包显示的码。
	AuthCode *string `json:"auth_code"`
}

func (o AuthPayer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.AuthCode != nil {
		toSerialize["auth_code"] = o.AuthCode
	}
	return json.Marshal(toSerialize)
}

func (o AuthPayer) String() string {
	var ret string
	if o.AuthCode == nil {
		ret += "auth_code:<nil>"
	} else {
		ret += fmt.Sprintf("auth_code:%v", *o.AuthCode)
	}
	return fmt.Sprintf("Payer{%s}", ret)
}

func (o AuthPayer) Clone() *AuthPayer {
	ret := AuthPayer{}

	if o.AuthCode != nil {
		ret.AuthCode = new(string)
		*ret.AuthCode = *o.AuthCode
	}

	return &ret
}

// CloseOrderRequest
type CloseOrderRequest struct {
	OutTradeNo *string `json:"out_trade_no"`
	// 直连商户号
	Mchid *string `json:"mchid"`
	//应用appid
	Appid *string `json:"appid"`
}

func (o CloseOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.OutTradeNo == nil {
		return nil, fmt.Errorf("field `OutTradeNo` is required and must be specified in CloseOrderRequest")
	}
	toSerialize["out_trade_no"] = o.OutTradeNo

	if o.Mchid == nil {
		return nil, fmt.Errorf("field `Mchid` is required and must be specified in CloseOrderRequest")
	}
	toSerialize["mchid"] = o.Mchid
	if o.Appid == nil {
		return nil, fmt.Errorf("field `Appid` is required and must be specified in CloseOrderRequest")
	}
	toSerialize["appid"] = o.Appid

	return json.Marshal(toSerialize)
}

func (o CloseOrderRequest) String() string {
	var ret string

	if o.OutTradeNo == nil {
		ret += "OutTradeNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutTradeNo:%v, ", *o.OutTradeNo)
	}

	if o.Mchid == nil {
		ret += "Mchid:<nil>"
	} else {
		ret += fmt.Sprintf("Mchid:%v", *o.Mchid)
	}
	if o.Appid == nil {
		ret += "Appid:<nil>"
	} else {
		ret += fmt.Sprintf("Appid:%v", *o.Appid)
	}
	return fmt.Sprintf("CloseOrderRequest{%s}", ret)
}

func (o CloseOrderRequest) Clone() *CloseOrderRequest {
	ret := CloseOrderRequest{}

	if o.OutTradeNo != nil {
		ret.OutTradeNo = new(string)
		*ret.OutTradeNo = *o.OutTradeNo
	}

	if o.Mchid != nil {
		ret.Mchid = new(string)
		*ret.Mchid = *o.Mchid
	}
	if o.Appid != nil {
		ret.Appid = new(string)
		*ret.Appid = *o.Appid
	}
	return &ret
}

// CloseRequest
type CloseRequest struct {
	// 直连商户号
	Mchid *string `json:"mchid"`
	//应用AppID。
	Appid *string `json:"appid"`
}

func (o CloseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Mchid == nil {
		return nil, fmt.Errorf("field `Mchid` is required and must be specified in CloseRequest")
	}
	toSerialize["mchid"] = o.Mchid
	if o.Appid == nil {
		return nil, fmt.Errorf("field `Appid` is required and must be specified in CloseRequest")
	}
	toSerialize["appid"] = o.Appid
	return json.Marshal(toSerialize)
}

func (o CloseRequest) String() string {
	var ret string
	if o.Mchid == nil {
		ret += "Mchid:<nil>"
	} else {
		ret += fmt.Sprintf("Mchid:%v", *o.Mchid)
	}
	if o.Appid == nil {
		ret += "Appid:<nil>, "
	} else {
		ret += fmt.Sprintf("Appid:%v, ", *o.Appid)
	}
	return fmt.Sprintf("CloseRequest{%s}", ret)
}

func (o CloseRequest) Clone() *CloseRequest {
	ret := CloseRequest{}

	if o.Mchid != nil {
		ret.Mchid = new(string)
		*ret.Mchid = *o.Mchid
	}
	if o.Appid != nil {
		ret.Appid = new(string)
		*ret.Appid = *o.Appid
	}
	return &ret
}

// Detail 优惠功能
type Detail struct {
	// 1.商户侧一张小票订单可能被分多次支付，订单原价用于记录整张小票的交易金额。 2.当订单原价与支付金额不相等，则不享受优惠。 3.该字段主要用于防止同一张小票分多次支付，以享受多次优惠的情况，正常支付订单不必上传此参数。
	CostPrice *int64 `json:"cost_price,omitempty"`
	// 商家小票ID。
	InvoiceId   *string       `json:"invoice_id,omitempty"`
	GoodsDetail []GoodsDetail `json:"goods_detail,omitempty"`
}

func (o Detail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.CostPrice != nil {
		toSerialize["cost_price"] = o.CostPrice
	}

	if o.InvoiceId != nil {
		toSerialize["invoice_id"] = o.InvoiceId
	}

	if o.GoodsDetail != nil {
		toSerialize["goods_detail"] = o.GoodsDetail
	}
	return json.Marshal(toSerialize)
}

func (o Detail) String() string {
	var ret string
	if o.CostPrice == nil {
		ret += "CostPrice:<nil>, "
	} else {
		ret += fmt.Sprintf("CostPrice:%v, ", *o.CostPrice)
	}

	if o.InvoiceId == nil {
		ret += "InvoiceId:<nil>, "
	} else {
		ret += fmt.Sprintf("InvoiceId:%v, ", *o.InvoiceId)
	}

	ret += fmt.Sprintf("GoodsDetail:%v", o.GoodsDetail)

	return fmt.Sprintf("Detail{%s}", ret)
}

func (o Detail) Clone() *Detail {
	ret := Detail{}

	if o.CostPrice != nil {
		ret.CostPrice = new(int64)
		*ret.CostPrice = *o.CostPrice
	}

	if o.InvoiceId != nil {
		ret.InvoiceId = new(string)
		*ret.InvoiceId = *o.InvoiceId
	}

	if o.GoodsDetail != nil {
		ret.GoodsDetail = make([]GoodsDetail, len(o.GoodsDetail))
		for i, item := range o.GoodsDetail {
			ret.GoodsDetail[i] = *item.Clone()
		}
	}

	return &ret
}

// GoodsDetail
type GoodsDetail struct {
	// 由半角的大小写字母、数字、中划线、下划线中的一种或几种组成。
	MerchantGoodsId *string `json:"merchant_goods_id"`
	// 微信支付定义的统一商品编号（没有可不传）。
	WechatpayGoodsId *string `json:"wechatpay_goods_id,omitempty"`
	// 商品的实际名称。
	GoodsName *string `json:"goods_name,omitempty"`
	// 用户购买的数量。
	Quantity *int64 `json:"quantity"`
	// 商品单价，单位为分。
	UnitPrice *int64 `json:"unit_price"`
}

func (o GoodsDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.MerchantGoodsId == nil {
		return nil, fmt.Errorf("field `MerchantGoodsId` is required and must be specified in GoodsDetail")
	}
	toSerialize["merchant_goods_id"] = o.MerchantGoodsId

	if o.WechatpayGoodsId != nil {
		toSerialize["wechatpay_goods_id"] = o.WechatpayGoodsId
	}

	if o.GoodsName != nil {
		toSerialize["goods_name"] = o.GoodsName
	}

	if o.Quantity == nil {
		return nil, fmt.Errorf("field `Quantity` is required and must be specified in GoodsDetail")
	}
	toSerialize["quantity"] = o.Quantity

	if o.UnitPrice == nil {
		return nil, fmt.Errorf("field `UnitPrice` is required and must be specified in GoodsDetail")
	}
	toSerialize["unit_price"] = o.UnitPrice
	return json.Marshal(toSerialize)
}

func (o GoodsDetail) String() string {
	var ret string
	if o.MerchantGoodsId == nil {
		ret += "MerchantGoodsId:<nil>, "
	} else {
		ret += fmt.Sprintf("MerchantGoodsId:%v, ", *o.MerchantGoodsId)
	}

	if o.WechatpayGoodsId == nil {
		ret += "WechatpayGoodsId:<nil>, "
	} else {
		ret += fmt.Sprintf("WechatpayGoodsId:%v, ", *o.WechatpayGoodsId)
	}

	if o.GoodsName == nil {
		ret += "GoodsName:<nil>, "
	} else {
		ret += fmt.Sprintf("GoodsName:%v, ", *o.GoodsName)
	}

	if o.Quantity == nil {
		ret += "Quantity:<nil>, "
	} else {
		ret += fmt.Sprintf("Quantity:%v, ", *o.Quantity)
	}

	if o.UnitPrice == nil {
		ret += "UnitPrice:<nil>"
	} else {
		ret += fmt.Sprintf("UnitPrice:%v", *o.UnitPrice)
	}

	return fmt.Sprintf("GoodsDetail{%s}", ret)
}

func (o GoodsDetail) Clone() *GoodsDetail {
	ret := GoodsDetail{}

	if o.MerchantGoodsId != nil {
		ret.MerchantGoodsId = new(string)
		*ret.MerchantGoodsId = *o.MerchantGoodsId
	}

	if o.WechatpayGoodsId != nil {
		ret.WechatpayGoodsId = new(string)
		*ret.WechatpayGoodsId = *o.WechatpayGoodsId
	}

	if o.GoodsName != nil {
		ret.GoodsName = new(string)
		*ret.GoodsName = *o.GoodsName
	}

	if o.Quantity != nil {
		ret.Quantity = new(int64)
		*ret.Quantity = *o.Quantity
	}

	if o.UnitPrice != nil {
		ret.UnitPrice = new(int64)
		*ret.UnitPrice = *o.UnitPrice
	}

	return &ret
}

// PrepayRequest
type PrepayRequest struct {
	// 公众号ID
	Appid *string `json:"appid"`
	// 直连商户号
	Mchid *string `json:"mchid"`
	// 商品描述
	Description *string `json:"description"`
	// 商户订单号
	OutTradeNo *string `json:"out_trade_no"`
	// 订单失效时间，格式为rfc3339格式
	TimeExpire *time.Time `json:"time_expire,omitempty"`
	// 附加数据
	Attach *string `json:"attach,omitempty"`
	// 有效性：1. HTTPS；2. 不允许携带查询串。
	NotifyUrl *string `json:"notify_url"`
	// 商品标记，代金券或立减优惠功能的参数。
	GoodsTag *string `json:"goods_tag,omitempty"`
	// 传入true时，支付成功消息和支付详情页将出现开票入口。需要在微信支付商户平台或微信公众平台开通电子发票功能，传此字段才可生效。
	SupportFapiao *bool       `json:"support_fapiao,omitempty"`
	Payer         *AuthPayer  `json:"payer,omitempty"`
	Amount        *Amount     `json:"amount"`
	Detail        *Detail     `json:"detail,omitempty"`
	SceneInfo     *SceneInfo  `json:"scene_info,omitempty"`
	SettleInfo    *SettleInfo `json:"settle_info,omitempty"`
}

func (o PrepayRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Appid == nil {
		return nil, fmt.Errorf("field `Appid` is required and must be specified in PrepayRequest")
	}
	toSerialize["appid"] = o.Appid

	if o.Mchid == nil {
		return nil, fmt.Errorf("field `Mchid` is required and must be specified in PrepayRequest")
	}
	toSerialize["mchid"] = o.Mchid

	if o.Description == nil {
		return nil, fmt.Errorf("field `Description` is required and must be specified in PrepayRequest")
	}
	toSerialize["description"] = o.Description

	if o.OutTradeNo == nil {
		return nil, fmt.Errorf("field `OutTradeNo` is required and must be specified in PrepayRequest")
	}
	toSerialize["out_trade_no"] = o.OutTradeNo

	if o.TimeExpire != nil {
		toSerialize["time_expire"] = o.TimeExpire.Format(time.RFC3339)
	}

	if o.Attach != nil {
		toSerialize["attach"] = o.Attach
	}

	if o.NotifyUrl == nil {
		return nil, fmt.Errorf("field `NotifyUrl` is required and must be specified in PrepayRequest")
	}
	toSerialize["notify_url"] = o.NotifyUrl

	if o.GoodsTag != nil {
		toSerialize["goods_tag"] = o.GoodsTag
	}

	if o.Payer == nil {
		return nil, fmt.Errorf("field `Payer` is required and must be specified in PrepayRequest")
	}
	toSerialize["payer"] = o.Payer

	if o.SupportFapiao != nil {
		toSerialize["support_fapiao"] = o.SupportFapiao
	}

	if o.Amount == nil {
		return nil, fmt.Errorf("field `Amount` is required and must be specified in PrepayRequest")
	}
	toSerialize["amount"] = o.Amount

	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}

	if o.SceneInfo != nil {
		toSerialize["scene_info"] = o.SceneInfo
	}

	if o.SettleInfo != nil {
		toSerialize["settle_info"] = o.SettleInfo
	}
	return json.Marshal(toSerialize)
}

func (o PrepayRequest) String() string {
	var ret string
	if o.Appid == nil {
		ret += "Appid:<nil>, "
	} else {
		ret += fmt.Sprintf("Appid:%v, ", *o.Appid)
	}

	if o.Mchid == nil {
		ret += "Mchid:<nil>, "
	} else {
		ret += fmt.Sprintf("Mchid:%v, ", *o.Mchid)
	}

	if o.Description == nil {
		ret += "Description:<nil>, "
	} else {
		ret += fmt.Sprintf("Description:%v, ", *o.Description)
	}

	if o.OutTradeNo == nil {
		ret += "OutTradeNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutTradeNo:%v, ", *o.OutTradeNo)
	}

	if o.TimeExpire == nil {
		ret += "TimeExpire:<nil>, "
	} else {
		ret += fmt.Sprintf("TimeExpire:%v, ", *o.TimeExpire)
	}

	if o.Attach == nil {
		ret += "Attach:<nil>, "
	} else {
		ret += fmt.Sprintf("Attach:%v, ", *o.Attach)
	}

	if o.NotifyUrl == nil {
		ret += "NotifyUrl:<nil>, "
	} else {
		ret += fmt.Sprintf("NotifyUrl:%v, ", *o.NotifyUrl)
	}

	if o.GoodsTag == nil {
		ret += "GoodsTag:<nil>, "
	} else {
		ret += fmt.Sprintf("GoodsTag:%v, ", *o.GoodsTag)
	}

	ret += fmt.Sprintf("Payer:%v, ", o.Payer)

	if o.SupportFapiao == nil {
		ret += "SupportFapiao:<nil>, "
	} else {
		ret += fmt.Sprintf("SupportFapiao:%v, ", *o.SupportFapiao)
	}

	ret += fmt.Sprintf("Amount:%v, ", o.Amount)

	ret += fmt.Sprintf("Detail:%v, ", o.Detail)

	ret += fmt.Sprintf("SceneInfo:%v, ", o.SceneInfo)

	ret += fmt.Sprintf("SettleInfo:%v", o.SettleInfo)

	return fmt.Sprintf("PrepayRequest{%s}", ret)
}

func (o PrepayRequest) Clone() *PrepayRequest {
	ret := PrepayRequest{}

	if o.Appid != nil {
		ret.Appid = new(string)
		*ret.Appid = *o.Appid
	}

	if o.Mchid != nil {
		ret.Mchid = new(string)
		*ret.Mchid = *o.Mchid
	}

	if o.Description != nil {
		ret.Description = new(string)
		*ret.Description = *o.Description
	}

	if o.OutTradeNo != nil {
		ret.OutTradeNo = new(string)
		*ret.OutTradeNo = *o.OutTradeNo
	}

	if o.TimeExpire != nil {
		ret.TimeExpire = new(time.Time)
		*ret.TimeExpire = *o.TimeExpire
	}

	if o.Attach != nil {
		ret.Attach = new(string)
		*ret.Attach = *o.Attach
	}

	if o.NotifyUrl != nil {
		ret.NotifyUrl = new(string)
		*ret.NotifyUrl = *o.NotifyUrl
	}

	if o.GoodsTag != nil {
		ret.GoodsTag = new(string)
		*ret.GoodsTag = *o.GoodsTag
	}

	if o.Payer != nil {
		ret.Payer = o.Payer.Clone()
	}

	if o.SupportFapiao != nil {
		ret.SupportFapiao = new(bool)
		*ret.SupportFapiao = *o.SupportFapiao
	}

	if o.Amount != nil {
		ret.Amount = o.Amount.Clone()
	}

	if o.Detail != nil {
		ret.Detail = o.Detail.Clone()
	}

	if o.SceneInfo != nil {
		ret.SceneInfo = o.SceneInfo.Clone()
	}

	if o.SettleInfo != nil {
		ret.SettleInfo = o.SettleInfo.Clone()
	}

	return &ret
}

// PrepayResponse
type PrepayResponse struct {
	// 商户应用ID
	Appid *string `json:"appid"`
	// 直连商户号
	Mchid *string `json:"mchid"`
	// 商户订单号
	OutTradeNo *string `json:"out_trade_no"`
	//【微信支付订单号】 微信支付系统生成的订单号。
	TransactionId *string `json:"transaction_id,omitempty"`
	//易类型，如下： MICROPAY：付款码支付  FACEPAY：刷脸支付
	TradeType *string `json:"trade_type,omitempty"`
	//【银行类型】 银行类型，采用字符串类型的银行标识。
	BankType *string `json:"bank_type,omitempty"`
	// 支付完成时间
	SuccessTime *string `json:"success_time,omitempty"`
	//【交易状态】 交易状态，如下：
	//* SUCCESS：支付成功
	//* REFUND：转入退款
	//* NOTPAY：未支付
	//* REVOKED：已撤销
	//* USERPAYING：用户支付中
	//* PAYERROR：支付失败
	TradeState *string `json:"trade_state"`
	//交易状态描述
	TradeStateDesc *string `json:"trade_state_desc,omitempty"`
	// 附加数据，在返回消息和查单API中原样返回，可作为自定义参数使用，实际情况下只有支付完成状态才会返回该字段
	Attach *string `json:"attach,omitempty"`
	Payer  *Payer  `json:"payer"`
	// 传入true时，支付成功消息和支付详情页将出现开票入口。需要在微信支付商户平台或微信公众平台开通电子发票功能，传此字段才可生效。
	Amount *Amount `json:"amount"`
	//优惠信息
	PromotionDetail *payments.PromotionDetail `json:"promotion_detail,omitempty"`
}

func (o PrepayResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Appid == nil {
		return nil, fmt.Errorf("field `Appid` is required and must be specified in PrepayRequest")
	}
	toSerialize["appid"] = o.Appid

	if o.Mchid == nil {
		return nil, fmt.Errorf("field `Mchid` is required and must be specified in PrepayRequest")
	}
	toSerialize["mchid"] = o.Mchid

	if o.OutTradeNo == nil {
		return nil, fmt.Errorf("field `OutTradeNo` is required and must be specified in PrepayRequest")
	}
	toSerialize["out_trade_no"] = o.OutTradeNo

	toSerialize["transaction_id"] = o.TransactionId

	if o.TradeType != nil {
		toSerialize["trade_type"] = o.TradeType
	}

	if o.Attach != nil {
		toSerialize["attach"] = o.Attach
	}

	if o.TradeStateDesc != nil {
		toSerialize["trade_state_desc"] = o.TradeStateDesc
	}

	if o.BankType != nil {
		toSerialize["bank_type"] = o.BankType
	}

	if o.Payer != nil {
		toSerialize["payer"] = o.Payer
	}

	if o.SuccessTime != nil {
		toSerialize["success_time"] = o.SuccessTime
	}

	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}

	if o.PromotionDetail != nil {
		toSerialize["promotion_detail"] = o.PromotionDetail
	}

	return json.Marshal(toSerialize)
}

func (o PrepayResponse) String() string {
	var ret string
	if o.Appid == nil {
		ret += "Appid:<nil>, "
	} else {
		ret += fmt.Sprintf("Appid:%v, ", *o.Appid)
	}

	if o.Mchid == nil {
		ret += "Mchid:<nil>, "
	} else {
		ret += fmt.Sprintf("Mchid:%v, ", *o.Mchid)
	}

	if o.OutTradeNo == nil {
		ret += "OutTradeNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutTradeNo:%v, ", *o.OutTradeNo)
	}

	if o.SuccessTime == nil {
		ret += "SuccessTime:<nil>, "
	} else {
		ret += fmt.Sprintf("SuccessTime:%v, ", *o.SuccessTime)
	}

	if o.Attach == nil {
		ret += "Attach:<nil>, "
	} else {
		ret += fmt.Sprintf("Attach:%v, ", *o.Attach)
	}

	ret += fmt.Sprintf("Payer:%v, ", o.Payer)

	if o.PromotionDetail == nil {
		ret += "SupportFapiao:<nil>, "
	} else {
		ret += fmt.Sprintf("PromotionDetail:%v, ", *o.PromotionDetail)
	}

	ret += fmt.Sprintf("Amount:%v, ", o.Amount)

	ret += fmt.Sprintf("TradeState:%v, ", o.TradeState)

	if o.TradeStateDesc == nil {
		ret += "TradeStateDesc:<nil>, "
	} else {
		ret += fmt.Sprintf("TradeStateDesc:%v, ", *o.TradeStateDesc)
	}

	return fmt.Sprintf("PrepayResponse{%s}", ret)
}

func (o PrepayResponse) Clone() *PrepayResponse {
	ret := PrepayResponse{}

	if o.Appid != nil {
		ret.Appid = new(string)
		*ret.Appid = *o.Appid
	}

	if o.Mchid != nil {
		ret.Mchid = new(string)
		*ret.Mchid = *o.Mchid
	}

	if o.TransactionId != nil {
		ret.TransactionId = new(string)
		*ret.TransactionId = *o.TransactionId
	}

	if o.OutTradeNo != nil {
		ret.OutTradeNo = new(string)
		*ret.OutTradeNo = *o.OutTradeNo
	}

	if o.SuccessTime != nil {
		ret.SuccessTime = new(string)
		*ret.SuccessTime = *o.SuccessTime
	}

	if o.Attach != nil {
		ret.Attach = new(string)
		*ret.Attach = *o.Attach
	}

	if o.TradeType != nil {
		ret.TradeType = new(string)
		*ret.TradeType = *o.TradeType
	}

	if o.TradeState != nil {
		ret.TradeState = new(string)
		*ret.TradeState = *o.TradeState
	}

	if o.Payer != nil {
		ret.Payer = o.Payer.Clone()
	}

	if o.TradeStateDesc != nil {
		ret.TradeStateDesc = new(string)
		*ret.TradeStateDesc = *o.TradeStateDesc
	}

	if o.Amount != nil {
		ret.Amount = o.Amount.Clone()
	}

	if o.PromotionDetail != nil {
		ret.PromotionDetail = o.PromotionDetail.Clone()
	}

	return &ret
}

// QueryOrderByIdRequest
type QueryOrderByIdRequest struct {
	TransactionId *string `json:"transaction_id"`
	// 直连商户号
	Mchid *string `json:"mchid"`
}

func (o QueryOrderByIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.TransactionId == nil {
		return nil, fmt.Errorf("field `TransactionId` is required and must be specified in QueryOrderByIdRequest")
	}
	toSerialize["transaction_id"] = o.TransactionId

	if o.Mchid == nil {
		return nil, fmt.Errorf("field `Mchid` is required and must be specified in QueryOrderByIdRequest")
	}
	toSerialize["mchid"] = o.Mchid
	return json.Marshal(toSerialize)
}

func (o QueryOrderByIdRequest) String() string {
	var ret string
	if o.TransactionId == nil {
		ret += "TransactionId:<nil>, "
	} else {
		ret += fmt.Sprintf("TransactionId:%v, ", *o.TransactionId)
	}

	if o.Mchid == nil {
		ret += "Mchid:<nil>"
	} else {
		ret += fmt.Sprintf("Mchid:%v", *o.Mchid)
	}

	return fmt.Sprintf("QueryOrderByIdRequest{%s}", ret)
}

func (o QueryOrderByIdRequest) Clone() *QueryOrderByIdRequest {
	ret := QueryOrderByIdRequest{}

	if o.TransactionId != nil {
		ret.TransactionId = new(string)
		*ret.TransactionId = *o.TransactionId
	}

	if o.Mchid != nil {
		ret.Mchid = new(string)
		*ret.Mchid = *o.Mchid
	}

	return &ret
}

// QueryOrderByOutTradeNoRequest
type QueryOrderByOutTradeNoRequest struct {
	OutTradeNo *string `json:"out_trade_no"`
	// 直连商户号
	Mchid *string `json:"mchid"`
}

func (o QueryOrderByOutTradeNoRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.OutTradeNo == nil {
		return nil, fmt.Errorf("field `OutTradeNo` is required and must be specified in QueryOrderByOutTradeNoRequest")
	}
	toSerialize["out_trade_no"] = o.OutTradeNo

	if o.Mchid == nil {
		return nil, fmt.Errorf("field `Mchid` is required and must be specified in QueryOrderByOutTradeNoRequest")
	}
	toSerialize["mchid"] = o.Mchid
	return json.Marshal(toSerialize)
}

func (o QueryOrderByOutTradeNoRequest) String() string {
	var ret string
	if o.OutTradeNo == nil {
		ret += "OutTradeNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutTradeNo:%v, ", *o.OutTradeNo)
	}

	if o.Mchid == nil {
		ret += "Mchid:<nil>"
	} else {
		ret += fmt.Sprintf("Mchid:%v", *o.Mchid)
	}

	return fmt.Sprintf("QueryOrderByOutTradeNoRequest{%s}", ret)
}

func (o QueryOrderByOutTradeNoRequest) Clone() *QueryOrderByOutTradeNoRequest {
	ret := QueryOrderByOutTradeNoRequest{}

	if o.OutTradeNo != nil {
		ret.OutTradeNo = new(string)
		*ret.OutTradeNo = *o.OutTradeNo
	}

	if o.Mchid != nil {
		ret.Mchid = new(string)
		*ret.Mchid = *o.Mchid
	}

	return &ret
}

// SceneInfo 支付场景描述
type SceneInfo struct {
	// 用户终端IP
	PayerClientIp *string `json:"payer_client_ip"`
	// 商户端设备号
	DeviceId  *string    `json:"device_id,omitempty"`
	StoreInfo *StoreInfo `json:"store_info,omitempty"`
}

func (o SceneInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.PayerClientIp == nil {
		return nil, fmt.Errorf("field `PayerClientIp` is required and must be specified in SceneInfo")
	}
	toSerialize["payer_client_ip"] = o.PayerClientIp

	if o.DeviceId != nil {
		toSerialize["device_id"] = o.DeviceId
	}

	if o.StoreInfo != nil {
		toSerialize["store_info"] = o.StoreInfo
	}
	return json.Marshal(toSerialize)
}

func (o SceneInfo) String() string {
	var ret string
	if o.PayerClientIp == nil {
		ret += "PayerClientIp:<nil>, "
	} else {
		ret += fmt.Sprintf("PayerClientIp:%v, ", *o.PayerClientIp)
	}

	if o.DeviceId == nil {
		ret += "DeviceId:<nil>, "
	} else {
		ret += fmt.Sprintf("DeviceId:%v, ", *o.DeviceId)
	}

	ret += fmt.Sprintf("StoreInfo:%v", o.StoreInfo)

	return fmt.Sprintf("SceneInfo{%s}", ret)
}

func (o SceneInfo) Clone() *SceneInfo {
	ret := SceneInfo{}

	if o.PayerClientIp != nil {
		ret.PayerClientIp = new(string)
		*ret.PayerClientIp = *o.PayerClientIp
	}

	if o.DeviceId != nil {
		ret.DeviceId = new(string)
		*ret.DeviceId = *o.DeviceId
	}

	if o.StoreInfo != nil {
		ret.StoreInfo = o.StoreInfo.Clone()
	}

	return &ret
}

// SettleInfo
type SettleInfo struct {
	// 是否指定分账
	ProfitSharing *bool `json:"profit_sharing,omitempty"`
}

func (o SettleInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.ProfitSharing != nil {
		toSerialize["profit_sharing"] = o.ProfitSharing
	}
	return json.Marshal(toSerialize)
}

func (o SettleInfo) String() string {
	var ret string
	if o.ProfitSharing == nil {
		ret += "ProfitSharing:<nil>"
	} else {
		ret += fmt.Sprintf("ProfitSharing:%v", *o.ProfitSharing)
	}

	return fmt.Sprintf("SettleInfo{%s}", ret)
}

func (o SettleInfo) Clone() *SettleInfo {
	ret := SettleInfo{}

	if o.ProfitSharing != nil {
		ret.ProfitSharing = new(bool)
		*ret.ProfitSharing = *o.ProfitSharing
	}

	return &ret
}

// StoreInfo 商户门店信息
type StoreInfo struct {
	// 商户侧门店编号
	Id *string `json:"id"`
	// 商户侧门店名称
	Name *string `json:"name,omitempty"`
	// 地区编码，详细请见微信支付提供的文档
	AreaCode *string `json:"area_code,omitempty"`
	// 详细的商户门店地址
	Address *string `json:"address,omitempty"`
}

func (o StoreInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Id == nil {
		return nil, fmt.Errorf("field `Id` is required and must be specified in StoreInfo")
	}
	toSerialize["id"] = o.Id

	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	if o.AreaCode != nil {
		toSerialize["area_code"] = o.AreaCode
	}

	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

func (o StoreInfo) String() string {
	var ret string
	if o.Id == nil {
		ret += "Id:<nil>, "
	} else {
		ret += fmt.Sprintf("Id:%v, ", *o.Id)
	}

	if o.Name == nil {
		ret += "Name:<nil>, "
	} else {
		ret += fmt.Sprintf("Name:%v, ", *o.Name)
	}

	if o.AreaCode == nil {
		ret += "AreaCode:<nil>, "
	} else {
		ret += fmt.Sprintf("AreaCode:%v, ", *o.AreaCode)
	}

	if o.Address == nil {
		ret += "Address:<nil>"
	} else {
		ret += fmt.Sprintf("Address:%v", *o.Address)
	}

	return fmt.Sprintf("StoreInfo{%s}", ret)
}

func (o StoreInfo) Clone() *StoreInfo {
	ret := StoreInfo{}

	if o.Id != nil {
		ret.Id = new(string)
		*ret.Id = *o.Id
	}

	if o.Name != nil {
		ret.Name = new(string)
		*ret.Name = *o.Name
	}

	if o.AreaCode != nil {
		ret.AreaCode = new(string)
		*ret.AreaCode = *o.AreaCode
	}

	if o.Address != nil {
		ret.Address = new(string)
		*ret.Address = *o.Address
	}

	return &ret
}

// Payer
type Payer struct {
	// 用户在商户appid下的唯一标识。
	Openid *string `json:"openid,omitempty"`
}

func (o Payer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Openid != nil {
		toSerialize["openid"] = o.Openid
	}
	return json.Marshal(toSerialize)
}

func (o Payer) String() string {
	var ret string
	if o.Openid == nil {
		ret += "Openid:<nil>"
	} else {
		ret += fmt.Sprintf("Openid:%v", *o.Openid)
	}

	return fmt.Sprintf("Payer{%s}", ret)
}

func (o Payer) Clone() *Payer {
	ret := Payer{}

	if o.Openid != nil {
		ret.Openid = new(string)
		*ret.Openid = *o.Openid
	}

	return &ret
}
