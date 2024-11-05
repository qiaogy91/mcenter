package impl_test

import (
	"context"
	"fmt"
	"github.com/qiaogy91/ioc"
	_ "github.com/qiaogy91/ioc/config/otlp"
	"github.com/qiaogy91/mcenter/apps/endpoint"
	"testing"
)

var (
	ctx = context.Background()
	c   endpoint.Service
)

func init() {
	cf := "/Users/qiaogy/GolandProjects/projects/github/CloudManager/mcenter/etc/application.yaml"
	if err := ioc.ConfigIocObject(cf); err != nil {
		panic(err)
	}
	c = endpoint.GetSvc()
}

func TestImpl_CreateTable(t *testing.T) {
	if err := c.CreateTable(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestImpl_CreateEndpoint(t *testing.T) {

	var specs []*endpoint.Spec

	actions := []string{"get", "list", "delete"}
	// 10个服务
	for i := 1; i <= 10; i++ {
		// 每个服务10个模块
		for j := 1; j <= 10; j++ {
			// 每个模块3个方法
			for _, action := range actions {
				spec := &endpoint.Spec{
					Service:  fmt.Sprintf("svc_%d", i),
					Resource: fmt.Sprintf("md_%d", j),
					Action:   action,
					Identity: fmt.Sprintf("svc_%d.md_%d.%s./md01/api/v1/secret", i, j, action),
					Desc:     "测试",
				}
				specs = append(specs, spec)
			}

		}

	}

	req := &endpoint.EndpointSpecSet{Items: specs}
	ins, err := c.CreateEndpoint(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	for _, item := range ins.Items {
		t.Logf("%+v", item)
	}
}

func TestImpl_DeleteEndpoint(t *testing.T) {
	req := &endpoint.DeleteEndpointRequest{Service: "svc01"}
	ins, err := c.DeleteEndpoint(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range ins.Items {
		t.Logf("%+v", item)
	}
}

func TestImpl_QueryEndpoint(t *testing.T) {
	req := &endpoint.QueryEndpointRequest{
		PageNumber: 1,
		PageSize:   10,
		QueryType:  endpoint.QueryType_QUERY_TYPE_SERVICE,
		Keyword:    "svc_3",
	}

	ins, err := c.QueryEndpoint(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range ins.Items {
		t.Logf("%+v", item)
	}
}
