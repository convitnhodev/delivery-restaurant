package restaurantbiz

import (
	"golang.org/x/net/context"
	"tap_code_lai/modules/restaurant/resraurantmodel"
)

type CreateRestaurantStore interface {
	Create(ctx context.Context, data *resraurantmodel.RestaurantCreate) error
}

type createRestaurantStore struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantStore(store CreateRestaurantStore) *createRestaurantStore {
	return &createRestaurantStore{store}
}

func (biz *createRestaurantStore) CreateRestaurant(ctx context.Context, data *resraurantmodel.RestaurantCreate) error {
	if err := data.Validata(); err != nil {
		return err
	}
	err := biz.store.Create(ctx, data)
	return err
}
