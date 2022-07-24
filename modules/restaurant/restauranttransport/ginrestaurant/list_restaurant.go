package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tap_code_lai/common"
	"tap_code_lai/component"
	"tap_code_lai/modules/restaurant/resraurantmodel"
	"tap_code_lai/modules/restaurant/restaurantbiz"
	"tap_code_lai/modules/restaurant/restaurantstorage"
	"tap_code_lai/modules/restaurant_like/storage"
)

func ListRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter resraurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		// set default
		paging.Fullfill()

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDbConnection())
		storeLike := storage.NewSQLStore(appCtx.GetMainDbConnection())
		biz := restaurantbiz.NewListRestaurantBiz(store, storeLike)
		data, err := biz.ListRestaurant(c.Request.Context(), nil, &filter, &paging, "User")
		if err != nil {
			panic(err)
		}

		for i := range data {
			data[i].Mark(false)

		}

		paging.NextCursor = data[len(data)-1].FakeId.String()

		c.JSON(http.StatusOK, common.NewSuccessReponse(data, paging, filter))
	}
}
