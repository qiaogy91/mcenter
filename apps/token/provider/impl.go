package provider

import (
	"context"
	"fmt"
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/log"
	"github.com/qiaogy91/mcenter/apps/token"
	"github.com/qiaogy91/mcenter/apps/user"
	"log/slog"
)

// IssueObjectInterface 所有注册到这个容器中的对象，都需要符合这个接口
type IssueObjectInterface interface {
	Name() string          // 名称
	Type() token.IssueType // 类型
	Init()                 // 初始化方法
	IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*user.User, error)
}

// Impl 是一个标准的Ioc 对象
// 同时也是一个存储器对象，用来存储各种 provider 的实现（飞书、钉钉、微信等.....），可通过这个对象来获取各种Provider 的实现
type Impl struct {
	ioc.ObjectImpl
	c   map[token.IssueType]IssueObjectInterface
	log *slog.Logger
	// 飞书认证所需
	AppID     string `json:"appID" yaml:"appID"`
	AppSecret string `json:"appSecret" yaml:"appSecret"`
	BaseUrl   string `json:"baseUrl" yaml:"baseUrl"`
}

func (i *Impl) Name() string  { return AppName }
func (i *Impl) Priority() int { return 300 }
func (i *Impl) Init() {
	i.log = log.Sub(AppName)
	for _, v := range i.c {
		v.Init()
		i.log.Info("Provider init", slog.String("name", v.Name()))
	}
}

func (i *Impl) Registry(o IssueObjectInterface) { i.c[o.Type()] = o }
func (i *Impl) Get(t token.IssueType) (IssueObjectInterface, error) {
	o, ok := i.c[t]
	if !ok {
		return nil, fmt.Errorf("%s not registry", t)
	}
	return o, nil
}

// 剥离出飞书需要的字段
func (i *Impl) GetAppId() string     { return i.AppID }
func (i *Impl) GetAppSecret() string { return i.AppSecret }
func (i *Impl) GetBaseUrl() string   { return i.BaseUrl }

func init() {
	ioc.Controller().Registry(&Impl{c: make(map[token.IssueType]IssueObjectInterface)})
}
