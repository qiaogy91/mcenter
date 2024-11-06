package client_test

import (
	"context"
	"fmt"
	"github.com/qiaogy91/ioc"
	_ "github.com/qiaogy91/ioc/config/otlp"
	"github.com/qiaogy91/mcenter/apps/endpoint"
	"github.com/qiaogy91/mcenter/apps/token"
	"github.com/qiaogy91/mcenter/provider/client"
	"testing"
)

var (
	ctx = context.Background()
	tc  token.RpcClient
	ec  endpoint.RpcClient
)

func init() {
	if err := ioc.ConfigIocObject("/Users/qiaogy/GolandProjects/projects/github/CloudManager/mcenter/etc/application.yaml"); err != nil {
		panic(err)
	}
	tc = client.GetSvc().TokenClient()
	ec = client.GetSvc().EndpointClient()
}

func TestClient_CreateToken(t *testing.T) {
	req := &token.IssueTokenRequest{
		Username: "user_23",
		Password: "redhat",
	}
	inst, err := tc.IssueToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", inst)
}

func TestClient_ValidateToken(t *testing.T) {
	req := &token.ValidateTokenRequest{AccessToken: "cslg812s4uq3e7gaa130"}
	inst, err := tc.ValidateToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", inst)
}

func TestClient_CreateEndpoint(t *testing.T) {

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
					Verb:     action,
					Identity: fmt.Sprintf("svc_%d.md_%d.%s./md01/api/v1/secret", i, j, action),
					Doc:      "测试",
				}
				specs = append(specs, spec)
			}

		}

	}

	req := &endpoint.EndpointSpecSet{Items: specs}
	ins, err := ec.CreateEndpoint(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	for _, item := range ins.Items {
		t.Logf("%+v", item)
	}
}
