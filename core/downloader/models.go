// Copyright 2021 Tencent Inc. All rights reserved.

package downloader

import (
	"encoding/json"
	"fmt"
	"time"
)

// rawCertificate 微信支付平台证书信息
type rawCertificate struct {
	// 证书序列号
	SerialNo *string `json:"serial_no"`
	// 证书有效期开始时间
	EffectiveTime *time.Time `json:"effective_time"`
	// 证书过期时间
	ExpireTime *time.Time `json:"expire_time"`
	// 为了保证安全性，微信支付在回调通知和平台证书下载接口中，对关键信息进行了AES-256-GCM加密
	EncryptCertificate *encryptCertificate `json:"encrypt_certificate"`
}

// MarshalJSON 自定义JSON序列化
func (o rawCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.SerialNo == nil {
		return nil, fmt.Errorf("field `SerialNo` is required and must be specified in rawCertificate")
	}
	toSerialize["serial_no"] = o.SerialNo

	if o.EffectiveTime == nil {
		return nil, fmt.Errorf("field `EffectiveTime` is required and must be specified in rawCertificate")
	}
	toSerialize["effective_time"] = o.EffectiveTime.Format(time.RFC3339)

	if o.ExpireTime == nil {
		return nil, fmt.Errorf("field `ExpireTime` is required and must be specified in rawCertificate")
	}
	toSerialize["expire_time"] = o.ExpireTime.Format(time.RFC3339)

	if o.EncryptCertificate == nil {
		return nil, fmt.Errorf("field `encryptCertificate` is required and must be specified in rawCertificate")
	}
	toSerialize["encrypt_certificate"] = o.EncryptCertificate
	return json.Marshal(toSerialize)
}

// String 自定义字符串表达
func (o rawCertificate) String() string {
	var ret string
	if o.SerialNo == nil {
		ret += "SerialNo:<nil>, "
	} else {
		ret += fmt.Sprintf("SerialNo:%v, ", *o.SerialNo)
	}

	if o.EffectiveTime == nil {
		ret += "EffectiveTime:<nil>, "
	} else {
		ret += fmt.Sprintf("EffectiveTime:%v, ", *o.EffectiveTime)
	}

	if o.ExpireTime == nil {
		ret += "ExpireTime:<nil>, "
	} else {
		ret += fmt.Sprintf("ExpireTime:%v, ", *o.ExpireTime)
	}

	ret += fmt.Sprintf("encryptCertificate:%v", o.EncryptCertificate)

	return fmt.Sprintf("rawCertificate{%s}", ret)
}

type downloadCertificatesResponse struct {
	// 平台证书列表
	Data []rawCertificate `json:"data,omitempty"`
}

// MarshalJSON 自定义JSON序列化
func (o downloadCertificatesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

// String 自定义字符串表达
func (o downloadCertificatesResponse) String() string {
	var ret string
	ret += fmt.Sprintf("Data:%v", o.Data)

	return fmt.Sprintf("downloadCertificatesResponse{%s}", ret)
}

// encryptCertificate 为了保证安全性，微信支付在回调通知和平台证书下载接口中，对关键信息进行了AES-256-GCM加密
type encryptCertificate struct {
	// 加密所使用的算法，目前可能取值仅为 AEAD_AES_256_GCM
	Algorithm *string `json:"algorithm"`
	// 加密所使用的随机字符串
	Nonce *string `json:"nonce"`
	// 附加数据包（可能为空）
	AssociatedData *string `json:"associated_data"`
	// 证书内容密文，解密后会获得证书完整内容
	Ciphertext *string `json:"ciphertext"`
}

// MarshalJSON 自定义JSON序列化
func (o encryptCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Algorithm == nil {
		return nil, fmt.Errorf("field `Algorithm` is required and must be specified in encryptCertificate")
	}
	toSerialize["algorithm"] = o.Algorithm

	if o.Nonce == nil {
		return nil, fmt.Errorf("field `Nonce` is required and must be specified in encryptCertificate")
	}
	toSerialize["nonce"] = o.Nonce

	if o.AssociatedData == nil {
		return nil, fmt.Errorf("field `AssociatedData` is required and must be specified in encryptCertificate")
	}
	toSerialize["associated_data"] = o.AssociatedData

	if o.Ciphertext == nil {
		return nil, fmt.Errorf("field `Ciphertext` is required and must be specified in encryptCertificate")
	}
	toSerialize["ciphertext"] = o.Ciphertext
	return json.Marshal(toSerialize)
}

// String 自定义字符串表达
func (o encryptCertificate) String() string {
	var ret string
	if o.Algorithm == nil {
		ret += "Algorithm:<nil>, "
	} else {
		ret += fmt.Sprintf("Algorithm:%v, ", *o.Algorithm)
	}

	if o.Nonce == nil {
		ret += "Nonce:<nil>, "
	} else {
		ret += fmt.Sprintf("Nonce:%v, ", *o.Nonce)
	}

	if o.AssociatedData == nil {
		ret += "AssociatedData:<nil>, "
	} else {
		ret += fmt.Sprintf("AssociatedData:%v, ", *o.AssociatedData)
	}

	if o.Ciphertext == nil {
		ret += "Ciphertext:<nil>"
	} else {
		ret += fmt.Sprintf("Ciphertext:%v", *o.Ciphertext)
	}

	return fmt.Sprintf("encryptCertificate{%s}", ret)
}
