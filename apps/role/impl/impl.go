package impl

import (
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/datasource"
	"github.com/qiaogy91/ioc/config/grpc"
	"github.com/qiaogy91/ioc/config/log"
	"github.com/qiaogy91/mcenter/apps/role"
	"gorm.io/gorm"
	"log/slog"
)

type Impl struct {
	ioc.ObjectImpl
	role.UnimplementedRpcServer
	log *slog.Logger
	db  *gorm.DB
}

func (i *Impl) Name() string  { return role.AppName }
func (i *Impl) Priority() int { return 388 }
func (i *Impl) Init() {
	i.log = log.Sub(role.AppName)
	i.db = datasource.DB()

	i.gRpcRegistry()
}

func (i *Impl) gRpcRegistry() {
	s := grpc.Get().Server()
	role.RegisterRpcServer(s, i)
}

func init() {
	ioc.Controller().Registry(&Impl{})
}
