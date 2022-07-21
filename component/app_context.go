package component

import "gorm.io/gorm"

type AppContext interface {
	GetMainDbConnection() *gorm.DB
	SecretKey() string
}

type appCtx struct {
	db        *gorm.DB
	secretKey string
}

func NewAppContext(db *gorm.DB, secretkey string) *appCtx {
	return &appCtx{db, secretkey}
}

func (ctx *appCtx) GetMainDbConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secretKey
}
