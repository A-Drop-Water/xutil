package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

type Registry struct {
	client *api.Client
}

// NewConsulRegistry 创建并返回一个新的 ConsulRegistry 实例
func NewConsulRegistry(ip string, port int) (*Registry, error) {
	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%d", ip, port) // 设置 Consul 的地址和端口
	client, err := api.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("创建 Consul 客户端失败: %w", err)
	}
	return &Registry{client: client}, nil
}

// RegisterHttpService HTTP服务注册
// 传的就是: 服务名、服务ID、服务IP+Port、健康检查信息
// 注册一个服务多个实例就是同一个name不同的id
func (m *Registry) RegisterHttpService(name, id, host string, port int, tags []string) error {
	check := &api.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d/health", host, port), // 健康检查的HTTP地址
		Interval:                       "10s",                                          // 健康检查间隔
		Timeout:                        "5s",                                           // 健康检查超时时间
		DeregisterCriticalServiceAfter: "1m",                                           // 如果服务不可用，多久后注销
	}
	cfg := &api.AgentServiceRegistration{
		ID:      id,
		Name:    name,
		Tags:    tags,
		Port:    port,
		Address: host,
		Check:   check,
	}
	// 直接调用client注册
	return m.client.Agent().ServiceRegister(cfg)
}

// DeRegisterService 服务注销
// 传的就是: 服务ID
func (m *Registry) DeRegisterService(id string) error {
	// 直接调用client注册
	return m.client.Agent().ServiceDeregister(id)
}
