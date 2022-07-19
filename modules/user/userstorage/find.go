package userstorage

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"tap_code_lai/common"
	"tap_code_lai/modules/user/usermodel"
)

func (s *sqlStore) FindUser(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*usermodel.User, error) {
	db := s.db
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var data usermodel.User
	if err := db.First(&data, conditions).Table(data.TableName()).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
