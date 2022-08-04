package rslikebiz

import (
	"golang.org/x/net/context"
	"tap_code_lai/common"
	restaurantlikemodel "tap_code_lai/modules/restaurant_like/model"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

type userLikeRestaurantBiz struct {
	store UserLikeRestaurantStore
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(ctx context.Context,
	data *restaurantlikemodel.Like) error {

	err := biz.store.Create(ctx, data)
	if err != nil {
		return common.ErrCannotCreateEntity("Like", err)
	}

	return nil
}
