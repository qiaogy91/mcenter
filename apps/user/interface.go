package user

import (
	"context"
	"github.com/qiaogy91/ioc"
)

const AppName = "user"

func GetSvc() Service { return ioc.Controller().Get(AppName).(Service) }

type Service interface {
	RpcServer
	CreateTable(ctx context.Context) error
}
