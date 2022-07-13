package restaurantbiz

import (
	"golang.org/x/net/context"
	"tap_code_lai/common"
	"tap_code_lai/modules/restaurant/resraurantmodel"
)

type FindRestaurantStore interface {
	FindByConditions(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*resraurantmodel.Restaurant, error)
}

type findRestaurantStore struct {
	store FindRestaurantStore
}

func NewFindRestaurantStore(store FindRestaurantStore) *findRestaurantStore {
	return &findRestaurantStore{store}
}

func (biz *findRestaurantStore) FindRestaurant(ctx context.Context, conditions map[string]interface{}) (*resraurantmodel.Restaurant, error) {
	data, err := biz.store.FindByConditions(ctx, conditions)
	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(resraurantmodel.EntityName, err)
		}
		return nil, common.ErrCannotGetEntity(resraurantmodel.EntityName, err)
	}

	if data.Status != 1 {
		return nil, common.ErrEntityDeleted(resraurantmodel.EntityName, nil)
	}
	return data, nil

}
