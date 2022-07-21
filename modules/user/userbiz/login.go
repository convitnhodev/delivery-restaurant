package userbiz

import (
	"golang.org/x/net/context"
	"tap_code_lai/component"
	"tap_code_lai/component/tokenprovider"
	"tap_code_lai/modules/user/usermodel"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type TokenConfig interface {
	GetAtExp() int
	GetRtExp() int
}

type loginBusiness struct {
	appCtx        component.AppContext
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	tkCfg         TokenConfig
}

func NewLoginBusiness(storeUser LoginStorage, tokenProvider tokenprovider.Provider, hasher Hasher, tkCfg TokenConfig) *loginBusiness {
	return &loginBusiness{storeUser: storeUser, tokenProvider: tokenProvider, hasher: hasher, tkCfg: tkCfg}
}

func (biz *loginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, err := biz.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		// error handling
		return nil, err
	}

	passHash := biz.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHash {
		// error handling
		return nil, err
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.tkCfg.GetAtExp())

	if err != nil {
		// error handling
		return nil, err
	}
	return accessToken, nil
}
