package restaurantbiz

import (
	"errors"
	"golang.org/x/net/context"
	"tap_code_lai/modules/restaurant/resraurantmodel"
)

type DeleteRestaurant interface {
	FindByConditions(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*resraurantmodel.Restaurant, error)

	DeleteByCondition(ctx context.Context, conditions map[string]interface{}) error
}

type deleteRestaurantStore struct {
	store DeleteRestaurant
}

func NewDeleteRestaurantStor(store DeleteRestaurant) *deleteRestaurantStore {
	return &deleteRestaurantStore{store: store}
}

func (biz *deleteRestaurantStore) DeleteRestaurant(ctx context.Context, id int) error {
	oldData, err := biz.store.FindByConditions(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data deleted")
	}

	biz.store.DeleteByCondition(ctx, map[string]interface{}{"id": id})

	return nil
}
