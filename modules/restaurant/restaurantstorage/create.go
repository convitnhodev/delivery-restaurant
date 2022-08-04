package restaurantstorage

import (
	"golang.org/x/net/context"
	"tap_code_lai/common"
	"tap_code_lai/modules/restaurant/resraurantmodel"
)

func (s *sqlStore) Create(ctx context.Context, data *resraurantmodel.RestaurantCreate) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
