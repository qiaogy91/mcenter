package provider

import (
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/mcenter/apps/token"
)

const AppName = "tokenProvider"

func GetSvc() Service { return ioc.Controller().Get(AppName).(Service) }

type Service interface {
	Registry(o IssueObjectInterface)
	Get(t token.IssueType) (IssueObjectInterface, error)

	GetAppId() string
	GetAppSecret() string
	GetBaseUrl() string
}
