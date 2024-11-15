package api

import (
	"github.com/qiaogy91/ioc/utils"
	"net/http"
)

func ErrUserCreateParams(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusBadRequest, 10001, "创建用户的参数错误", e)
}

func ErrUserCreate(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusInternalServerError, 10002, "创建用户过程错误", e)
}

func ErrUserDeleteParams(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusBadRequest, 10003, "删除用户参数错误", e)
}

func ErrUserDelete(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusInternalServerError, 10004, "删除用户参内部误", e)
}

func ErrUserDescParams(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusBadRequest, 10005, "用户详情参数错误", e)
}

func ErrUserDesc(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusInternalServerError, 10006, "用户详情参内部误", e)
}

func ErrUserUpdateParams(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusBadRequest, 10007, "用户更新参数错误", e)
}

func ErrUserUpdate(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusInternalServerError, 10008, "用户更新参内部误", e)
}

func ErrUserQueryParams(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusBadRequest, 10009, "用户列表参数错误", e)
}

func ErrUserQuery(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusInternalServerError, 10010, "用户列表参内部误", e)
}
