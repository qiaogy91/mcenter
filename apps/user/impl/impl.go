package impl

import (
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/datasource"
	"github.com/qiaogy91/ioc/config/log"
	"github.com/qiaogy91/mcenter/apps/user"
	"gorm.io/gorm"
	"log/slog"
)

var _ user.Service = &Impl{}

type Impl struct {
	ioc.ObjectImpl
	user.UnimplementedRpcServer
	log *slog.Logger
	db  *gorm.DB
}

func (i *Impl) Name() string  { return user.AppName }
func (i *Impl) Priority() int { return 399 }
func (i *Impl) Init() {
	i.log = log.Sub(user.AppName)
	i.db = datasource.DB()
}

func init() {
	ioc.Controller().Registry(&Impl{})
}
