package api

import (
	"github.com/qiaogy91/ioc/utils"
	"net/http"
)

func ErrTokenParams(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusBadRequest, 10001, "请求参数错误", e)
}

func ErrTokenIssue(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusInternalServerError, 10002, "token签发错误", e)
}

func ErrDeleteTokenRequest(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusBadRequest, 10003, "删除令牌参数错误", e)
}

func ErrDeleteToken(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusInternalServerError, 10004, "删除令牌内部错误", e)
}

func ErrQueryTokenRequest(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusBadRequest, 10005, "查询令牌参数错误", e)
}

func ErrQueryToken(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusInternalServerError, 10006, "查询令牌内部错误", e)
}

func ErrValidateTokenRequest(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusBadRequest, 10007, "校验令牌参数错误", e)
}

func ErrValidateToken(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusInternalServerError, 10008, "校验令牌内部错误", e)
}

func ErrRefreshTokenRequest(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusBadRequest, 10009, "刷新令牌参数错误", e)
}

func ErrRefreshToken(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusInternalServerError, 10010, "刷新令牌内部错误", e)
}
func ErrGetTokenRequest(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusBadRequest, 10011, "令牌详情参数错误", e)
}

func ErrGetToken(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusInternalServerError, 10012, "令牌详情内部错误", e)
}
