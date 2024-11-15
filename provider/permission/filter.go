package permission

import (
	"errors"
	"github.com/emicklei/go-restful/v3"
	"github.com/qiaogy91/ioc/labels"
	"github.com/qiaogy91/ioc/utils"
	"github.com/qiaogy91/mcenter/apps/token"
)

func (i *Impl) Filter() restful.FilterFunction {
	return func(r *restful.Request, w *restful.Response, chain *restful.FilterChain) {
		sr := r.SelectedRoute()
		if sr == nil {
			chain.ProcessFilter(r, w)
			return
		}
		meta := NewMeta(sr.Metadata())
		permEnable := meta.GetBool(labels.PermissionEnabled)
		if !permEnable {
			chain.ProcessFilter(r, w)
			return
		}

		tk := r.Attribute(labels.ContextTokenKey).(*token.Token)
		if !tk.RoleSet.ValidatePermission(i.app, meta.GetString(labels.Resource), meta.GetString(labels.Action)) {
			utils.SendFailed(w, ErrPermUnauthorized(errors.New("权限校验失败")))
			return
		}

		chain.ProcessFilter(r, w)
	}
}
