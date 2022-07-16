package restaurantlikemodel

import (
	"time"
)

type Like struct {
	RestaurantId int        `json:"restaurant_id" gorm:"column:restaurant_id"`
	UserId       int        `json:"user_id" gorm:"column:user_id"`
	CreateAt     *time.Time `json:"create_at" gorm:"column:created_at"`
}

func (l Like) TableName() string {
	return "restaurant_likes"
}
