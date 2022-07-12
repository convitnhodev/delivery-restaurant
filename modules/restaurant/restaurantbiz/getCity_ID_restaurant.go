package restaurantbiz

import (
	"golang.org/x/net/context"
	"tap_code_lai/modules/restaurant/resraurantmodel"
)

type Find_city_RestaurantStore interface {
	FindByConditions(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*resraurantmodel.Restaurant, error)
}

type find_city_RestaurantStore struct {
	store FindRestaurantStore
}

func NewFind_city_RestaurantStore(store FindRestaurantStore) *find_city_RestaurantStore {
	return &find_city_RestaurantStore{store}
}

func (biz *find_city_RestaurantStore) Find_city_Restaurant(ctx context.Context, conditions map[string]interface{}) (*resraurantmodel.Restaurant, error) {
	data, err := biz.store.FindByConditions(ctx, conditions)
	if err != nil {
		return data, err
	}
	return data, nil
}
