package role

import (
	"context"
	"github.com/qiaogy91/ioc"
	"slices"
)

const AppName = "role"

func GetSvc() Service { return ioc.Controller().Get(AppName).(Service) }

type Service interface {
	RpcServer
	CreateTable(ctx context.Context) error
}

// ValidatePermission 校验给定的 service、resource、action 是否在权限允许的范围之内
func (rs *RoleSet) ValidatePermission(v, s, a string) bool {

	// 任意一个不符合，就进入下一轮判断
	for _, r := range rs.Items {
		if !slices.Contains(r.Spec.Service, v) && !slices.Contains(r.Spec.Service, "*") {
			continue
		}
		if !slices.Contains(r.Spec.Resource, s) && !slices.Contains(r.Spec.Resource, "*") {
			continue
		}
		if !slices.Contains(r.Spec.Actions, a) && !slices.Contains(r.Spec.Actions, "*") {
			continue
		}
		return true
	}
	return false
}
