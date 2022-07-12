package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tap_code_lai/component"
	"tap_code_lai/modules/restaurant/resraurantmodel"
	"tap_code_lai/modules/restaurant/restaurantbiz"
	"tap_code_lai/modules/restaurant/restaurantstorage"
)

func FindCity_IDRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter resraurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDbConnection())
		biz := restaurantbiz.NewFind_city_RestaurantStore(store)
		data, err := biz.Find_city_Restaurant(c.Request.Context(), map[string]interface{}{"city_id": filter.CityId})
		if err != nil {
			c.JSON(400, err)
		}

		c.JSON(http.StatusOK, data)
	}
}
