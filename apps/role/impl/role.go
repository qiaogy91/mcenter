package impl

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/qiaogy91/mcenter/apps/role"
	"strings"
)

func (i *Impl) CreateTable(ctx context.Context) error {
	return i.db.WithContext(ctx).AutoMigrate(&role.Role{})
}

func (i *Impl) CreateRole(ctx context.Context, req *role.Spec) (*role.Role, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, err
	}

	ins := &role.Role{Spec: req}
	if err := i.db.WithContext(ctx).Model(&role.Role{}).Create(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

func (i *Impl) DescRole(ctx context.Context, req *role.DescRoleRequest) (*role.Role, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, err
	}
	ins := &role.Role{}
	if err := i.db.WithContext(ctx).Model(&role.Role{}).Where("id = ?", req.Id).First(ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *Impl) QueryRole(ctx context.Context, req *role.QueryRoleRequest) (*role.RoleSet, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, err
	}

	ins := &role.RoleSet{}
	sql := i.db.WithContext(ctx).Model(&role.Role{})

	offset := int((req.PageNum - 1) * req.PageSize)
	limit := int(req.PageSize)

	switch req.QueryType {
	case role.QueryType_QUERY_TYPE_ROLE_NAME:
		sql = sql.Where("name = ?", req.Keyword)
	case role.QueryType_QUERY_TYPE_DESC:
		sql = sql.Where("description like ?", "%"+req.Keyword+"%")
	case role.QueryType_QUERY_TYPE_ROLE_IDS:
		sql = sql.Where("id in ?", strings.Split(req.Keyword, ","))
	}

	if err := sql.Count(&ins.Total).Offset(offset).Limit(limit).Find(&ins.Items).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *Impl) DeleteRole(ctx context.Context, req *role.DeleteRoleRequest) (*role.Role, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, err
	}

	ins, err := i.DescRole(ctx, &role.DescRoleRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}

	if err := i.db.WithContext(ctx).Where("id = ?", req.Id).Delete(&role.Role{}).Error; err != nil {
		return nil, err
	}
	return ins, nil
}
