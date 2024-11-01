package client

import (
	"context"
	"fmt"
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/mcenter/apps/token"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Credential 为每个RPC 请求提供凭据
type Credential struct{ token string }

func (cr *Credential) RequireTransportSecurity() bool { return false }
func (cr *Credential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"Token": cr.token}, nil
}

// McenterClient 客户端对象
type McenterClient struct {
	ioc.ObjectImpl
	Addr  string `json:"addr" yaml:"addr"`
	Port  int    `json:"port" yaml:"port"`
	Token string `json:"token" yaml:"token"`
	conn  *grpc.ClientConn
}

func (mc *McenterClient) Name() string  { return AppName }
func (mc *McenterClient) Priority() int { return 201 }
func (mc *McenterClient) Init() {
	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", mc.Addr, mc.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()), // 非TLS
		grpc.WithPerRPCCredentials(&Credential{token: mc.Token}), // 凭据
	)

	if err != nil {
		panic(err)
	}
	mc.conn = conn
}

func (mc *McenterClient) TokenClient() token.RpcClient { return token.NewRpcClient(mc.conn) }

func init() {
	ioc.Default().Registry(&McenterClient{})
}
