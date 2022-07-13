package common

import "time"

type SQLModel struct {
	Id       int        `json:"-" gorm:"column:id"`
	Status   int        `json:"status" gorm:"column:status;default:1"`
	CreateAt *time.Time `json:"create_at" gorm:"column:created_at"`
	UpdateAt *time.Time `json:"update_at" gorm:"column:updated_at"`
}
