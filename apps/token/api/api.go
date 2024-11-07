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
	ws.Route(ws.POST("login").To(h.Login).
		Doc("测试接口").
		Metadata(restfulspec.KeyOpenAPITags, tag),
	)
}
func init() {
	ioc.Api().Registry(&Handler{})
}
