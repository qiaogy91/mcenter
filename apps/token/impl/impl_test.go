package impl_test

import (
	"context"
	"github.com/qiaogy91/ioc"
	_ "github.com/qiaogy91/ioc/config/otlp"
	_ "github.com/qiaogy91/mcenter/apps"
	"github.com/qiaogy91/mcenter/apps/token"
	"testing"
)

var (
	ctx = context.Background()
	c   token.Service
)

func init() {
	if err := ioc.ConfigIocObject("/Users/qiaogy/GolandProjects/projects/github/CloudManager/mcenter/etc/application.yaml"); err != nil {
		panic(err)
	}
	c = token.GetSvc()
}

func TestImpl_CreateTable(t *testing.T) {
	if err := c.CreateTable(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestImpl_IssueToken(t *testing.T) {
	for i := 1; i <= 100; i++ {
		req := &token.IssueTokenRequest{
			Username: "user_66",
			Password: "redhat",
		}

		inst, err := c.IssueToken(ctx, req)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", inst)
	}
}

func TestImpl_ValidateToken(t *testing.T) {
	req := &token.ValidateTokenRequest{AccessToken: "cshmn8is4uq1o285fh6g"}
	inst, err := c.ValidateToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", inst)
}

func TestImpl_QueryToken(t *testing.T) {
	req := &token.QueryTokenRequest{
		PageNum:   2,
		PageSize:  10,
		QueryType: token.QueryType_QUERY_TYPE_USERNAME,
		Keyword:   "user_66",
	}

	inst, err := c.QueryToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	for _, item := range inst.Items {
		t.Logf("%+v", item)
	}
}
