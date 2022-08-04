package rslikebiz

import (
	"golang.org/x/net/context"
	"tap_code_lai/common"
	restaurantlikemodel "tap_code_lai/modules/restaurant_like/model"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

type IncreaseLikeCountStore interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeRestaurantBiz struct {
	store    UserLikeRestaurantStore
	incStore IncreaseLikeCountStore
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore, incStore IncreaseLikeCountStore) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store, incStore}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(ctx context.Context,
	data *restaurantlikemodel.Like) error {

	err := biz.store.Create(ctx, data)
	if err != nil {
		return common.ErrCannotCreateEntity("Like", err)
	}

	// side effect

	go func() {
		_ = biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId)
	}()

	return nil
}
