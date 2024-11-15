package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/qiaogy91/ioc/utils"
	"github.com/qiaogy91/mcenter/apps/endpoint"
	"strconv"
)

func (h *Handler) QueryEndpoint(r *restful.Request, w *restful.Response) {
	num, err := strconv.ParseInt(r.QueryParameter("pageNum"), 10, 64)
	size, err := strconv.ParseInt(r.QueryParameter("pageSize"), 10, 64)
	t, err := strconv.ParseInt(r.QueryParameter("queryType"), 10, 64)
	if err != nil {
		utils.SendFailed(w, ErrQueryEndpointParams(err))
		return
	}

	req := &endpoint.QueryEndpointRequest{
		PageNumber: num,
		PageSize:   size,
		QueryType:  endpoint.QueryType(t),
		Keyword:    r.QueryParameter("keyword"),
	}

	ins, err := h.svc.QueryEndpoint(r.Request.Context(), req)
	if err != nil {
		utils.SendFailed(w, ErrQueryEndpoint(err))
		return
	}

	utils.SendSuccess(w, ins)
}
