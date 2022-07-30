package restaurantlikemodel

import (
	"tap_code_lai/common"
	"time"
)

const EntityName = "UserLikeRestaurant"

type Like struct {
	RestaurantId int                `json:"restaurant_id" gorm:"column:restaurant_id"`
	UserId       int                `json:"user_id" gorm:"column:user_id"`
	CreatedAt    *time.Time         `json:"create_at" gorm:"column:created_at"`
	User         *common.SimpleUser `json:"user" gorm:"preload:false"`
}

func (l Like) TableName() string {
	return "restaurant_likes"
}
