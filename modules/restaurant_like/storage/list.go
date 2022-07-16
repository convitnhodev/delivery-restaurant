package storage

import (
	"golang.org/x/net/context"
	"tap_code_lai/common"
	"tap_code_lai/modules/restaurant_like/model"
)

func (s *sqlStore) GetRestaurantLike(ctx context.Context, ids []int) (map[int]int, error) {
	mapLike := make(map[int]int)
	db := s.db

	type sqlData struct {
		RestaurantId int `json:"restaurant_id" gorm:"column:restaurant_id"`
		LikeCount    int `json:"count" gorm:"column:count"`
	}

	var listLike []sqlData

	if err := db.Select("restaurant_id, count(restaurant_id) as count").
		Table(restaurantlikemodel.Like{}.TableName()).
		Group("restaurant_id").
		Where("restaurant_id in (?)", ids).
		Scan(&listLike).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range listLike {
		mapLike[item.RestaurantId] = int(item.LikeCount)
	}

	return mapLike, nil
}
