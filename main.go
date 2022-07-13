package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"tap_code_lai/component"
	"tap_code_lai/middleware"
	"tap_code_lai/modules/restaurant/restauranttransport/ginrestaurant"
)

func main() {
	dsn := os.Getenv("MYSQL_CONNECTION")
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB) error {
	r := gin.Default()
	appCtx := component.NewAppContext(db)

	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/v1")

	restaurant := v1.Group("/restaurants")
	{
		restaurant.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurant.GET("/:id", ginrestaurant.FindIDRestaurant(appCtx))
		restaurant.GET("", ginrestaurant.FindCity_IDRestaurant(appCtx))
		restaurant.GET("/list", ginrestaurant.ListRestaurant(appCtx))
		restaurant.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurant.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	}
	return r.Run()
}
