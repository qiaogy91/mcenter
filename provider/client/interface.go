package client

import (
	"github.com/qiaogy91/ioc"
)

const AppName = "mcenterClient"

func GetSvc() *McenterClient { return ioc.Default().Get(AppName).(*McenterClient) }
