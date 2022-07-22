package usermodel

import (
	"tap_code_lai/common"
)

type User struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email"`
	// dont return password
	Password  string `json:"-" gorm:"column:password"`
	LastName  string `json:"last_name" gorm:"column:last_name"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	Phone     string `json:"phone" gorm:"column:phone"`
	Role      string `json:"-" gorm:"column:role"`
	Salt      string `json:"-" gorm:"column:salt"`
	//Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}
func (u *User) GetRole() string {
	return u.Role
}

func (User) TableName() string {
	return "users"
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email"`
	Password        string `json:"password" gorm:"column:password"`
	LastName        string `json:"last_name" gorm:"column:last_name"`
	FirstName       string `json:"first_name" gorm:"column:first_name"`
	Role            string `json:"-" gorm:"column:role"`
	Salt            string `json:"-" gorm:"column:salt"`
	//Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (u UserCreate) TableName() string {
	return "users"
}

func (data *UserCreate) Mark(isAdminOrOwner bool) {
	data.GenUID(common.DbUser)
}

func (data *User) Mark(isAdminOrOwner bool) {
	data.GenUID(common.DbUser)
}

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email" form:"email"`
	Password string `json:"password" form:"password" gorm:"column:password form:password"`
}

func (u UserLogin) TableName() string {
	return "users"
}
