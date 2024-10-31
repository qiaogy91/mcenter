package token

import "github.com/qiaogy91/ioc/utils"

func ErrTokenCreate(e error) *utils.ApiException {
	return utils.NewApiException(10010, "Token 创建失败", e)
}

func ErrTokenValidate(e error) *utils.ApiException {
	return utils.NewApiException(10011, "Token 参数校验错误", e)
}
func ErrTokenNotFound(e error) *utils.ApiException {
	return utils.NewApiException(10012, "Token 查询错误", e)
}
func ErrTokenExpired(e error) *utils.ApiException {
	return utils.NewApiException(10013, "Token 已经过期", e)
}
func ErrTokenUpdate(e error) *utils.ApiException {
	return utils.NewApiException(10014, "Token 刷新错误", e)
}
func ErrTokenDelete(e error) *utils.ApiException {
	return utils.NewApiException(10015, "Token 删除错误", e)
}

func ErrTokenQuery(e error) *utils.ApiException {
	return utils.NewApiException(10016, "Token 查询错误", e)
}
