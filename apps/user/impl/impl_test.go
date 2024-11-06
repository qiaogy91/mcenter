package impl_test

import (
	"context"
	"fmt"
	"github.com/qiaogy91/ioc"
	_ "github.com/qiaogy91/ioc/config/otlp"
	_ "github.com/qiaogy91/mcenter/apps"
	"github.com/qiaogy91/mcenter/apps/user"
	"testing"
)

var (
	ctx = context.Background()
	c   user.Service
)

func init() {
	if err := ioc.ConfigIocObject("/Users/qiaogy/GolandProjects/projects/github/CloudManager/mcenter/etc/application.yaml"); err != nil {
		panic(err)
	}
	c = user.GetSvc()
}

func TestImpl_CreateTable(t *testing.T) {
	if err := c.CreateTable(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestImpl_CreateUser(t *testing.T) {
	for i := 1; i <= 100; i++ {
		req := &user.CreateUserRequest{
			Username: fmt.Sprintf("user_%d", i),
			Password: "redhat",
		}

		inst, err := c.CreateUser(ctx, req)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", inst)
	}

}
func TestImpl_DescUser(t *testing.T) {
	req := &user.DescUserRequest{Id: 1}
	inst, err := c.DescUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", inst)
}
func TestImpl_UpdateUser(t *testing.T) {
	req := &user.UpdateUserRequest{
		Id: 1,
		Spec: &user.CreateUserRequest{
			Username: "qiaogy",
			Password: "redhat",
			RoleId:   []int64{1, 2, 3},
		},
	}
	ins, err := c.UpdateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", ins)
}
func TestImpl_DeleteUser(t *testing.T) {
	req := &user.DeleteUserRequest{Id: 100}
	inst, err := c.DeleteUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", inst)
}

func TestImpl_QueryUser(t *testing.T) {
	req := &user.QueryUserRequest{
		PageNum:   2,
		PageSize:  10,
		QueryType: user.QueryType_QUERY_TYPE_USERNAME,
		Keyword:   "user",
	}

	inst, err := c.QueryUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	for _, item := range inst.Items {
		t.Logf("%+v", item)
	}
}
