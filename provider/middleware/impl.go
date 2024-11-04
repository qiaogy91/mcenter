package middleware

import (
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/gorestful"
	"github.com/qiaogy91/ioc/config/log"
	"github.com/qiaogy91/mcenter/apps/token"
	"github.com/qiaogy91/mcenter/provider/client"
	"log/slog"
)

const AppName = "mcenterFilter"

type Impl struct {
	ioc.ObjectImpl
	c   token.RpcClient
	log *slog.Logger
}

func (i *Impl) Name() string  { return AppName }
func (i *Impl) Priority() int { return 202 }
func (i *Impl) Init() {
	i.log = log.Sub(AppName)
	i.c = client.GetSvc().TokenClient()

	ws := gorestful.RootContainer()
	ws.Filter(i.Filter())
	i.log.Info("McenterFilter added")
}

func init() {
	ioc.Default().Registry(&Impl{})
}
