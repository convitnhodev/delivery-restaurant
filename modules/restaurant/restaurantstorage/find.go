package restaurantstorage

import (
	"golang.org/x/net/context"
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
	err := db.Find(&data, conditions).Table(data.TableName()).Error
	if err != nil {
		return &data, err
	}
	return &data, nil
}
