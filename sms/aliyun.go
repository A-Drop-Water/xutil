package sms

import (
	"context"
	"encoding/json"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v5/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/credentials-go/credentials"
)

type AliyunService struct {
	client   *dysmsapi20170525.Client
	appId    string
	signName string
}

func NewAliyunService(cfg *Config) Service {
	cli, err := createClientAli(cfg)
	if err != nil {
		panic(err)
	}
	s := &AliyunService{
		client:   cli,
		appId:    cfg.AppId,
		signName: cfg.SignName,
	}
	return s
}

func createClientAli(cfg *Config) (*dysmsapi20170525.Client, error) {
	// 工程代码建议使用更安全的无AK方式，凭据配置方式请参见：https://help.aliyun.com/document_detail/378661.html。
	var (
		credential credentials.Credential
		err        error
	)

	// 初始化id和secret信息
	if cfg == nil || (cfg.AccessKeyId == "" && cfg.AccessKeySecret == "") {
		credential, err = credentials.NewCredential(nil)
	} else {
		credential, err = credentials.NewCredential(&credentials.Config{
			AccessKeyId:     &cfg.AccessKeyId,
			AccessKeySecret: &cfg.AccessKeySecret,
		})
	}
	if err != nil {
		return nil, err
	}

	config := &openapi.Config{
		Credential: credential,
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Dysmsapi
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	cli := &dysmsapi20170525.Client{}
	cli, err = dysmsapi20170525.NewClient(config)
	return cli, err
}

func buildSliceToString(p []string) string {
	// 将这一系列手机号通过,分割
	res := ""
	for k, v := range p {
		res += v
		if k != len(p)-1 {
			res += ","
		}
	}
	return res
}

func (s *AliyunService) SendSMS(context context.Context, templateId string, phones []string, params map[string]string) error {
	// 1. 初始化request
	paramStr, _ := json.Marshal(params)
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(buildSliceToString(phones)), // 格式是什么?
		SignName:      tea.String(s.signName),
		TemplateCode:  tea.String(templateId),
		TemplateParam: tea.String(string(paramStr)),
	}
	// 2. 发送这个请求

	// TODO 后续考虑如何处理错误信息
	_, err := s.client.SendSms(sendSmsRequest)
	if err != nil {
		return err
	}
	return nil
}
