package userstorage

import (
	"golang.org/x/net/context"
	"tap_code_lai/common"
	"tap_code_lai/modules/user/usermodel"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
	db := s.db.Begin()

	if err := db.Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
