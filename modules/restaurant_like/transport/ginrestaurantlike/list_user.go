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

//get/v1/restaurants/:id/likes-users

func ListUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter := restaurantlikemodel.Filter{
			RestaurantId: int(uid.GetLocalID()),
		}

		//id, err := strconv.Atoi(c.Param("id"))

		//
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		// set default
		paging.Fullfill()

		store := storage.NewSQLStore(appCtx.GetMainDbConnection())
		biz := rslikebiz.NewListUserLikeRestaurantBiz(store)
		data, err := biz.ListUsers(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range data {
			data[i].Mark(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessReponse(data, paging, filter))
	}
}
