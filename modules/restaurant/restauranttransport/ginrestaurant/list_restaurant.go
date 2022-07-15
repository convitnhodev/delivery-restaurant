package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tap_code_lai/common"
	"tap_code_lai/component"
	"tap_code_lai/modules/restaurant/resraurantmodel"
	"tap_code_lai/modules/restaurant/restaurantbiz"
	"tap_code_lai/modules/restaurant/restaurantstorage"
)

func ListRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter resraurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
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
		biz := restaurantbiz.NewListRestaurantStore(store)
		data, err := biz.ListRestaurant(c.Request.Context(), nil, &filter, &paging)
		if err != nil {
			c.JSON(400, err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessReponse(data))
	}
}
