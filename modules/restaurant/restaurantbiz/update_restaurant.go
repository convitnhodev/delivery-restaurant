package restaurantbiz

import (
	"errors"
	"golang.org/x/net/context"
	"tap_code_lai/modules/restaurant/resraurantmodel"
)

type UpdateRestaurantStore interface {
	FindByConditions(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*resraurantmodel.Restaurant, error)

	UpdateByCondition(ctx context.Context,
		data *resraurantmodel.RestaurantUpdate,
		conditions map[string]interface{}) error
}

type updateRestaurantStore struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantStor(store UpdateRestaurantStore) *updateRestaurantStore {
	return &updateRestaurantStore{store: store}
}

func (biz *updateRestaurantStore) UpdateRestaurant(ctx context.Context, id int, data *resraurantmodel.RestaurantUpdate) error {
	oldData, err := biz.store.FindByConditions(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data deleted")
	}

	biz.store.UpdateByCondition(ctx, data, map[string]interface{}{"id": id})

	return nil

}
