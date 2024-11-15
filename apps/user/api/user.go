package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/qiaogy91/ioc/utils"
	"github.com/qiaogy91/mcenter/apps/user"
	"strconv"
)

func (h *Handler) CreateUser(r *restful.Request, w *restful.Response) {
	req := &user.CreateUserRequest{}

	if err := r.ReadEntity(req); err != nil {
		utils.SendFailed(w, ErrUserCreateParams(err))
		return
	}

	ins, err := h.svc.CreateUser(r.Request.Context(), req)
	if err != nil {
		utils.SendFailed(w, ErrUserCreate(err))
		return
	}

	utils.SendSuccess(w, ins)
}

func (h *Handler) DescUser(r *restful.Request, w *restful.Response) {
	_uid := r.PathParameter("uid")
	uid, err := strconv.ParseInt(_uid, 10, 64)
	if err != nil {
		utils.SendFailed(w, ErrUserDescParams(err))
		return
	}
	req := &user.DescUserRequest{Id: uid}
	ins, err := h.svc.DescUser(r.Request.Context(), req)
	if err != nil {
		utils.SendFailed(w, ErrUserDesc(err))
		return
	}

	utils.SendSuccess(w, ins)
}

func (h *Handler) UpdateUser(r *restful.Request, w *restful.Response) {
	_uid := r.PathParameter("uid")
	uid, err := strconv.ParseInt(_uid, 10, 64)
	if err != nil {
		utils.SendFailed(w, ErrUserUpdateParams(err))
		return
	}

	spec := &user.CreateUserRequest{}
	if err := r.ReadEntity(spec); err != nil {
		utils.SendFailed(w, ErrUserCreateParams(err))
		return
	}

	req := &user.UpdateUserRequest{Id: uid, Spec: spec}
	ins, err := h.svc.UpdateUser(r.Request.Context(), req)
	if err != nil {
		utils.SendFailed(w, ErrUserUpdate(err))
		return
	}

	utils.SendSuccess(w, ins)
}

func (h *Handler) QueryUser(r *restful.Request, w *restful.Response) {

	num, err := strconv.ParseInt(r.QueryParameter("pageNum"), 10, 64)
	size, err := strconv.ParseInt(r.QueryParameter("pageSize"), 10, 64)
	t, err := strconv.ParseInt(r.QueryParameter("queryType"), 10, 64)
	if err != nil {
		utils.SendFailed(w, ErrUserQueryParams(err))
		return
	}

	req := &user.QueryUserRequest{
		PageNum:   num,
		PageSize:  size,
		QueryType: user.QueryType(t),
		Keyword:   r.QueryParameter("keyword"),
	}

	ins, err := h.svc.QueryUser(r.Request.Context(), req)
	if err != nil {
		utils.SendFailed(w, ErrUserQuery(err))
		return
	}

	utils.SendSuccess(w, ins)
}

func (h *Handler) DeleteUser(r *restful.Request, w *restful.Response) {
	_uid := r.PathParameter("uid")
	uid, err := strconv.ParseInt(_uid, 10, 64)
	if err != nil {
		utils.SendFailed(w, ErrUserDeleteParams(err))
		return
	}
	req := &user.DeleteUserRequest{Id: uid}
	ins, err := h.svc.DeleteUser(r.Request.Context(), req)
	if err != nil {
		utils.SendFailed(w, ErrUserDelete(err))
		return
	}

	utils.SendSuccess(w, ins)
}
