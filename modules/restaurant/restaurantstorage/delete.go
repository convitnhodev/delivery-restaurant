package restaurantstorage

import (
	"golang.org/x/net/context"
	"tap_code_lai/common"
	"tap_code_lai/modules/restaurant/resraurantmodel"
)

func (s *sqlStore) DeleteByCondition(ctx context.Context, conditions map[string]interface{}) error {
	db := s.db

	if err := db.Table(resraurantmodel.Restaurant{}.TableName()).Where(conditions).Update("status", 0).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
