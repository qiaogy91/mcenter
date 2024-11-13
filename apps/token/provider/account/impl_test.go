package account_test

import (
	"context"
	"github.com/qiaogy91/ioc"
	_ "github.com/qiaogy91/ioc/config/otlp"
	_ "github.com/qiaogy91/mcenter/apps"
	"github.com/qiaogy91/mcenter/apps/token"
	"github.com/qiaogy91/mcenter/apps/token/provider"
	"testing"
)

var (
	ctx = context.Background()
	c   = provider.GetProvider(token.IssueType_ISSUE_TYPE_ACCOUNT)
)

func init() {
	if err := ioc.ConfigIocObject("/Users/qiaogy/GolandProjects/projects/github/CloudManager/mcenter/etc/application.yaml"); err != nil {
		panic(err)
	}
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
