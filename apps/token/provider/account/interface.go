package account

import (
	"github.com/qiaogy91/mcenter/apps/token"
	"github.com/qiaogy91/mcenter/apps/token/provider"
)

const (
	AppName = "account"
	Type    = token.IssueType_ISSUE_TYPE_ACCOUNT
)

// GetSvc 从Provider 容器中获取一个实现
func GetSvc() Service {
	ins, err := provider.GetSvc().Get(Type)
	if err != nil {
		panic(err)
	}
	return ins
}

type Service interface {
	provider.IssueObjectInterface
}
