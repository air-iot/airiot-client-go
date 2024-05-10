package api_client_go

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strings"
)

type KratosRegistryClient struct {
	registry *etcd.Registry
}

func NewKartosRegistryClient(cli *clientv3.Client, options ...etcd.Option) *KratosRegistryClient {
	return &KratosRegistryClient{registry: etcd.New(cli, options...)}
}

// GetServiceInstances 从注册中心查询指定服务的实例列表. 可以通过 env 过滤服务实例
//
// serviceName: 服务名称
//
// allowEmptyEnv: 是否允许环境为空的实例. 如果为 true, 则返回的实例列表中包含未定义 env 元数据或者 env 为空的实例
//
// envs: 环境列表. 如果为空, 则返回所有实例. 如果不为空, 则返回指定环境的实例
func (reg *KratosRegistryClient) GetServiceInstances(ctx context.Context, serviceName string, allowEmptyEnv bool, envs ...string) ([]*registry.ServiceInstance, error) {
	serviceInstances, err := reg.registry.GetService(ctx, serviceName)
	if err != nil {
		return nil, err
	}

	if len(serviceInstances) == 0 {
		return nil, nil
	} else if len(envs) == 0 {
		return serviceInstances, nil
	}

	finalServiceInstances := make([]*registry.ServiceInstance, 0, len(serviceInstances))
	for i := range serviceInstances {
		inst := serviceInstances[i]
		if env, ok := inst.Metadata["env"]; !ok || env == "" {
			if allowEmptyEnv {
				finalServiceInstances = append(finalServiceInstances, inst)
			}
		} else {
			for ei := range envs {
				if strings.Compare(envs[ei], env) == 0 {
					finalServiceInstances = append(finalServiceInstances, inst)
					break
				}
			}
		}
	}

	return finalServiceInstances, nil
}

// GetServiceEndpoints 获取服务实例的指定协议的端点. 如果 schemes 为空, 则返回所有端点
//
// serviceInstances: 服务实例列表
//
// schemes: 协议列表, 例如: http, https, grpc 等. 如果为空, 则返回所有端点. 如果不为空, 则返回指定协议的端点
func (reg *KratosRegistryClient) GetServiceEndpoints(serviceInstances []*registry.ServiceInstance, schemes ...string) []string {
	var endpoints []string
	for i := range serviceInstances {
		inst := serviceInstances[i]
		for j := range inst.Endpoints {
			ep := inst.Endpoints[j]
			if len(schemes) == 0 {
				endpoints = append(endpoints, ep)
			} else {
				for k := range schemes {
					if strings.HasPrefix(ep, schemes[k]) {
						endpoints = append(endpoints, ep)
						break
					}
				}
			}
		}
	}
	return endpoints
}

// GetServiceEndpointsByServiceName 从注册中心查询指定服务的实例列表, 并获取指定协议的端点
//
// serviceName: 服务名称
//
// allowEmptyEnv: 是否允许环境为空的实例. 如果为 true, 则返回的实例列表中包含未定义 env 元数据或者 env 为空的实例
//
// envs: 环境列表. 如果为空, 则返回所有实例. 如果不为空, 则返回指定环境的实例
//
// schemes: 协议列表, 例如: http, https, grpc 等. 如果为空, 则返回所有端点. 如果不为空, 则返回指定协议的端点
func (reg *KratosRegistryClient) GetServiceEndpointsByServiceName(ctx context.Context, serviceName string, allowEmptyEnv bool, envs []string, schemes []string) ([]string, error) {
	serviceInstances, err := reg.GetServiceInstances(ctx, serviceName, allowEmptyEnv, envs...)
	if err != nil {
		return nil, err
	} else if len(serviceInstances) == 0 {
		return nil, nil
	}
	return reg.GetServiceEndpoints(serviceInstances, schemes...), nil
}
