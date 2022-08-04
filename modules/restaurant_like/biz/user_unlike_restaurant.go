package rslikebiz

import (
	"golang.org/x/net/context"
	"tap_code_lai/common"
)

type UserUnLikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restaurantId int) error
}

type userUnLikeRestaurantBiz struct {
	store UserUnLikeRestaurantStore
}

func NewUserUnLikeRestaurantBiz(store UserUnLikeRestaurantStore) *userUnLikeRestaurantBiz {
	return &userUnLikeRestaurantBiz{store}
}

func (biz *userUnLikeRestaurantBiz) UnLikeRestaurant(ctx context.Context,
	userId, restaurantId int) error {

	err := biz.store.Delete(ctx, userId, restaurantId)
	if err != nil {
		return common.ErrCannotCreateEntity("UnLike", err)
	}

	return nil
}
