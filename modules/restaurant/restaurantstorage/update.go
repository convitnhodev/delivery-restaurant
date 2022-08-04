package restaurantstorage

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"tap_code_lai/common"
	"tap_code_lai/modules/restaurant/resraurantmodel"
)

func (s *sqlStore) UpdateByCondition(ctx context.Context,
	data *resraurantmodel.RestaurantUpdate,
	conditons map[string]interface{}) error {
	db := s.db
	if err := db.Table(resraurantmodel.RestaurantUpdate{}.TableName()).Where(conditons).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStore) IncreaseLikeCount(ctx context.Context, id int) error {
	db := s.db
	if err := db.Table(resraurantmodel.Restaurant{}.TableName()).Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStore) DecreaseLikeCount(ctx context.Context, id int) error {
	db := s.db
	if err := db.Table(resraurantmodel.Restaurant{}.TableName()).Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
