package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tap_code_lai/component"
	"tap_code_lai/modules/restaurant/restaurantbiz"
	"tap_code_lai/modules/restaurant/restaurantstorage"
)

func FindIDRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(400, err)
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDbConnection())
		biz := restaurantbiz.NewFindRestaurantStore(store)
		data, err := biz.FindRestaurant(c.Request.Context(), map[string]interface{}{"id": id})
		if err != nil {
			c.JSON(400, err)
		}

		c.JSON(http.StatusOK, data)
	}
}
