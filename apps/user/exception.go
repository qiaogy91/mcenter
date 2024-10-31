package user

import "github.com/qiaogy91/ioc/utils"

func ErrUserValidate(e error) *utils.ApiException {
	return utils.NewApiException(10001, "用户查询参数错误", e)
}
func ErrUserCreate(e error) *utils.ApiException {
	return utils.NewApiException(10002, "用户创建错误", e)
}
func ErrUserQuery(e error) *utils.ApiException {
	return utils.NewApiException(10003, "用户查询错误", e)
}
func ErrUserDelete(e error) *utils.ApiException {
	return utils.NewApiException(10004, "用户删除错误", e)
}

func ErrUserPasswordCheck(e error) *utils.ApiException {
	return utils.NewApiException(10005, "用户密码校验错误", e)
}

func ErrUserUpdate(e error) *utils.ApiException {
	return utils.NewApiException(10006, "用户更新失败", e)
}
