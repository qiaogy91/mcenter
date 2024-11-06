package authentication

import (
	"github.com/qiaogy91/ioc/utils"
	"net/http"
)

func ErrTokenAuth(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusForbidden, 10001, "Token校验失败", e)
}
