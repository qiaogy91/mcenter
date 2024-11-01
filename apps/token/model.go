package token

import (
	"github.com/qiaogy91/mcenter/apps/user"
	"github.com/rs/xid"
	"time"
)

func NewToken(u *user.User) *Token {
	return &Token{
		UserId:          u.Meta.Id,
		Username:        u.Spec.Username,
		AccessToken:     xid.New().String(),
		RefreshToken:    xid.New().String(),
		CreatedAt:       time.Now().Unix(),
		AccessExpireAt:  time.Now().Add(1 * time.Hour).Unix(),
		RefreshExpireAt: time.Now().Add(4 * time.Hour).Unix(),
	}
}
