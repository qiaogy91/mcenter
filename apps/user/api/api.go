package api

import (
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/gorestful"
	"github.com/qiaogy91/ioc/config/log"
	"github.com/qiaogy91/ioc/labels"
	"github.com/qiaogy91/mcenter/apps/user"
	"log/slog"
)

type Handler struct {
	ioc.ObjectImpl
	log *slog.Logger
	svc user.Service
}

func (h *Handler) Name() string  { return user.AppName }
func (h *Handler) Priority() int { return 404 }
func (h *Handler) Init() {
	h.log = log.Sub(user.AppName)
	h.svc = user.GetSvc()

	h.registry()
}

func (h *Handler) registry() {
	ws := gorestful.ModuleWebservice(h)

	tags := []string{"用户管理"}

	ws.Route(ws.POST("/").To(h.CreateUser).
		Doc("用户创建").
		Metadata(labels.ApiTags, tags).
		Reads(user.CreateUserRequest{}).
		Returns(200, "userInstance", user.User{}))

	ws.Route(ws.DELETE("/{uid}").To(h.DeleteUser).
		Doc("用户删除").
		Metadata(labels.ApiTags, tags).
		Param(ws.PathParameter("uid", "用户ID")).
		Returns(200, "userInstance", user.User{}))

	ws.Route(ws.GET("/{uid}").To(h.DescUser).
		Doc("用户详情").
		Metadata(labels.ApiTags, tags).
		Param(ws.PathParameter("uid", "用户ID")).
		Returns(200, "userInstance", user.User{}))

	ws.Route(ws.PATCH("/{uid}").To(h.UpdateUser).
		Doc("用户更新").
		Metadata(labels.ApiTags, tags).
		Param(ws.PathParameter("uid", "用户ID")).
		Returns(200, "userInstance", user.User{}))

	ws.Route(ws.GET("/").To(h.QueryUser).
		Doc("用户列表").
		Metadata(labels.ApiTags, tags).
		Reads(user.QueryUserRequest{}).
		Returns(200, "userInstance", user.UserSet{}))
}

func init() {
	ioc.Api().Registry(&Handler{})
}
