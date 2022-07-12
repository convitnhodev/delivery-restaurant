package restaurantstorage

import (
	"golang.org/x/net/context"
	"tap_code_lai/modules/restaurant/resraurantmodel"
)

func (s *sqlStore) ListByConditions(ctx context.Context,
	conditions map[string]interface{},
	filter *resraurantmodel.Filter,
	moreKeys ...string) ([]resraurantmodel.Restaurant, error) {

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(resraurantmodel.Restaurant{}.TableName()).Where(conditions).Where("status in (1)")

	if v := filter; v != nil {
		if v.CityId > 0 {
			db = db.Where("city_id = ?", v.CityId)
		}
	}

	var result []resraurantmodel.Restaurant

	if err := db.Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil

}
