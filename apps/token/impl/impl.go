package impl

import (
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/datasource"
	"github.com/qiaogy91/ioc/config/log"
	"github.com/qiaogy91/mcenter/apps/token"
	"github.com/qiaogy91/mcenter/apps/user"
	"gorm.io/gorm"
	"log/slog"
)

var _ token.Service = &Impl{}

type Impl struct {
	ioc.ObjectImpl
	token.UnimplementedRpcServer
	log *slog.Logger
	db  *gorm.DB
	svc user.Service
}

func (i *Impl) Name() string  { return token.AppName }
func (i *Impl) Priority() int { return 399 }
func (i *Impl) Init() {
	i.log = log.Sub(token.AppName)
	i.db = datasource.DB()
	i.svc = user.GetSvc()
}

func init() {
	ioc.Controller().Registry(&Impl{})
}
