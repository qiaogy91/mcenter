package endpoint

import (
	"context"
	"fmt"
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/application"
	"github.com/qiaogy91/ioc/config/gorestful"
	"github.com/qiaogy91/ioc/config/log"
	"github.com/qiaogy91/mcenter/apps/endpoint"
	"github.com/qiaogy91/mcenter/provider/client"
	"log/slog"
	"strings"
)

const AppName = "endpointRegistry"

type Impl struct {
	ioc.ObjectImpl
	c   endpoint.RpcClient
	log *slog.Logger
}

func (i *Impl) Name() string  { return AppName }
func (i *Impl) Priority() int { return 999 } // 等待所有 api 注册完成后，再初始化
func (i *Impl) Init() {
	i.log = log.Sub(AppName)
	i.c = client.GetSvc().EndpointClient()

	// 将所有api 进行注册
	i.Registry()
}

func (i *Impl) Registry() {
	ins, err := i.c.CreateEndpoint(context.Background(), i.EndpointSpecSet())
	if err != nil {
		panic(err)
	}

	for _, ep := range ins.Items {
		i.log.Info(fmt.Sprintf("Registry Endpoint %s", ep.Spec.Identity))
	}
}

func (i *Impl) EndpointSpecSet() *endpoint.EndpointSpecSet {
	root := gorestful.RootContainer()
	eps := &endpoint.EndpointSpecSet{}
	svc := application.Get().ApplicationName()

	for _, ws := range root.RegisteredWebServices() {
		for _, r := range ws.Routes() {

			meta := NewMetaData(r.Metadata)
			resource := strings.Trim(ws.RootPath(), "/")
			action := meta.GetString("action")
			doc := meta.GetString("doc")

			// 强制限制路由元数据必须包含 resource、action、doc
			if resource == "" || action == "" || doc == "" {
				continue
			}

			eps.Items = append(eps.Items, &endpoint.Spec{
				Service:  svc,
				Resource: resource,
				Action:   action,
				Identity: fmt.Sprintf("%s.%s.%s.%s", svc, resource, r.Method, r.Path),
				Doc:      doc,
			})
		}
	}
	return eps
}

func init() {
	ioc.Api().Registry(&Impl{})
}
