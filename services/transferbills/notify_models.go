// Copyright 2021 Tencent Inc. All rights reserved.
//
// 商家转账到用户零钱API - 通知相关模型
//
// API version: 1.0.0

package transferbills

import (
	"encoding/json"
	"fmt"
	"time"
)

// 转账通知事件类型
const TransferNotifyEventType = "MCHTRANSFER.BILL.FINISHED"

// 转账通知状态常量
const (
	TransferNotifyStateSuccess   = "SUCCESS"    // 转账成功
	TransferNotifyStateFail      = "FAIL"       // 转账失败
	TransferNotifyStateCancelled = "CANCELLED"  // 转账已撤销
)

// TransferNotifyContent 转账通知内容
type TransferNotifyContent struct {
	// 商户号
	Mchid *string `json:"mchid"`
	// 应用ID
	Appid *string `json:"appid"`
	// 商户订单号
	OutBillNo *string `json:"out_bill_no"`
	// 微信转账单号
	TransferBillNo *string `json:"transfer_bill_no"`
	// 转账场景ID
	TransferSceneId *string `json:"transfer_scene_id"`
	// 用户openid
	Openid *string `json:"openid"`
	// 转账金额，单位为分
	TransferAmount *int64 `json:"transfer_amount"`
	// 转账备注
	TransferRemark *string `json:"transfer_remark"`
	// 收款方真实姓名，加密字段
	UserName *string `json:"user_name,omitempty"`
	// 转账状态：SUCCESS-转账成功, FAIL-转账失败, CANCELLED-转账已撤销
	State *string `json:"state"`
	// 失败原因，转账失败时存在
	FailReason *string `json:"fail_reason,omitempty"`
	// 转账创建时间
	CreateTime *time.Time `json:"create_time"`
	// 转账更新时间
	UpdateTime *time.Time `json:"update_time"`
}

func (o TransferNotifyContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Mchid != nil {
		toSerialize["mchid"] = o.Mchid
	}

	if o.Appid != nil {
		toSerialize["appid"] = o.Appid
	}

	if o.OutBillNo != nil {
		toSerialize["out_bill_no"] = o.OutBillNo
	}

	if o.TransferBillNo != nil {
		toSerialize["transfer_bill_no"] = o.TransferBillNo
	}

	if o.TransferSceneId != nil {
		toSerialize["transfer_scene_id"] = o.TransferSceneId
	}

	if o.Openid != nil {
		toSerialize["openid"] = o.Openid
	}

	if o.TransferAmount != nil {
		toSerialize["transfer_amount"] = o.TransferAmount
	}

	if o.TransferRemark != nil {
		toSerialize["transfer_remark"] = o.TransferRemark
	}

	if o.UserName != nil {
		toSerialize["user_name"] = o.UserName
	}

	if o.State != nil {
		toSerialize["state"] = o.State
	}

	if o.FailReason != nil {
		toSerialize["fail_reason"] = o.FailReason
	}

	if o.CreateTime != nil {
		toSerialize["create_time"] = o.CreateTime.Format(time.RFC3339)
	}

	if o.UpdateTime != nil {
		toSerialize["update_time"] = o.UpdateTime.Format(time.RFC3339)
	}

	return json.Marshal(toSerialize)
}

func (o TransferNotifyContent) String() string {
	var ret string

	if o.Mchid == nil {
		ret += "Mchid:<nil>, "
	} else {
		ret += fmt.Sprintf("Mchid:%v, ", *o.Mchid)
	}

	if o.Appid == nil {
		ret += "Appid:<nil>, "
	} else {
		ret += fmt.Sprintf("Appid:%v, ", *o.Appid)
	}

	if o.OutBillNo == nil {
		ret += "OutBillNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutBillNo:%v, ", *o.OutBillNo)
	}

	if o.TransferBillNo == nil {
		ret += "TransferBillNo:<nil>, "
	} else {
		ret += fmt.Sprintf("TransferBillNo:%v, ", *o.TransferBillNo)
	}

	if o.TransferSceneId == nil {
		ret += "TransferSceneId:<nil>, "
	} else {
		ret += fmt.Sprintf("TransferSceneId:%v, ", *o.TransferSceneId)
	}

	if o.Openid == nil {
		ret += "Openid:<nil>, "
	} else {
		ret += fmt.Sprintf("Openid:%v, ", *o.Openid)
	}

	if o.TransferAmount == nil {
		ret += "TransferAmount:<nil>, "
	} else {
		ret += fmt.Sprintf("TransferAmount:%v, ", *o.TransferAmount)
	}

	if o.TransferRemark == nil {
		ret += "TransferRemark:<nil>, "
	} else {
		ret += fmt.Sprintf("TransferRemark:%v, ", *o.TransferRemark)
	}

	if o.UserName == nil {
		ret += "UserName:<nil>, "
	} else {
		ret += fmt.Sprintf("UserName:%v, ", *o.UserName)
	}

	if o.State == nil {
		ret += "State:<nil>, "
	} else {
		ret += fmt.Sprintf("State:%v, ", *o.State)
	}

	if o.FailReason == nil {
		ret += "FailReason:<nil>, "
	} else {
		ret += fmt.Sprintf("FailReason:%v, ", *o.FailReason)
	}

	if o.CreateTime == nil {
		ret += "CreateTime:<nil>, "
	} else {
		ret += fmt.Sprintf("CreateTime:%v, ", *o.CreateTime)
	}

	if o.UpdateTime == nil {
		ret += "UpdateTime:<nil>"
	} else {
		ret += fmt.Sprintf("UpdateTime:%v", *o.UpdateTime)
	}

	return fmt.Sprintf("TransferNotifyContent{%s}", ret)
}

func (o TransferNotifyContent) Clone() *TransferNotifyContent {
	ret := TransferNotifyContent{}

	if o.Mchid != nil {
		ret.Mchid = new(string)
		*ret.Mchid = *o.Mchid
	}

	if o.Appid != nil {
		ret.Appid = new(string)
		*ret.Appid = *o.Appid
	}

	if o.OutBillNo != nil {
		ret.OutBillNo = new(string)
		*ret.OutBillNo = *o.OutBillNo
	}

	if o.TransferBillNo != nil {
		ret.TransferBillNo = new(string)
		*ret.TransferBillNo = *o.TransferBillNo
	}

	if o.TransferSceneId != nil {
		ret.TransferSceneId = new(string)
		*ret.TransferSceneId = *o.TransferSceneId
	}

	if o.Openid != nil {
		ret.Openid = new(string)
		*ret.Openid = *o.Openid
	}

	if o.TransferAmount != nil {
		ret.TransferAmount = new(int64)
		*ret.TransferAmount = *o.TransferAmount
	}

	if o.TransferRemark != nil {
		ret.TransferRemark = new(string)
		*ret.TransferRemark = *o.TransferRemark
	}

	if o.UserName != nil {
		ret.UserName = new(string)
		*ret.UserName = *o.UserName
	}

	if o.State != nil {
		ret.State = new(string)
		*ret.State = *o.State
	}

	if o.FailReason != nil {
		ret.FailReason = new(string)
		*ret.FailReason = *o.FailReason
	}

	if o.CreateTime != nil {
		ret.CreateTime = new(time.Time)
		*ret.CreateTime = *o.CreateTime
	}

	if o.UpdateTime != nil {
		ret.UpdateTime = new(time.Time)
		*ret.UpdateTime = *o.UpdateTime
	}

	return &ret
}