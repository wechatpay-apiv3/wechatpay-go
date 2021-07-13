// Copyright 2021 Tencent Inc. All rights reserved.

package cipher

import "context"

// Cipher 使用证书对数据中进行原地加密/使用私钥对数据进行原地解密的功能
type Cipher interface {
	// Encrypt 使用证书对数据进行原地加密，密文会直接体现在入参 in 中，并返回加密所使用的证书序列号
	Encrypt(ctx context.Context, in interface{}) (string, error)

	// Decrypt 使用私钥对数据进行原地解密，明文会直接体现在入参 in 中，无返回
	Decrypt(ctx context.Context, in interface{}) error
}
