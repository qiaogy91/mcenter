package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/gorestful"
	"github.com/qiaogy91/ioc/config/log"
	"github.com/qiaogy91/mcenter/apps/token"
	"log/slog"
)

type Handler struct {
	ioc.ObjectImpl
	log *slog.Logger
	svc token.Service
}

func (h *Handler) Name() string  { return token.AppName }
func (h *Handler) Priority() int { return 401 }
func (h *Handler) Init() {
	h.log = log.Sub(token.AppName)
	h.svc = token.GetSvc()

	h.registry()
}

func (h *Handler) registry() {
	tag := []string{"令牌管理"}

	ws := gorestful.ModuleWebservice(h)
	ws.Route(ws.POST("account/login").To(h.AccountLogin).
		Doc("账号登录").
		Metadata(restfulspec.KeyOpenAPITags, tag),
	)

	ws.Route(ws.GET("feishu/login").To(h.FeishuLogin).
		Doc("飞书登录").
		Metadata(restfulspec.KeyOpenAPITags, tag),
	)

	// todo
	// 1. 颁发令牌
	// 2. 撤销令牌
	// 3. 校验令牌
	// 4. 令牌颁发记录

	for _, ws := range gorestful.RootContainer().RegisteredWebServices() {
		for _, r := range ws.Routes() {
			h.log.Info("Registry", slog.String("doc", r.Doc), slog.String("method", r.Method), slog.String("path", r.Path))
		}
	}
}
func init() {
	ioc.Api().Registry(&Handler{})
}
