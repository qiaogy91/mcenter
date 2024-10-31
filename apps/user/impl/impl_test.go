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

func TestImpl_GetUser(t *testing.T) {
	req := &user.GetUserRequest{Username: "qiaogy"}
	inst, err := c.GetUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", inst)
}

func TestImpl_DeleteUser(t *testing.T) {
	req := &user.DeleteUserRequest{Username: "qiaogy"}
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

func TestImpl_UpdatePassword(t *testing.T) {
	req := &user.UpdatePasswordRequest{
		Username:    "user_12",
		Password:    "redhat",
		NewPassword: "redhat.123",
	}

	inst, err := c.UpdatePassword(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", inst)
}

func TestCheckPassword(t *testing.T) {
	inst, err := c.GetUser(ctx, &user.GetUserRequest{Username: "user_12"})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", inst)

	if err := inst.CheckPassword("redhat.123"); err != nil {
		t.Fatal(err)
	}

	t.Logf("密码正确")
}
