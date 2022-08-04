package ginrestaurantlike

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tap_code_lai/common"
	"tap_code_lai/component"
	rslikebiz "tap_code_lai/modules/restaurant_like/biz"
	restaurantlikemodel "tap_code_lai/modules/restaurant_like/model"
	"tap_code_lai/modules/restaurant_like/storage"
)

//POST/v1/restaurant/:id/likes

func UserLikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.Like{
			UserId:       requester.GetUserId(),
			RestaurantId: int(uid.GetLocalID()),
		}

		store := storage.NewSQLStore(appCtx.GetMainDbConnection())
		biz := rslikebiz.NewUserLikeRestaurantBiz(store)

		err = biz.LikeRestaurant(c.Request.Context(), &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessReponse(true))
	}
}
