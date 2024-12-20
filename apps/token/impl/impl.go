package impl

import (
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/datasource"
	"github.com/qiaogy91/ioc/config/grpc"
	"github.com/qiaogy91/ioc/config/log"
	"github.com/qiaogy91/mcenter/apps/token"
	"github.com/qiaogy91/mcenter/apps/token/provider"
	"github.com/qiaogy91/mcenter/apps/user"
	"gorm.io/gorm"
	"log/slog"
)

var _ token.Service = &Impl{}

type Impl struct {
	ioc.ObjectImpl
	token.UnimplementedRpcServer
	log      *slog.Logger
	db       *gorm.DB
	usv      user.Service
	Provider provider.Conf `json:"provider" yaml:"provider"`
}

func (i *Impl) Name() string  { return token.AppName }
func (i *Impl) Priority() int { return 303 }
func (i *Impl) Init() {
	i.log = log.Sub(token.AppName)
	i.db = datasource.DB()
	i.usv = user.GetSvc()

	// 各provider 开始初始化
	provider.InitProvider(i.Provider)

	i.grpcRegistry()
}

func (i *Impl) grpcRegistry() {
	s := grpc.Get().Server()
	token.RegisterRpcServer(s, i)
}

func init() {
	ioc.Controller().Registry(&Impl{})
}
