package resraurantmodel

import (
	"errors"
	"strings"
	"tap_code_lai/common"
)

const EntityName = "restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name"`
	Addr            string `json:"address" gorm:"addr"`
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name"`
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
