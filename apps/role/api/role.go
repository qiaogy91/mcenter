package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/qiaogy91/ioc/utils"
	"github.com/qiaogy91/mcenter/apps/role"
	"strconv"
)

func (h *Handler) CreateRole(r *restful.Request, w *restful.Response) {
	req := &role.Spec{}
	if err := r.ReadEntity(req); err != nil {
		utils.SendFailed(w, ErrRoleCreateParams(err))
		return
	}

	ins, err := h.svc.CreateRole(r.Request.Context(), req)
	if err != nil {
		utils.SendFailed(w, ErrRoleCreate(err))
		return
	}
	utils.SendSuccess(w, ins)
}

func (h *Handler) DescRole(r *restful.Request, w *restful.Response) {
	_id := r.PathParameter("uid")
	id, err := strconv.ParseInt(_id, 10, 64)
	if err != nil {
		utils.SendFailed(w, ErrRoleDescParams(err))
		return
	}
	req := &role.DescRoleRequest{Id: id}

	ins, err := h.svc.DescRole(r.Request.Context(), req)
	if err != nil {
		utils.SendFailed(w, ErrRoleDesc(err))
		return
	}
	utils.SendSuccess(w, ins)
}

func (h *Handler) QueryRole(r *restful.Request, w *restful.Response) {
	num, err := strconv.ParseInt(r.QueryParameter("pageNum"), 10, 64)
	size, err := strconv.ParseInt(r.QueryParameter("pageSize"), 10, 64)
	t, err := strconv.ParseInt(r.QueryParameter("queryType"), 10, 64)
	if err != nil {
		utils.SendFailed(w, ErrUserQueryParams(err))
		return
	}

	req := &role.QueryRoleRequest{
		PageNum:   num,
		PageSize:  size,
		QueryType: role.QueryType(t),
		Keyword:   r.QueryParameter("keyword"),
	}

	ins, err := h.svc.QueryRole(r.Request.Context(), req)
	if err != nil {
		utils.SendFailed(w, ErrRoleQuery(err))
		return
	}

	utils.SendSuccess(w, ins)

}

func (h *Handler) DeleteRole(r *restful.Request, w *restful.Response) {
	id, err := strconv.ParseInt(r.PathParameter("uid"), 10, 64)
	if err != nil {
		utils.SendFailed(w, ErrRoleDeleteParams(err))
		return
	}

	req := &role.DeleteRoleRequest{Id: id}
	ins, err := h.svc.DeleteRole(r.Request.Context(), req)
	if err != nil {
		utils.SendFailed(w, ErrRoleDelete(err))
		return
	}

	utils.SendSuccess(w, ins)
}
