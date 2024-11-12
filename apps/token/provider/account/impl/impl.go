package impl

import (
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/log"
	"github.com/qiaogy91/mcenter/apps/token"
	"github.com/qiaogy91/mcenter/apps/token/provider"
	"github.com/qiaogy91/mcenter/apps/token/provider/account"
	"github.com/qiaogy91/mcenter/apps/token/provider/feishu"
	"github.com/qiaogy91/mcenter/apps/user"
	"log/slog"
)

type Impl struct {
	ioc.ObjectImpl
	log *slog.Logger
	usv user.Service
}

func (i *Impl) Type() token.IssueType { return account.Type }
func (i *Impl) Name() string          { return account.AppName }
func (i *Impl) Priority() int         { return 399 }
func (i *Impl) Init() {
	i.log = log.Sub(feishu.AppName)
	i.usv = user.GetSvc()
}

func init() {
	provider.GetSvc().Registry(&Impl{})
}
