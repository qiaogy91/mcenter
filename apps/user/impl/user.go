package impl

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/qiaogy91/mcenter/apps/user"
)

func (i *Impl) CreateTable(ctx context.Context) error {
	return i.db.WithContext(ctx).AutoMigrate(&user.User{})
}

func (i *Impl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.User, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, user.ErrUserValidate(err)
	}

	ins := user.NewUser(req)
	if err := i.db.WithContext(ctx).Model(&user.User{}).Create(ins).Error; err != nil {
		return nil, user.ErrUserCreate(err)
	}
	return ins, nil
}

func (i *Impl) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.User, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, user.ErrUserValidate(err)
	}

	inst, err := i.GetUser(ctx, &user.GetUserRequest{Username: req.Username})
	if err != nil {
		return nil, err
	}

	if err := i.db.WithContext(ctx).Model(&user.User{}).Where("username = ?", req.Username).Delete(inst).Error; err != nil {
		return nil, user.ErrUserDelete(err)
	}

	return inst, nil
}

func (i *Impl) UpdatePassword(ctx context.Context, req *user.UpdatePasswordRequest) (*user.User, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, user.ErrUserValidate(err)
	}

	inst, err := i.GetUser(ctx, &user.GetUserRequest{Username: req.Username})
	if err != nil {
		return nil, err
	}

	if err := inst.CheckPassword(req.Password); err != nil {
		return nil, user.ErrUserPasswordCheck(err)
	}

	if req.Password == req.NewPassword {
		return nil, user.ErrUserPasswordCheck(errors.New("密码未修改"))
	}

	inst.Spec.Password = req.NewPassword
	inst.MakePasswordHash()

	if err := i.db.WithContext(ctx).Model(&user.User{}).Where("username = ?", req.Username).Updates(inst).Error; err != nil {
		return nil, user.ErrUserUpdate(err)
	}

	return inst, nil
}

func (i *Impl) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.User, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, user.ErrUserValidate(err)
	}

	ins := &user.User{}
	if err := i.db.WithContext(ctx).Model(&user.User{}).Where("username = ?", req.Username).First(ins).Error; err != nil {
		return nil, user.ErrUserQuery(err)
	}
	return ins, nil
}

func (i *Impl) QueryUser(ctx context.Context, req *user.QueryUserRequest) (*user.UserSet, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, user.ErrUserValidate(err)
	}

	sql := i.db.WithContext(ctx).Model(&user.User{})
	offset := int((req.PageNum - 1) * req.PageSize)
	limit := int(req.PageSize)

	switch req.QueryType {
	case user.QueryType_QUERY_TYPE_USERNAME:
		sql = sql.Where("username like ?", "%"+req.Keyword+"%")
	}

	inst := &user.UserSet{}
	if err := sql.Count(&inst.Total).Offset(offset).Limit(limit).Find(&inst.Items).Error; err != nil {
		return nil, user.ErrUserQuery(err)
	}

	return inst, nil
}
