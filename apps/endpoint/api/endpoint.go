package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/qiaogy91/ioc/utils"
	"github.com/qiaogy91/mcenter/apps/endpoint"
)

func (h *Handler) QueryEndpoint(r *restful.Request, w *restful.Response) {
	req := &endpoint.QueryEndpointRequest{}
	if err := utils.Decoder.Decode(req, r.Request.URL.Query()); err != nil {
		utils.SendFailed(w, ErrQueryEndpointParams(err))
		return
	}

	ins, err := h.svc.QueryEndpoint(r.Request.Context(), req)
	if err != nil {
		utils.SendFailed(w, ErrQueryEndpoint(err))
		return
	}

	utils.SendSuccess(w, ins)
}
