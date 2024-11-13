package account

import (
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/log"
	"github.com/qiaogy91/mcenter/apps/token"
	"github.com/qiaogy91/mcenter/apps/token/provider"
	"github.com/qiaogy91/mcenter/apps/user"
	"log/slog"
)

var _ provider.Provider = &Impl{}

const (
	AppName      = "account"
	ProviderType = token.IssueType_ISSUE_TYPE_ACCOUNT
)

type Impl struct {
	ioc.ObjectImpl
	log *slog.Logger
	usv user.Service
}

func (i *Impl) Type() token.IssueType { return ProviderType }
func (i *Impl) Name() string          { return AppName }
func (i *Impl) Priority() int         { return 399 }
func (i *Impl) Init(conf provider.Conf) {
	i.log = log.Sub(AppName)
	i.usv = user.GetSvc()
}

func init() {
	provider.AddProvider(&Impl{})
}
