package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"tap_code_lai/common"
	"tap_code_lai/component"
	"tap_code_lai/middleware"
	"tap_code_lai/modules/restaurant/restauranttransport/ginrestaurant"
	"tap_code_lai/modules/restaurant_like/transport/ginrestaurantlike"
	"tap_code_lai/modules/user/usertransport/ginuser"
)

func main() {
	dsn := os.Getenv("MYSQL_CONNECTION")
	secretKey := os.Getenv("SECRET_KEY")
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db = db.Debug()

	if err := runService(db, secretKey); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB, secretKey string) error {
	// khong co gi ca
	r := gin.Default()
	// should use os variable
	appCtx := component.NewAppContext(db, secretKey)

	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/v1")

	restaurant := v1.Group("/restaurants", middleware.RequireAuth(appCtx))
	{
		restaurant.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurant.GET("/:id", ginrestaurant.FindIDRestaurant(appCtx))
		restaurant.GET("", ginrestaurant.FindCity_IDRestaurant(appCtx))
		restaurant.GET("/list", ginrestaurant.ListRestaurant(appCtx))
		restaurant.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurant.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))

		restaurant.GET("/:id/liked-users", ginrestaurantlike.ListUser(appCtx))
	}

	user := v1.Group("/users", middleware.RequireAuth(appCtx))
	{
		user.POST("/register", ginuser.Register(appCtx))
		user.POST("/login", ginuser.Login(appCtx))
		user.GET("/profile", ginuser.GetProfile(appCtx))
	}

	v1.GET("/encode-uid", func(c *gin.Context) {
		type reqData struct {
			DbType int `form:"type"`
			RealId int `form:"id"`
		}

		var req reqData
		c.ShouldBind(&req)

		//c.JSON(http.StatusOK, common.SimpleSuccessReponse(req))
		c.JSON(http.StatusOK, gin.H{
			// note
			"uid": common.NewUID(uint32(req.RealId), req.DbType, 1).String(),
		})
	})
	return r.Run()
}
