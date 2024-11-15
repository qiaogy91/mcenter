package impl

import (
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/datasource"
	"github.com/qiaogy91/ioc/config/grpc"
	"github.com/qiaogy91/ioc/config/log"
	"github.com/qiaogy91/mcenter/apps/endpoint"
	"gorm.io/gorm"
	"log/slog"
)

type Impl struct {
	ioc.ObjectImpl
	endpoint.UnimplementedRpcServer
	log *slog.Logger
	db  *gorm.DB
}

func (i *Impl) Name() string  { return endpoint.AppName }
func (i *Impl) Priority() int { return 301 }
func (i *Impl) Init() {
	i.log = log.Sub(endpoint.AppName)
	i.db = datasource.DB()
	i.RpcRegistry()
}

func (i *Impl) RpcRegistry() {
	s := grpc.Get().Server()
	endpoint.RegisterRpcServer(s, i)
}

func init() {
	ioc.Controller().Registry(&Impl{})
}
