package impl

import (
	"context"
	"github.com/qiaogy91/mcenter/apps/token"
	"github.com/qiaogy91/mcenter/apps/user"
)

// IssueToken 飞书校验
func (i *Impl) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*user.User, error) {
	r := &user.GetUserRequest{Username: req.Username}
	u, err := i.usv.GetUser(ctx, r)
	if err != nil {
		return nil, err
	}

	if err := u.CheckPassword(req.Password); err != nil {
		return nil, err
	}

	return u, nil
}
