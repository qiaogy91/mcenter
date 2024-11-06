package permission

import (
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/application"
	"github.com/qiaogy91/ioc/config/gorestful"
	"github.com/qiaogy91/ioc/config/log"
	"log/slog"
)

const AppName = "mcenterPermFilter"

type Impl struct {
	ioc.ObjectImpl
	log *slog.Logger
	app string
}

func (i *Impl) Name() string  { return AppName }
func (i *Impl) Priority() int { return 205 }
func (i *Impl) Init() {
	i.log = log.Sub(AppName)
	i.app = application.Get().ApplicationName()

	root := gorestful.RootContainer()
	root.Filter(i.Filter())
	i.log.Info("mcenterPermFilter added")
}

func init() {
	ioc.Default().Registry(&Impl{})
}
