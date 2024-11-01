package client_test

import (
	"context"
	"github.com/qiaogy91/ioc"
	_ "github.com/qiaogy91/ioc/config/otlp"
	"github.com/qiaogy91/mcenter/apps/token"
	"github.com/qiaogy91/mcenter/provider/client"
	"testing"
)

var (
	ctx = context.Background()
	c   token.RpcClient
)

func init() {
	if err := ioc.ConfigIocObject("/Users/qiaogy/GolandProjects/projects/github/CloudManager/mcenter/etc/application.yaml"); err != nil {
		panic(err)
	}
	c = client.GetSvc().TokenClient()
}

func TestClient_CreateToken(t *testing.T) {
	req := &token.IssueTokenRequest{
		Username: "user_23",
		Password: "redhat",
	}
	inst, err := c.IssueToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", inst)
}

func TestClient_ValidateToken(t *testing.T) {
	req := &token.ValidateTokenRequest{AccessToken: "csi4342s4uq61n13k4dg"}
	inst, err := c.ValidateToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", inst)
}
