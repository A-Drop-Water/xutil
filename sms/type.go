package sms

import (
	"context"
)

type Service interface {
	SendSMS(context context.Context, templateId string, phones []string, params map[string]string) error
}

// Config SMS配置
type Config struct {
	AccessKeyId     string
	AccessKeySecret string
	AppId           string // 应用ID
	SignName        string // 应用统一签名
}
