package ginuser

import (
	"github.com/gin-gonic/gin"
	"tap_code_lai/common"
	"tap_code_lai/component"
	"tap_code_lai/component/hasher"
	"tap_code_lai/modules/user/userbiz"
	"tap_code_lai/modules/user/usermodel"
	"tap_code_lai/modules/user/userstorage"
)

func Login(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			common.ErrInvalidRequest(err)
		}

		db := appCtx.GetMainDbConnection()

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMD5Hash()
		biz := userbiz.NewLoginBusiness(store) tokenProvider, md5, appCtx.GetTokenConfig())
	}
}
