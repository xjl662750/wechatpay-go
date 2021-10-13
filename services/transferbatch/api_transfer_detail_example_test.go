// Copyright 2021 Tencent Inc. All rights reserved.
//
// 商家转账对外API
//
// - 场景及业务流程： 商户可通过该产品实现同时向多个用户微信零钱进行转账的操作，可用于发放奖金补贴、佣金货款结算、员工报销等场景。 [https://pay.weixin.qq.com/index.php/public/product/detail?pid=108&productType=0](https://pay.weixin.qq.com/index.php/public/product/detail?pid=108&productType=0) - 接入步骤：     - 商户在微信支付商户平台开通“批量转账到零钱”产品权限，并勾选“使用API方式发起转账”。     - 调用批量转账接口，对多个用户微信零钱发起转账。     - 调用查询批次接口，可获取到转账批次详情及当前状态。     - 调用查询明细接口，可获取到单条转账明细详情及当前状态。
//
// API version: 1.0.0

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package transferbatch_test

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/transferbatch"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

func ExampleTransferDetailApiService_GetTransferDetailByNo() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Print("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
	}

	svc := transferbatch.TransferDetailApiService{Client: client}
	resp, result, err := svc.GetTransferDetailByNo(ctx,
		transferbatch.GetTransferDetailByNoRequest{
			BatchId:  core.String("1030000071100999991182020050700019480001"),
			DetailId: core.String("1040000071100999991182020050700019500100"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call GetTransferDetailByNo err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleTransferDetailApiService_GetTransferDetailByOutNo() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Print("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
	}

	svc := transferbatch.TransferDetailApiService{Client: client}
	resp, result, err := svc.GetTransferDetailByOutNo(ctx,
		transferbatch.GetTransferDetailByOutNoRequest{
			OutDetailNo: core.String("x23zy545Bd5436"),
			OutBatchNo:  core.String("plfk2020042013"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call GetTransferDetailByOutNo err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}
