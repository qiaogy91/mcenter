package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/gorestful"
	"github.com/qiaogy91/ioc/config/log"
	"github.com/qiaogy91/ioc/labels"
	"github.com/qiaogy91/mcenter/apps/token"
	"log/slog"
)

type Handler struct {
	ioc.ObjectImpl
	log *slog.Logger
	svc token.Service
}

func (h *Handler) Name() string  { return token.AppName }
func (h *Handler) Priority() int { return 403 }
func (h *Handler) Init() {
	h.log = log.Sub(token.AppName)
	h.svc = token.GetSvc()

	h.registry()
}

func (h *Handler) registry() {
	tags := []string{"令牌管理"}

	// 颁发令牌
	ws := gorestful.ModuleWebservice(h)
	ws.Route(ws.POST("account/login").To(h.AccountLogin).
		Doc("颁发令牌(账号登录)").
		Reads(token.IssueTokenRequest{}).
		Metadata(restfulspec.KeyOpenAPITags, tags),
	)
	ws.Route(ws.POST("feishu/login").To(h.FeishuLogin).
		Doc("颁发令牌(飞书登录)").
		Reads(token.IssueTokenRequest{}).
		Metadata(restfulspec.KeyOpenAPITags, tags),
	)

	// 令牌撤销
	ws.Route(ws.DELETE("").To(h.DeleteToken).
		Doc("令牌撤销").
		Metadata(labels.ApiTags, tags).
		Reads(token.DeleteTokenRequest{}).
		Returns(200, "令牌对象", token.Token{}))

	// 令牌查询
	ws.Route(ws.GET("").To(h.QueryToken).
		Doc("令牌查询").
		Metadata(labels.ApiTags, tags).
		Reads(token.QueryTokenRequest{}).
		Returns(200, "令牌列表", token.TokenSet{}))

	// 令牌详情
	ws.Route(ws.GET("/{accessToken}").To(h.DescToken).
		Doc("令牌详情").
		Metadata(labels.ApiTags, tags).
		Reads(token.DescTokenRequest{}).
		Returns(200, "令牌对象", token.Token{}))

	// 令牌校验
	ws.Route(ws.POST("/validate").To(h.ValidateToken).
		Doc("令牌校验").
		Metadata(labels.ApiTags, tags).
		Reads(token.ValidateTokenRequest{}).
		Returns(200, "令牌对象", token.Token{}))

	// 令牌刷新
	ws.Route(ws.POST("/refresh").To(h.RefreshToken).
		Doc("令牌刷新").
		Metadata(labels.ApiTags, tags).
		Reads(token.RefreshTokenRequest{}).
		Returns(200, "令牌对象", token.Token{}))
}

func init() {
	ioc.Api().Registry(&Handler{})
}
