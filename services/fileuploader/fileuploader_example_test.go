package fileuploader_test

import (
	"context"
	"os"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/services/fileuploader"
)

func ExampleImageUploader_Upload() {
	var (
		ctx    context.Context
		client *core.Client
	)
	// 假设已获得初始化后的 core.Client

	file, err := os.Open("picture.jpg")
	if err != nil {
		return
	}
	defer file.Close()

	svc := fileuploader.ImageUploader{Client: client}
	resp, result, err := svc.Upload(ctx, file, "picture.jpg", consts.ImageJPG)

	// TODO: 处理返回结果
	_, _, _ = resp, result, err
}

func ExampleVideoUploader_Upload() {
	var (
		ctx    context.Context
		client *core.Client
	)
	// 假设已获得初始化后的 core.Client

	file, err := os.Open("video.mp4")
	if err != nil {
		return
	}
	defer file.Close()

	svc := fileuploader.VideoUploader{Client: client}
	resp, result, err := svc.Upload(ctx, file, "video.mp4", consts.VideoMP4)

	// TODO: 处理返回结果
	_, _, _ = resp, result, err
}

func ExampleMarketingImageUploader_Upload() {
	var (
		ctx    context.Context
		client *core.Client
	)
	// 假设已获得初始化后的 core.Client

	file, err := os.Open("picture.jpg")
	if err != nil {
		return
	}
	defer file.Close()

	svc := fileuploader.MarketingImageUploader{Client: client}
	resp, result, err := svc.Upload(ctx, file, "picture.jpg", consts.ImageJPG)

	// TODO: 处理返回结果
	_, _, _ = resp, result, err
}

func ExampleMchBizUploader_Upload() {
	var (
		ctx    context.Context
		client *core.Client
	)
	// 假设已获得初始化后的 core.Client

	file, err := os.Open("picture.jpg")
	if err != nil {
		return
	}
	defer file.Close()

	svc := fileuploader.MchBizUploader{Client: client}
	resp, result, err := svc.Upload(ctx, file, "picture.jpg", consts.ImageJPG)

	// TODO: 处理返回结果
	_, _, _ = resp, result, err
}
