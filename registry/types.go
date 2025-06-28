package registry

import (
	"github.com/A-Drop-Water/xutil/registry/consul"
)

type Registry interface {
	RegisterHttpService(name, id, host string, port int, tags []string) error
	DeRegisterService(id string) error
}

func NewConsulRegistry(ip string, port int) Registry {
	rt, err := consul.NewConsulRegistry(ip, port)
	if err != nil {
		panic(err)
	}
	return rt
}
