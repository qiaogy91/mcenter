package permission

import (
	"github.com/qiaogy91/ioc/utils"
	"net/http"
)

func ErrPermUnauthorized(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusUnauthorized, 10001, "权限校验失败", e)
}
