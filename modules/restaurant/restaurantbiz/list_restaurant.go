package restaurantbiz

import (
	"errors"
	"golang.org/x/net/context"
	"tap_code_lai/common"
	"tap_code_lai/modules/restaurant/resraurantmodel"
)

type ListRestaurantStore interface {
	ListByConditions(ctx context.Context,
		conditions map[string]interface{},
		filter *resraurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]resraurantmodel.Restaurant, error)
}

type listRestaurantStore struct {
	store ListRestaurantStore
}

func NewListRestaurantStore(store ListRestaurantStore) *listRestaurantStore {
	return &listRestaurantStore{store: store}
}

func (biz *listRestaurantStore) ListRestaurant(ctx context.Context,
	conditions map[string]interface{},
	filter *resraurantmodel.Filter, paging *common.Paging,
	moreKeys ...string) ([]resraurantmodel.Restaurant, error) {

	if filter.CityId < 0 {
		return nil, errors.New("City_id must >= 0")
	}

	result, err := biz.store.ListByConditions(ctx, nil, filter, paging)
	if err != nil {
		return nil, err
	}

	return result, nil

}
