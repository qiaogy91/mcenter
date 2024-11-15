package api

import (
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/gorestful"
	"github.com/qiaogy91/ioc/config/log"
	"github.com/qiaogy91/ioc/labels"
	"github.com/qiaogy91/mcenter/apps/endpoint"
	"github.com/qiaogy91/mcenter/apps/token"
	"log/slog"
)

type Handler struct {
	ioc.ObjectImpl
	log *slog.Logger
	svc endpoint.Service
}

func (h *Handler) Name() string  { return endpoint.AppName }
func (h *Handler) Priority() int { return 401 }
func (h *Handler) Init() {
	h.log = log.Sub(token.AppName)
	h.svc = endpoint.GetSvc()

	h.registry()
}

func (h *Handler) registry() {
	tags := []string{"端点管理"}

	// 颁发令牌
	ws := gorestful.ModuleWebservice(h)

	// 令牌撤销
	ws.Route(ws.GET("/").To(h.QueryEndpoint).
		Doc("端点查询").
		Metadata(labels.ApiTags, tags).
		Reads(endpoint.QueryEndpointRequest{}).
		Returns(200, "端点列表", endpoint.EndpointSet{}))

}

func init() {
	ioc.Api().Registry(&Handler{})
}
