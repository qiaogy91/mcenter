package impl

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/qiaogy91/mcenter/apps/role"
	"github.com/qiaogy91/mcenter/apps/user"
)

func (i *Impl) CreateTable(ctx context.Context) error {
	return i.db.WithContext(ctx).AutoMigrate(&user.User{})
}

func (i *Impl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.User, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, err
	}

	ins := user.NewUser(req)
	if err := i.db.WithContext(ctx).Model(&user.User{}).Create(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}
func (i *Impl) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.User, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, err
	}

	ins, err := i.DescUser(ctx, &user.DescUserRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}

	ins.Spec = req.Spec
	ins.MakePasswordHash()
	if err := i.db.WithContext(ctx).Model(&user.User{}).Where("id = ?", req.Id).Updates(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

func (i *Impl) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.User, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, err
	}

	inst, err := i.DescUser(ctx, &user.DescUserRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}

	if err := i.db.WithContext(ctx).Model(&user.User{}).Where("id = ?", req.Id).Delete(inst).Error; err != nil {
		return nil, err
	}

	return inst, nil
}

func (i *Impl) DescUser(ctx context.Context, req *user.DescUserRequest) (*user.User, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, err
	}

	ins := &user.User{}
	if err := i.db.WithContext(ctx).Model(&user.User{}).Where("id = ?", req.Id).First(ins).Error; err != nil {
		return nil, err
	}

	in := &role.QueryRoleRequest{
		PageNum:   1,
		PageSize:  1000,
		QueryType: role.QueryType_QUERY_TYPE_ROLE_IDS,
		Ids:       ins.Spec.RoleId,
	}
	roleSet, err := i.role.QueryRole(ctx, in)
	if err != nil {
		return nil, err
	}
	ins.RoleSet = roleSet
	return ins, nil
}

func (i *Impl) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.User, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, err
	}

	ins := &user.User{}
	if err := i.db.WithContext(ctx).Model(&user.User{}).Where("username = ?", req.Username).First(ins).Error; err != nil {
		return nil, err
	}

	return ins, nil

}
func (i *Impl) QueryUser(ctx context.Context, req *user.QueryUserRequest) (*user.UserSet, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, err
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
		return nil, err
	}

	return inst, nil
}
