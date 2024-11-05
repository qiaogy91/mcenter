package endpoint

import (
	"context"
	"github.com/qiaogy91/ioc"
)

const AppName = "endpoint"

func GetSvc() Service { return ioc.Controller().Get(AppName).(Service) }

type Service interface {
	RpcServer
	CreateTable(context.Context) error
}

func (s *EndpointSpecSet) Services() []string {
	var svc []string
	var dic = map[string]int{}

	for _, item := range s.Items {
		dic[item.Service] += 1
	}

	for k := range dic {
		svc = append(svc, k)
	}
	return svc
}

func (s *EndpointSpecSet) EndpointSet() *EndpointSet {
	ep := &EndpointSet{}
	for _, item := range s.Items {
		ep.Items = append(ep.Items, &Endpoint{Spec: item})
		ep.Total += 1
	}
	return ep
}
