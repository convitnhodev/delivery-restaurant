package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"tap_code_lai/modules/restaurant/restauranttransport/ginrestaurant"
)

func main() {
	dsn := "foot_delivery:Thaothaothao2230@tcp(localhost:3306)/foot_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB) error {
	r := gin.Default()

	restaurant := r.Group("/restaurants")
	{
		restaurant.POST("", ginrestaurant.CreateRestaurant(db))

	}

	return r.Run()
}
