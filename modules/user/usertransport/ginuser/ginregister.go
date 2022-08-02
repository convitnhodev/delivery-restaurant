package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tap_code_lai/component"
	"tap_code_lai/component/hasher"
	"tap_code_lai/modules/user/userbiz"
	"tap_code_lai/modules/user/usermodel"
	"tap_code_lai/modules/user/userstorage"
)

func Register(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDbConnection()

		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			//panic(err)
			c.JSON(http.StatusOK, gin.H{"data": err})
			return
		}

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMD5Hash()
		biz := userbiz.NewRegisterBiz(store, md5)
		if err := biz.Register(c.Request.Context(), &data); err != nil {
			//panic(err)
			c.JSON(http.StatusOK, gin.H{"data": err})
			return
		}

		data.Mark(false)
		c.JSON(http.StatusOK, gin.H{"data": data.FakeId})
	}
}
