package encryptors

import (
	"context"
	"crypto/x509"
	"fmt"
	"sync"
	"time"

	"git.code.oa.com/emmetzhang/codegen-go-sdk/utils"
)

type WechatPayEncryptor struct {
	// 微信支付平台证书 证书序列号 -> 证书
	certificates map[string]*x509.Certificate
	// 当前可用且有效期最远的证书序列号
	bestCertSerial string
	// 读写锁，用于保证 certificates 与 bestCertSerial 的访问并发安全性
	lock sync.RWMutex
}

// NewWechatPayEncryptor 新建一个 WechatPayEncryptor
//
// 如果需要在并发环境（GoRoutine）中共享使用它，请使用 NewConcurrentSafeWechatPayEncryptor 生成并发安全版本
func NewWechatPayEncryptor(certificates []*x509.Certificate) *WechatPayEncryptor {
	e := WechatPayEncryptor{}
	e.SetCertificates(certificates)
	return &e
}

// SelectCertificate 选择合适的微信支付平台证书用于加密
func (e *WechatPayEncryptor) SelectCertificate(ctx context.Context) (serial string, err error) {
	e.rLock()
	bestCert, ok := e.certificates[e.bestCertSerial]
	e.rUnlock()

	if ok && utils.IsCertValid(*bestCert, time.Now()) {
		return e.bestCertSerial, nil
	}

	// 当前 bestCertSerial 不再合法，需要更新
	e.wLock()
	defer e.wUnlock()

	e.updateBestCertSerial()

	if e.bestCertSerial == "" {
		return "", fmt.Errorf("no valid certificate found")
	}

	// 新更新得到的最佳选择
	return e.bestCertSerial, nil
}

// Encrypt 对字符串加密
func (e *WechatPayEncryptor) Encrypt(ctx context.Context, serial, plaintext string) (ciphertext string, err error) {
	e.rLock()
	cert, ok := e.certificates[serial]
	e.rUnlock()

	if !ok {
		return plaintext, fmt.Errorf("cert for EncryptSerial not found")
	}

	return utils.EncryptOAEPWithCertificate(plaintext, cert)
}

// SetCertificates 重置微信支付平台证书列表
func (e *WechatPayEncryptor) SetCertificates(certificates []*x509.Certificate) {
	c := make(map[string]*x509.Certificate)

	now := time.Now()
	for _, cert := range certificates {
		if !utils.IsCertExpired(*cert, now) {
			c[utils.GetCertificateSerialNumber(*cert)] = cert
		}
	}

	e.wLock()
	defer e.wUnlock()

	e.certificates = c
	e.updateBestCertSerial()
}

// AddCertificate 增加微信支付平台证书，如果证书已经过期会自动忽略
func (e *WechatPayEncryptor) AddCertificate(cert *x509.Certificate) {
	if utils.IsCertExpired(*cert, time.Now()) {
		return
	}

	e.wLock()
	defer e.wUnlock()

	e.certificates[utils.GetCertificateSerialNumber(*cert)] = cert
	e.updateBestCertSerial()
}

// RemoveCertificate 删除微信支付平台证书，证书不存在时自动忽略
func (e *WechatPayEncryptor) RemoveCertificate(serial string) {
	e.wLock()
	defer e.wUnlock()

	delete(e.certificates, serial)

	if serial == e.bestCertSerial {
		e.updateBestCertSerial()
	}
}

// rLock 读锁
func (e *WechatPayEncryptor) rLock() {
	e.lock.RLock()
}

// rUnlock 读解锁
func (e *WechatPayEncryptor) rUnlock() {
	e.lock.RUnlock()
}

// wLock 写锁
func (e *WechatPayEncryptor) wLock() {
	e.lock.Lock()
}

// wUnlock 写解锁
func (e *WechatPayEncryptor) wUnlock() {
	e.lock.Unlock()
}

// updateBestCertSerial 重新选择最佳微信支付平台证书
// 根据微信支付平台证书文档说明，应优先使用最新的证书（即启用时间最晚）
// https://wechatpay-api.gitbook.io/wechatpay-api-v3/jie-kou-wen-dang/ping-tai-zheng-shu#zhu-yi-shi-xiang
// 注意：由于 RWMutex 锁不可重入，updateBestCertSerial 方法不会尝试获取锁，应由调用方获取写锁
func (e *WechatPayEncryptor) updateBestCertSerial() {
	e.bestCertSerial = ""

	now := time.Now()
	for serial, cert := range e.certificates {
		if !utils.IsCertValid(*cert, now) {
			continue
		}

		if e.bestCertSerial == "" {
			e.bestCertSerial = serial
		} else {
			bestCert := e.certificates[e.bestCertSerial]
			if cert.NotBefore.After(bestCert.NotBefore) {
				e.bestCertSerial = serial
			}
		}
	}
}
