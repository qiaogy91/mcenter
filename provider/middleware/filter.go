package middleware

import (
	"errors"
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"github.com/qiaogy91/ioc/utils"
	"github.com/qiaogy91/mcenter/apps/token"
)

// Filter 对每一个Rest 请求，内部调用gRPC 接口去校验
func (i *Impl) Filter() restful.FilterFunction {
	return func(r *restful.Request, w *restful.Response, chain *restful.FilterChain) {
		fmt.Printf("我是认证中间件\n")
		// 没有匹配到路由
		sr := r.SelectedRoute()
		if sr == nil {
			utils.SendFailed(w, ErrRouteSelect(errors.New("selected route is nil")))
			return
		}

		// 元数据
		md := NewMetaData(sr.Metadata())

		// 开启认证
		if md.GetBool("auth") {
			ak := r.Request.Header.Get("Authorization")
			in := &token.ValidateTokenRequest{AccessToken: ak}
			tk, err := i.c.ValidateToken(r.Request.Context(), in)
			if err != nil {
				utils.SendFailed(w, ErrTokenAuth(err))
				return
			}

			r.SetAttribute("token", tk)
		}

		chain.ProcessFilter(r, w)
	}
}
