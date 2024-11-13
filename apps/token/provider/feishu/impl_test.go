package feishu_test

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
	c   = provider.GetProvider(token.IssueType_ISSUE_TYPE_FEISHU)
)

func init() {
	cf := "/Users/qiaogy/GolandProjects/projects/github/CloudManager/mcenter/etc/application.yaml"
	if err := ioc.ConfigIocObject(cf); err != nil {
		panic(err)
	}
}

func TestImpl_IssueToken(t *testing.T) {
	req := &token.IssueTokenRequest{
		IssueType: token.IssueType_ISSUE_TYPE_FEISHU,
		Code:      "8f7sa1ad27124bea821ba028a315186a",
		State:     "STATE",
	}

	ins, err := c.IssueToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", ins)
}
