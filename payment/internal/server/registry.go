package server

import (
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/hashicorp/consul/api"
	"log"
	"payment/internal/conf"
)

// NewRegistrar 使用Consul作为注册中心
func NewRegistrar(c *conf.Registry) registry.Registrar {
	// 创建consul客户端
	configs := api.DefaultConfig()
	// 从conf/conf.proto定义Consul配置与configs/config.yml配置文件中读取consul的配置
	configs.Address = c.Consul.Address
	configs.Scheme = c.Consul.Scheme
	// 创建consul客户端
	consulClient, err := api.NewClient(configs)
	if err != nil {
		log.Fatal(err)
	}

	// 创建consul注册中心
	r := consul.New(consulClient, consul.WithHealthCheck(c.Consul.HealthCheck))
	return r
}
