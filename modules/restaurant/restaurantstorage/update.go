package restaurantstorage

import (
	"golang.org/x/net/context"
	"tap_code_lai/modules/restaurant/resraurantmodel"
)

func (s *sqlStore) UpdateByCondition(ctx context.Context,
	data *resraurantmodel.RestaurantUpdate,
	conditons map[string]interface{}) error {
	db := s.db
	if err := db.Table(resraurantmodel.RestaurantUpdate{}.TableName()).Where(conditons).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
