package impl_test

import (
	"context"
	"github.com/qiaogy91/ioc"
	_ "github.com/qiaogy91/ioc/config/otlp"
	_ "github.com/qiaogy91/mcenter/apps"
	"github.com/qiaogy91/mcenter/apps/token"
	"github.com/qiaogy91/mcenter/apps/token/provider"
	"github.com/qiaogy91/mcenter/apps/token/provider/feishu"
	"testing"
)

var (
	ctx = context.Background()
	c   feishu.Service
)

func init() {
	cf := "/Users/qiaogy/GolandProjects/projects/github/CloudManager/mcenter/etc/application.yaml"
	if err := ioc.ConfigIocObject(cf); err != nil {
		panic(err)
	}

	ins, err := provider.GetSvc().Get(feishu.Type)
	if err != nil {
		panic(err)
	}
	c = ins
}

func TestImpl_IssueToken(t *testing.T) {
	req := &token.IssueTokenRequest{
		IssueType: token.IssueType_ISSUE_TYPE_FEISHU,
		Code:      "6f0m034afdcc41e6842722fa3af6f88b",
		State:     "STATE",
	}

	ins, err := c.IssueToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", ins)
}
