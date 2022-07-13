package restaurantstorage

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"tap_code_lai/common"
	"tap_code_lai/modules/restaurant/resraurantmodel"
)

func (s *sqlStore) FindByConditions(ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*resraurantmodel.Restaurant, error) {
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var data resraurantmodel.Restaurant
	err := db.First(&data, conditions).Table(data.TableName()).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
