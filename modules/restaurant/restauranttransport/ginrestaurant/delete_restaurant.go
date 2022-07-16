package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tap_code_lai/common"
	"tap_code_lai/component"
	"tap_code_lai/modules/restaurant/restaurantbiz"
	"tap_code_lai/modules/restaurant/restaurantstorage"
)

func DeleteRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//id, err := strconv.Atoi(c.Param("id"))

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			c.JSON(400, err)
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDbConnection())
		biz := restaurantbiz.NewDeleteRestaurantStor(store)

		if err := biz.DeleteRestaurant(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			c.JSON(401, err)
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessReponse(true))

	}
}
