package api

import (
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/gorestful"
	"github.com/qiaogy91/ioc/config/log"
	"github.com/qiaogy91/ioc/labels"
	"github.com/qiaogy91/mcenter/apps/role"
	"github.com/qiaogy91/mcenter/apps/user"
	"log/slog"
)

type Handler struct {
	ioc.ObjectImpl
	log *slog.Logger
	svc role.Service
}

func (h *Handler) Name() string  { return role.AppName }
func (h *Handler) Priority() int { return 402 }
func (h *Handler) Init() {
	h.log = log.Sub(user.AppName)
	h.svc = role.GetSvc()

	h.registry()
}

func (h *Handler) registry() {
	ws := gorestful.ModuleWebservice(h)

	tags := []string{"角色管理"}
	ws.Route(ws.POST("").To(h.CreateRole).
		Doc("角色创建").
		Metadata(labels.ApiTags, tags).
		Reads(role.Spec{}).
		Returns(200, "角色实例", role.Role{}))

	ws.Route(ws.GET("/{uid}").To(h.DescRole).
		Doc("角色详情").
		Metadata(labels.ApiTags, tags).
		Param(ws.PathParameter("uid", "角色ID")).
		Returns(200, "角色实例", role.Role{}))

	ws.Route(ws.GET("").To(h.QueryRole).
		Doc("角色列表").
		Metadata(labels.ApiTags, tags).
		Reads(role.QueryRoleRequest{}).
		Returns(200, "角色列表", role.RoleSet{}))

	ws.Route(ws.DELETE("/{uid}").To(h.DeleteRole).
		Doc("角色删除").
		Metadata(labels.ApiTags, tags).
		Param(ws.PathParameter("uid", "角色ID")).
		Returns(200, "角色实例", role.Role{}))

}

func init() {
	ioc.Api().Registry(&Handler{})
}
