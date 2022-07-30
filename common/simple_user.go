package common

type SimpleUser struct {
	SQLModel `json:",inline"`
	//Email    string `json:"email" gorm:"column:email"`
	// dont return password

	LastName  string `json:"last_name" gorm:"column:last_name"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	Phone     string `json:"phone" gorm:"column:phone"`

	//Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (SimpleUser) TableName() string {
	return "users"
}

func (data *SimpleUser) Mark(isAdminOrOwner bool) {
	data.GenUID(DbUser)
}
