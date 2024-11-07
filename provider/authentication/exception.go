package authentication

import (
	"github.com/qiaogy91/ioc/utils"
	"net/http"
)

func ErrTokenAuth(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusForbidden, 10001, "Token认证失败", e)
}

func ErrReadCookie(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusForbidden, 10002, "Cookie读取错误", e)
}
