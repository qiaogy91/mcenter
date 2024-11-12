package impl_test

import (
	"context"
	"github.com/qiaogy91/ioc"
	_ "github.com/qiaogy91/ioc/config/otlp"
	_ "github.com/qiaogy91/mcenter/apps"
	"github.com/qiaogy91/mcenter/apps/token"
	"github.com/qiaogy91/mcenter/apps/token/provider"
	"github.com/qiaogy91/mcenter/apps/token/provider/account"
	"testing"
)

var (
	ctx = context.Background()
	c   account.Service
)

func init() {
	if err := ioc.ConfigIocObject("/Users/qiaogy/GolandProjects/projects/github/CloudManager/mcenter/etc/application.yaml"); err != nil {
		panic(err)
	}

	ins, err := provider.GetSvc().Get(account.Type)
	if err != nil {
		panic(err)
	}
	c = ins
}

func TestImpl_IssueToken(t *testing.T) {
	req := &token.IssueTokenRequest{
		IssueType: token.IssueType_ISSUE_TYPE_ACCOUNT,
		Username:  "qiaogy",
		Password:  "redhat",
	}

	ins, err := c.IssueToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", ins)
}
