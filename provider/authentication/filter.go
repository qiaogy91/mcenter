package authentication

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/qiaogy91/ioc/utils"
	"github.com/qiaogy91/mcenter/apps/token"
)

// Filter 对每一个Rest 请求，内部调用gRPC 接口去校验
func (i *Impl) Filter() restful.FilterFunction {
	return func(r *restful.Request, w *restful.Response, chain *restful.FilterChain) {
		// 没有匹配到路由
		sr := r.SelectedRoute()
		if sr == nil {
			chain.ProcessFilter(r, w)
			return
		}

		// 元数据
		meta := NewMetaData(sr.Metadata())
		authEnable := meta.GetBool(AuthKey)
		if !authEnable {
			chain.ProcessFilter(r, w)
			return
		}

		// 开启了认证，优先从Header 中读取、如果找不到再取Cookie 中读取
		ak := r.Request.Header.Get(AuthHeader)
		if ak == "" {
			cookie, err := r.Request.Cookie(AttrTokenKey)
			if err != nil {
				utils.SendFailed(w, ErrReadCookie(err))
				return
			}
			ak = cookie.Value
		}

		in := &token.ValidateTokenRequest{AccessToken: ak}
		tk, err := i.c.ValidateToken(r.Request.Context(), in)
		if err != nil {
			utils.SendFailed(w, ErrTokenAuth(err))
			return
		}

		r.SetAttribute(AttrTokenKey, tk)
		chain.ProcessFilter(r, w)
	}
}
