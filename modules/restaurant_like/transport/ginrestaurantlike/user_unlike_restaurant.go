package ginrestaurantlike

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tap_code_lai/common"
	"tap_code_lai/component"
	rslikebiz "tap_code_lai/modules/restaurant_like/biz"
	"tap_code_lai/modules/restaurant_like/storage"
)

//POST/v1/restaurant/:id/unlike

func UserUnLikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := storage.NewSQLStore(appCtx.GetMainDbConnection())
		biz := rslikebiz.NewUserUnLikeRestaurantBiz(store)

		err = biz.UnLikeRestaurant(c.Request.Context(), requester.GetUserId(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessReponse(true))
	}
}
