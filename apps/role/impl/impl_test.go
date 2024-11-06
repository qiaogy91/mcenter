package impl_test

import (
	"context"
	"fmt"
	"github.com/qiaogy91/ioc"
	_ "github.com/qiaogy91/ioc/config/otlp"
	"github.com/qiaogy91/mcenter/apps/role"
	"testing"
)

var (
	ctx = context.Background()
	c   role.Service
)

func init() {
	cf := "/Users/qiaogy/GolandProjects/projects/github/CloudManager/mcenter/etc/application.yaml"
	if err := ioc.ConfigIocObject(cf); err != nil {
		panic(err)
	}
	c = role.GetSvc()
}

func TestImpl_CreateTable(t *testing.T) {
	if err := c.CreateTable(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestImpl_CreateRole(t *testing.T) {
	for i := 1; i <= 3; i++ {
		req := &role.Spec{
			Name:        fmt.Sprintf("role_%d", i),
			Namespace:   "default",
			Service:     []string{"svc01", "svc02"},
			Resource:    []string{"md01", "md02"},
			Actions:     []string{"get", "list"},
			Description: "管理员",
		}

		ins, err := c.CreateRole(ctx, req)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", ins)
	}

}

func TestImpl_DescRole(t *testing.T) {
	req := &role.DescRoleRequest{Id: 1}

	ins, err := c.DescRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", ins)
}

func TestImpl_DeleteRole(t *testing.T) {
	req := &role.DeleteRoleRequest{Id: 1}
	ins, err := c.DeleteRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", ins)
}

func TestImpl_QueryRole(t *testing.T) {
	req := &role.QueryRoleRequest{
		PageNum:   3,
		PageSize:  10,
		QueryType: role.QueryType_QUERY_TYPE_DESC,
		Keyword:   "管理员",
	}

	ins, err := c.QueryRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%d\n", ins.Total)
	for _, item := range ins.Items {
		t.Logf("%+v", item)
	}

}
