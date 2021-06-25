package main

import (
	"context"
	"crypto/x509"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

var (
	mchID             string
	mchSerialNo       string
	mchPrivateKeyPath string
	mchAPIv3Key       string

	wechatPayCertificatePath string
	outputPath               string
)

func init() {
	flag.StringVar(&mchID, "m", "", "【必传】`商户号`")
	flag.StringVar(&mchSerialNo, "s", "", "【必传】`商户证书序列号`")
	flag.StringVar(&mchPrivateKeyPath, "p", "", "【必传】`商户私钥路径`")
	flag.StringVar(&mchAPIv3Key, "k", "", "【必传】`商户APIv3密钥`")

	flag.StringVar(&wechatPayCertificatePath, "c", "", "【可选】`商户平台证书路径`，用于验签。省略则跳过验签")
	flag.StringVar(&outputPath, "o", "./", "【可选】`证书下载保存目录`")
}

func main() {
	flag.Parse()
	flag.Usage = usage

	if err := checkArgs(); err != nil {
		reportError("参数有误：%v", err)
		usage()
		os.Exit(2)
	}

	ctx := context.Background()
	client, err := createClient(ctx)
	if err != nil {
		reportError("%v", err)
		os.Exit(2)
	}

	d, err := downloader.NewCertificateDownloaderWithClient(ctx, client, mchAPIv3Key)
	if err != nil {
		reportError("下载证书失败：%v", err)
		os.Exit(2)
	}

	err = saveCertificates(ctx, d)
	if err != nil {
		reportError("%v", err)
	}

	os.Exit(0)
}

func reportError(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format+"\n", a...)
}

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, "usage of wechatpay_download_certs:\n")
	flag.PrintDefaults()
	os.Exit(2)
}

type paramError struct {
	name    string
	value   string
	message string
}

func (e paramError) Error() string {
	if e.value != "" {
		return fmt.Sprintf("%v(%v) %v", e.name, e.value, e.message)
	}
	return fmt.Sprintf("%v %v", e.name, e.message)
}

func checkArgs() error {
	if mchID == "" {
		return paramError{"商户号", mchID, "必传"}
	}

	if mchSerialNo == "" {
		return paramError{"商户证书序列号", mchSerialNo, "必传"}
	}

	if mchPrivateKeyPath == "" {
		return paramError{"商户平台证书路径", mchPrivateKeyPath, "必传"}
	}

	fileInfo, err := os.Stat(mchPrivateKeyPath)
	if err != nil {
		return paramError{"商户私钥路径", mchPrivateKeyPath, fmt.Sprintf("有误: %v", err)}
	}
	if fileInfo.IsDir() {
		return paramError{"商户私钥路径", mchPrivateKeyPath, "不是合法的文件路径"}
	}

	if mchAPIv3Key == "" {
		return paramError{"商户APIv3密钥", mchAPIv3Key, "必传"}
	}

	if wechatPayCertificatePath != "" {
		fileInfo, err := os.Stat(wechatPayCertificatePath)
		if err != nil {
			return paramError{"商户平台证书路径", wechatPayCertificatePath, fmt.Sprintf("有误：%v", err)}
		}
		if fileInfo.IsDir() {
			return paramError{"商户平台证书路径", wechatPayCertificatePath, "不是合法的文件路径"}
		}
	}

	err = os.MkdirAll(outputPath, os.ModePerm)
	if err != nil {
		return paramError{"证书下载保存目录", outputPath, fmt.Sprintf("创建失败：%v", err)}
	}

	return nil
}

func saveCertificates(ctx context.Context, d *downloader.CertificateDownloader) error {
	for serialNo, certContent := range d.ExportAll(ctx) {
		outputFilePath := filepath.Join(outputPath, fmt.Sprintf("wechatpay_%v.pem", serialNo))

		f, err := os.Create(outputFilePath)
		if err != nil {
			return fmt.Errorf("创建证书文件`%v`失败：%v", outputFilePath, err)
		}

		_, err = f.WriteString(certContent + "\n")
		if err != nil {
			return fmt.Errorf("写入证书到`%v`失败: %v", outputFilePath, err)
		}

		fmt.Printf("写入证书到`%v`成功\n", outputFilePath)
	}
	return nil
}

func createClient(ctx context.Context) (*core.Client, error) {
	privateKey, err := utils.LoadPrivateKeyWithPath(mchPrivateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("商户私钥有误：%v", err)
	}

	var client *core.Client
	if wechatPayCertificatePath != "" {
		wechatPayCertificate, err := utils.LoadCertificateWithPath(wechatPayCertificatePath)
		if err != nil {
			return nil, fmt.Errorf("平台证书有误：%v", err)
		}
		client, err = core.NewClient(
			ctx, option.WithMerchantCredential(mchID, mchSerialNo, privateKey),
			option.WithWechatPayCertificate([]*x509.Certificate{wechatPayCertificate}),
		)
		if err != nil {
			return nil, fmt.Errorf("创建 Client 失败：%v", err)
		}
	} else {
		client, err = core.NewClient(
			ctx, option.WithMerchantCredential(mchID, mchSerialNo, privateKey), option.WithoutValidator(),
		)
		if err != nil {
			return nil, fmt.Errorf("创建 Client 失败：%v", err)
		}
	}

	return client, nil
}
