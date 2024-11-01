package impl

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/qiaogy91/mcenter/apps/token"
	"github.com/qiaogy91/mcenter/apps/user"
	"time"
)

func (i *Impl) CreateTable(ctx context.Context) error {
	return i.db.WithContext(ctx).AutoMigrate(&token.Token{})
}

func (i *Impl) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, token.ErrTokenValidate(err)
	}

	r := &user.GetUserRequest{Username: req.Username}
	u, err := i.svc.GetUser(ctx, r)
	if err != nil {
		return nil, err
	}

	if err := u.CheckPassword(req.Password); err != nil {
		return nil, err
	}

	inst := token.NewToken(u)
	if err := i.db.WithContext(ctx).Model(&token.Token{}).Create(inst).Error; err != nil {
		return nil, token.ErrTokenCreate(err)
	}
	return inst, nil
}

func (i *Impl) DeleteToken(ctx context.Context, req *token.DeleteTokenRequest) (*token.Token, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, token.ErrTokenValidate(err)
	}

	inst, err := i.GetToken(ctx, &token.GetTokenRequest{AccessToken: req.AccessToken})
	if err != nil {
		return nil, err
	}

	if err := i.db.WithContext(ctx).Where("access_token = ? AND refresh_token = ?", req.AccessToken, req.RefreshToken).
		Delete(&token.Token{}).
		Error; err != nil {
		return nil, token.ErrTokenDelete(err)
	}

	return inst, nil
}

func (i *Impl) GetToken(ctx context.Context, req *token.GetTokenRequest) (*token.Token, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, token.ErrTokenValidate(err)
	}

	inst := &token.Token{}
	if err := i.db.WithContext(ctx).Model(&token.Token{}).Where("access_token = ?", req.AccessToken).First(inst).Error; err != nil {
		return nil, token.ErrTokenNotFound(err)
	}

	return inst, nil
}

func (i *Impl) RefreshToken(ctx context.Context, req *token.RefreshTokenRequest) (*token.Token, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, token.ErrTokenValidate(err)
	}

	inst, err := i.GetToken(ctx, &token.GetTokenRequest{AccessToken: req.AccessToken})
	if err != nil {
		return nil, err
	}

	// 延长token过期时间
	inst.AccessExpireAt = time.Now().Add(1 * time.Hour).Unix()
	inst.RefreshExpireAt = time.Now().Add(4 * time.Hour).Unix()

	if err := i.db.Debug().WithContext(ctx).Model(&token.Token{}).Where("access_token = ?", inst.AccessToken).Updates(inst).Error; err != nil {
		return nil, token.ErrTokenUpdate(err)
	}

	return inst, nil
}

func (i *Impl) QueryToken(ctx context.Context, req *token.QueryTokenRequest) (*token.TokenSet, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, token.ErrTokenValidate(err)
	}

	sql := i.db.WithContext(ctx).Model(&token.Token{})
	inst := &token.TokenSet{}
	offset := int((req.PageNum - 1) * req.PageSize)
	limit := int(req.PageSize)

	switch req.QueryType {
	case token.QueryType_QUERY_TYPE_USERNAME:
		u, err := i.svc.GetUser(ctx, &user.GetUserRequest{Username: req.Keyword})
		if err != nil {
			return nil, err
		}
		sql = sql.Where("user_id = ?", u.Meta.Id)
	}

	if err := sql.Count(&inst.Total).Offset(offset).Limit(limit).Find(&inst.Items).Error; err != nil {
		return nil, token.ErrTokenQuery(err)
	}
	return inst, nil
}

func (i *Impl) ValidateToken(ctx context.Context, req *token.ValidateTokenRequest) (*token.Token, error) {
	// 1. 判断是否存在该token
	inst, err := i.GetToken(ctx, &token.GetTokenRequest{AccessToken: req.AccessToken})
	if err != nil {
		return nil, err
	}

	// 2. AccessToken 未过期
	if inst.AccessExpireAt > time.Now().Unix() {
		return inst, nil
	}

	// 2. RefreshToken 未过期
	if inst.RefreshExpireAt > time.Now().Unix() {
		in := &token.RefreshTokenRequest{AccessToken: inst.AccessToken, RefreshToken: inst.RefreshToken}
		inst, err = i.RefreshToken(ctx, in)
		if err != nil {
			return nil, err
		}
		return inst, nil
	}

	// 3. 过期后删除
	in := &token.DeleteTokenRequest{
		AccessToken:  inst.AccessToken,
		RefreshToken: inst.RefreshToken,
	}
	if _, err := i.DeleteToken(ctx, in); err != nil {
		return nil, err
	}

	return nil, token.ErrTokenExpired(errors.New("token expired"))
}
