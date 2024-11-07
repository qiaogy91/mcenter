package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/qiaogy91/ioc/utils"
	"github.com/qiaogy91/mcenter/apps/token"
	"net/http"
)

func (h *Handler) Login(r *restful.Request, w *restful.Response) {
	ins := &token.IssueTokenRequest{}
	if err := r.ReadEntity(ins); err != nil {
		utils.SendFailed(w, ErrTokenParams(err))
		return
	}

	tk, err := h.svc.IssueToken(r.Request.Context(), ins)
	if err != nil {
		utils.SendFailed(w, ErrTokenIssue(err))
		return
	}
	http.SetCookie(w.ResponseWriter, &http.Cookie{
		Name:  "token",
		Value: tk.AccessToken,
	})
	utils.SendSuccess(w, tk)
}
