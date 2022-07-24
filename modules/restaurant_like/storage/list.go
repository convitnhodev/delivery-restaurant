package storage

import (
	 "fmt"
	 "github.com/btcsuite/btcutil/base58"
	"golang.org/x/net/context"
	"tap_code_lai/common"
	"tap_code_lai/modules/restaurant/resraurantmodel"
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

func (s *sqlStore) GetUserLikeRestaurant(ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]common.SimpleUser, error) {

	db := s.db

	var result []common.SimpleUser

	db = db.Table(resraurantmodel.Restaurant{}.TableName()).Where(conditions)

	if v := filter; v != nil {
		if v.RestaurantId > 0 {
			db = db.Where("restaurant_id = ?", v.RestaurantId)
		}
	}

	if err := db.Table(resraurantmodel.Restaurant{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if paging.FakeCursor != "" {
		if uid, err := common.FromBase58(paging.FakeCursor); err == nil {
			db = db.Where("created_at < ?", uid.GetLocalID())
		}
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("created_at desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i, items := range result {
		if i == len(result)-1 {
			cursorStr := base58.Encode([]byte(fmt.Sprint("%v", items.CreatedAt)))"))
		}
	}

	return result, nil
}
