package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"tap_code_lai/modules/restaurant/resraurantmodel"
	"tap_code_lai/modules/restaurant/restaurantbiz"
	"tap_code_lai/modules/restaurant/restaurantstorage"
)

func CreateRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var data resraurantmodel.RestaurantCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewCreateRestaurantStore(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, &data)
	}
}
