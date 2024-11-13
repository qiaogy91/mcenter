package provider

import (
	"context"
	"fmt"
	"github.com/qiaogy91/ioc/config/log"
	"github.com/qiaogy91/mcenter/apps/token"
	"github.com/qiaogy91/mcenter/apps/user"
	"log/slog"
)

const AppName = "tokenProvider"

// Provider 每种实现的接口定义
type Provider interface {
	Name() string
	Init(conf Conf)                                                                   // 初始化，
	Type() token.IssueType                                                            // 有多个Provider 的实现，用来唯一区别每种实现
	IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*user.User, error) // 每种Provider 都需要实现的方法
}

var (
	container   = make(map[token.IssueType]Provider)
	AddProvider = func(o Provider) { container[o.Type()] = o }
	GetProvider = func(t token.IssueType) Provider {
		v, ok := container[t]
		if !ok {
			panic(fmt.Errorf("get provider err, because provider not added"))
		}
		return v
	}

	InitProvider = func(conf Conf) {
		l := log.Sub(AppName)
		for _, o := range container {
			o.Init(conf)
			l.Info("add provider", slog.String("name", o.Name()), slog.Any("type", o.Type()))
		}
	}
)
