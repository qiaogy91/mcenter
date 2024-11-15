package api

import (
	"github.com/qiaogy91/ioc/utils"
	"net/http"
)

func ErrRoleCreateParams(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusBadRequest, 10001, "创建角色的参数错误", e)
}

func ErrRoleCreate(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusInternalServerError, 10002, "创建角色过程错误", e)
}

func ErrRoleDeleteParams(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusBadRequest, 10003, "删除角色参数错误", e)
}

func ErrRoleDelete(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusInternalServerError, 10004, "删除角色参内部误", e)
}

func ErrRoleDescParams(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusBadRequest, 10005, "角色详情参数错误", e)
}

func ErrRoleDesc(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusInternalServerError, 10006, "角色详情参内部误", e)
}

func ErrUserQueryParams(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusBadRequest, 10009, "角色列表参数错误", e)
}

func ErrRoleQuery(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusInternalServerError, 10010, "角色列表参内部误", e)
}
