package impl

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/qiaogy91/mcenter/apps/endpoint"
)

func (i *Impl) CreateTable(ctx context.Context) error {
	return i.db.WithContext(ctx).AutoMigrate(&endpoint.Endpoint{})
}
func (i *Impl) CreateEndpoint(ctx context.Context, req *endpoint.EndpointSpecSet) (*endpoint.EndpointSet, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, err
	}

	sql := i.db.WithContext(ctx).Model(&endpoint.Endpoint{})

	// 删除
	if err := sql.Where("service in ?", req.Services()).Delete(&endpoint.Endpoint{}).Error; err != nil {
		return nil, err
	}

	// 新增
	ins := req.EndpointSet()
	if err := sql.CreateInBatches(ins.Items, int(ins.Total)).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *Impl) QueryEndpoint(ctx context.Context, req *endpoint.QueryEndpointRequest) (*endpoint.EndpointSet, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, err
	}

	sql := i.db.WithContext(ctx).Model(&endpoint.Endpoint{})
	ins := &endpoint.EndpointSet{}

	switch req.QueryType {
	case endpoint.QueryType_QUERY_TYPE_SERVICE:
		sql = sql.Where("service = ?", req.Keyword)
	case endpoint.QueryType_QUERY_TYPE_RESOURCE:
		sql = sql.Where("resource = ?", req.Keyword)
	case endpoint.QueryType_QUERY_TYPE_ACTION:
		sql = sql.Where("action = ?", req.Keyword)

	}

	offset := int((req.PageNumber - 1) * req.PageSize)
	limit := int(req.PageSize)

	if err := sql.Offset(offset).Limit(limit).Count(&ins.Total).Find(&ins.Items).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *Impl) DeleteEndpoint(ctx context.Context, req *endpoint.DeleteEndpointRequest) (*endpoint.EndpointSet, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, err
	}

	ins := &endpoint.EndpointSet{}
	sql := i.db.WithContext(ctx).Model(&endpoint.Endpoint{}).Where("service = ?", req.Service)

	if err := sql.Count(&ins.Total).Find(&ins.Items).Delete(&endpoint.Endpoint{}).Error; err != nil {
		return nil, err
	}
	return ins, nil
}
