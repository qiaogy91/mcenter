package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/qiaogy91/ioc/utils"
	"github.com/qiaogy91/mcenter/apps/token"
)

func (h *Handler) DeleteToken(r *restful.Request, w *restful.Response) {
	req := &token.DeleteTokenRequest{}
	if err := r.ReadEntity(req); err != nil {
		utils.SendFailed(w, ErrDeleteTokenRequest(err))
		return
	}

	ins, err := h.svc.DeleteToken(r.Request.Context(), req)
	if err != nil {
		utils.SendFailed(w, ErrDeleteToken(err))
		return
	}

	utils.SendSuccess(w, ins)
}

func (h *Handler) QueryToken(r *restful.Request, w *restful.Response) {

	req := &token.QueryTokenRequest{}
	if err := utils.Decoder.Decode(req, r.Request.URL.Query()); err != nil {
		utils.SendFailed(w, ErrQueryTokenRequest(err))
		return
	}

	ins, err := h.svc.QueryToken(r.Request.Context(), req)
	if err != nil {
		utils.SendFailed(w, ErrQueryToken(err))
		return
	}

	utils.SendSuccess(w, ins)
}

func (h *Handler) ValidateToken(r *restful.Request, w *restful.Response) {
	req := &token.ValidateTokenRequest{}
	if err := r.ReadEntity(req); err != nil {
		utils.SendFailed(w, ErrValidateTokenRequest(err))
		return
	}

	ins, err := h.svc.ValidateToken(r.Request.Context(), req)
	if err != nil {
		utils.SendFailed(w, ErrValidateToken(err))
		return
	}

	utils.SendSuccess(w, ins)
}

func (h *Handler) RefreshToken(r *restful.Request, w *restful.Response) {
	req := &token.RefreshTokenRequest{}
	if err := r.ReadEntity(req); err != nil {
		utils.SendFailed(w, ErrRefreshTokenRequest(err))
		return
	}

	ins, err := h.svc.RefreshToken(r.Request.Context(), req)
	if err != nil {
		utils.SendFailed(w, ErrRefreshToken(err))
		return
	}
	utils.SendSuccess(w, ins)
}

func (h *Handler) DescToken(r *restful.Request, w *restful.Response) {
	req := &token.DescTokenRequest{AccessToken: r.PathParameter("accessToken")}
	ins, err := h.svc.DescToken(r.Request.Context(), req)
	if err != nil {
		utils.SendFailed(w, ErrGetToken(err))
		return
	}
	utils.SendSuccess(w, ins)
}
