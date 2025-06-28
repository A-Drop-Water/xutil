package sms

import (
	"context"
	"fmt"
)

// MemoryService 本地实现
// 其实就是没有实现
type MemoryService struct {
}

func NewMemoryService() Service {
	return &MemoryService{}
}

func (s *MemoryService) SendSMS(context context.Context, templateId string, phones []string, params map[string]string) error {
	fmt.Println(fmt.Sprintf("成功发送短信,使用的templateid为: %s,发送的手机号为: %v,发送的模板参数为 %v", templateId, phones, params))
	return nil
}
