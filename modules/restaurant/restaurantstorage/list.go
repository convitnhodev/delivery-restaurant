package restaurantstorage

import (
	"golang.org/x/net/context"
	"tap_code_lai/common"
	"tap_code_lai/modules/restaurant/resraurantmodel"
)

func (s *sqlStore) ListByConditions(ctx context.Context,
	conditions map[string]interface{},
	filter *resraurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]resraurantmodel.Restaurant, error) {

	db := s.db

	var result []resraurantmodel.Restaurant

	if err := db.Table(resraurantmodel.Restaurant{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(resraurantmodel.Restaurant{}.TableName()).Where(conditions).Where("status in (1)")

	if v := filter; v != nil {
		if v.CityId > 0 {
			db = db.Where("city_id = ?", v.CityId)
		}
	}

	if paging.FakeCursor != "" {
		if uid, err := common.FromBase58(paging.FakeCursor); err == nil {
			db = db.Where("id > ?", uid.GetLocalID())
		}

	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil

}
