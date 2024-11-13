package feishu

import (
	lark "github.com/larksuite/oapi-sdk-go/v3"
	"github.com/qiaogy91/ioc/config/log"
	"github.com/qiaogy91/mcenter/apps/token"
	"github.com/qiaogy91/mcenter/apps/token/provider"
	"github.com/qiaogy91/mcenter/apps/user"
	"log/slog"
	"time"
)

// 受到 Provider 约束
var _ provider.Provider = &Impl{}

const (
	AppName      = "feiShu"
	ProviderType = token.IssueType_ISSUE_TYPE_FEISHU
)

type Impl struct {
	log *slog.Logger
	c   *lark.Client
	usv user.Service
}

func (i *Impl) Name() string          { return AppName }
func (i *Impl) Type() token.IssueType { return ProviderType }
func (i *Impl) Init(conf provider.Conf) {
	i.log = log.Sub(AppName)
	i.usv = user.GetSvc()
	i.c = lark.NewClient(conf.Feishu.AppID, conf.Feishu.AppSecret,
		lark.WithEnableTokenCache(true),           // 自动缓存 AccessToken
		lark.WithLogReqAtDebug(true),              // 开启 Http 请求参数和响应参数的日志打印开
		lark.WithReqTimeout(5*time.Second),        // 请求调用的超时时间，0表示永不超时
		lark.WithOpenBaseUrl(conf.Feishu.BaseUrl), // 设置飞书域名，私有化的时候需要更改下，比如runwork
	)
}

func init() {
	provider.AddProvider(&Impl{})
}
