package feishu

import (
	"context"
	"errors"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkauthen "github.com/larksuite/oapi-sdk-go/v3/service/authen/v1"
	"github.com/qiaogy91/mcenter/apps/token"
	"github.com/qiaogy91/mcenter/apps/user"
	"github.com/rs/xid"
	"gorm.io/gorm"
	"log/slog"
)

// OidcAccessToken 根据登录预授权码获取 user_access_token
func (i *Impl) oidcAccessToken(ctx context.Context, code string) (*larkauthen.CreateOidcAccessTokenRespData, error) {
	req := larkauthen.NewCreateOidcAccessTokenReqBuilder().
		Body(larkauthen.NewCreateOidcAccessTokenReqBodyBuilder().
			GrantType(`authorization_code`).
			Code(code).
			Build()).
		Build()

	// 发起请求
	ins, err := i.c.Authen.OidcAccessToken.Create(ctx, req)

	// 处理错误
	if err != nil {
		i.log.Error("create err", slog.Any("err", err.Error()))
		return nil, err
	}

	// 服务端错误处理
	if !ins.Success() {
		i.log.Error("OidcAccessToken err", slog.Any("err", ins.CodeError.Error()))
		return nil, ins.CodeError
	}
	return ins.Data, nil

}

// GetUserInfo 通过 user_access_token 获取登录的用户信息
func (i *Impl) getUserInfo(ctx context.Context, token string) (*larkauthen.GetUserInfoRespData, error) {
	ins, err := i.c.Authen.UserInfo.Get(ctx, larkcore.WithUserAccessToken(token))

	// 处理错误
	if err != nil {
		i.log.Error("create err", slog.Any("err", err.Error()))
		return nil, err
	}

	// 服务端错误处理
	if !ins.Success() {
		i.log.Error("OidcAccessToken err", slog.Any("err", ins.CodeError.Error()))
		return nil, ins.CodeError
	}

	return ins.Data, nil
}

// IssueToken 飞书校验
func (i *Impl) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*user.User, error) {
	// 从临时code 码获取用户 access_token
	tk, err := i.oidcAccessToken(ctx, req.Code)
	if err != nil {
		return nil, err
	}

	// 根据access_token 获取用户身份信息
	info, err := i.getUserInfo(ctx, *tk.AccessToken)
	if err != nil {
		return nil, err
	}

	// 同步数据库
	usr, err := i.usv.GetUser(ctx, &user.GetUserRequest{Username: *info.OpenId})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			usr, err = i.usv.CreateUser(ctx, &user.CreateUserRequest{
				Username: *info.OpenId,
				Password: xid.New().String(),
				RoleId:   nil,
			})
			if err != nil {
				return nil, err
			}
			return usr, nil
		}
		return nil, err
	}
	return usr, nil
}
