package api

import (
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/gorestful"
	"github.com/qiaogy91/ioc/config/log"
	"github.com/qiaogy91/mcenter/apps/token"
	"log/slog"
)

type Handler struct {
	ioc.ObjectImpl
	svc token.Service
	log *slog.Logger
}

func (h *Handler) Name() string  { return token.AppName }
func (h *Handler) Priority() int { return 499 }
func (h *Handler) Init() {
	h.log = log.Sub(token.AppName)
	h.svc = token.GetSvc()

}

func (h *Handler) httpRegistry() {
	ws := gorestful.ModuleWebservice(h)
	ws.Route(ws.POST("").To(h.IssueToken))
}

func init() {
	ioc.Api().Registry(&Handler{})
}
