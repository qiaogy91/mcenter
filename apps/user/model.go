package user

import "golang.org/x/crypto/bcrypt"

func NewUser(req *CreateUserRequest) *User {
	ins := &User{Spec: req}
	ins.MakePasswordHash()
	return ins
}

func (u *User) MakePasswordHash() {
	bs, err := bcrypt.GenerateFromPassword([]byte(u.Spec.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	u.Spec.Password = string(bs)
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Spec.Password), []byte(password))
}
