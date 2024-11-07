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
