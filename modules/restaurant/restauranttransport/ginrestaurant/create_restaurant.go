package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"tap_code_lai/common"
	"tap_code_lai/component"
	"tap_code_lai/modules/restaurant/resraurantmodel"
	"tap_code_lai/modules/restaurant/restaurantbiz"
	"tap_code_lai/modules/restaurant/restaurantstorage"
)

func CreateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var data resraurantmodel.RestaurantCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// compel type of data Requester ( because Requester is interface)
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		//fmt.Println(requester)
		data.OwnerId = requester.GetUserId()

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDbConnection())
		biz := restaurantbiz.NewCreateRestaurantStore(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		data.GenUID(common.DbTypeRestaurant)

		c.JSON(200, common.SimpleSuccessReponse(data.FakeId))
	}
}
