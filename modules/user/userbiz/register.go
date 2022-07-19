package userbiz

import (
	"golang.org/x/net/context"
	"tap_code_lai/common"
	"tap_code_lai/modules/user/usermodel"
)

type RegisterStore interface {
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
	FindUser(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*usermodel.User, error)
}

// hash password
type Hasher interface {
	Hash(data string) string
}

type RegisterBiz struct {
	registerStore RegisterStore
	hasher        Hasher
}

func NewRegisterBiz(registerStore RegisterStore, hasher Hasher) *RegisterBiz {
	return &RegisterBiz{registerStore, hasher}
}

func (biz *RegisterBiz) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, err := biz.registerStore.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return common.ErrEntityExisted("User Register", err)
	}

	// random salt
	salt := common.GenSalt(50)

	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user" // hard code
	data.Status = 1

	if err := biz.registerStore.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity("User Register", err)
	}
	return nil
}
