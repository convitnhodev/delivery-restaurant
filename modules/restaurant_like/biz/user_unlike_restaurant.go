package rslikebiz

import (
	"golang.org/x/net/context"
	"tap_code_lai/common"
)

type UserUnLikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restaurantId int) error
}

type DecreaseLikeCountStore interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type userUnLikeRestaurantBiz struct {
	store    UserUnLikeRestaurantStore
	decStore DecreaseLikeCountStore
}

func NewUserUnLikeRestaurantBiz(store UserUnLikeRestaurantStore, decStore DecreaseLikeCountStore) *userUnLikeRestaurantBiz {
	return &userUnLikeRestaurantBiz{store, decStore}
}

func (biz *userUnLikeRestaurantBiz) UnLikeRestaurant(ctx context.Context,
	userId, restaurantId int) error {

	err := biz.store.Delete(ctx, userId, restaurantId)
	if err != nil {
		return common.ErrCannotCreateEntity("UnLike", err)
	}

	_ = biz.decStore.DecreaseLikeCount(ctx, restaurantId)
	return nil
}
