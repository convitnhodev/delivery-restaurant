package storage

import (
	"golang.org/x/net/context"
	"tap_code_lai/common"
	restaurantlikemodel "tap_code_lai/modules/restaurant_like/model"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantlikemodel.Like) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil

}
