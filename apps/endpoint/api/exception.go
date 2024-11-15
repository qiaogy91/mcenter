package api

import (
	"github.com/qiaogy91/ioc/utils"
	"net/http"
)

func ErrQueryEndpointParams(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusBadRequest, 10001, "端点查询参数错误", e)
}

func ErrQueryEndpoint(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusInternalServerError, 10002, "端点查询错误", e)
}
