package resraurantmodel

import (
	"errors"
	"strings"
	"tap_code_lai/common"
)

const EntityName = "restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string             `json:"name" gorm:"column:name"`
	UserId          int                `json:"-" gorm:"column:owner_id"`
	Addr            string             `json:"address" gorm:"addr"`
	User            *common.SimpleUser `json:"user" gorm:"preload:false;"`
	LikeCount       int                `json:"like_count" gorm:"column:like_count;"`
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name"`
	UserId          int    `json:"-" gorm:"column:owner_id;"`
	Addr            string `json:"address" gorm:"addr"`
}

type RestaurantUpdate struct {
	common.SQLModel `json:",inline"`
	Name            *string `json:"name" gorm:"column:name"`
	Addr            *string `json:"address" gorm:"addr"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

func (RestaurantUpdate) TableName() string {
	return "restaurants"
}

func (RestaurantCreate) TableName() string {
	return "restaurants"
}

func (res *RestaurantCreate) Validata() error {
	// delete excess space
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return errors.New("restaurant name can not be blank")
	}

	return nil
}

func (data *Restaurant) Mark(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)

	u := data.User
	if u != nil {
		u.Mark(isAdminOrOwner)
	}
}
