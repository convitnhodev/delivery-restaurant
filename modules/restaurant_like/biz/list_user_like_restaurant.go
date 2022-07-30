package rslikebiz

import (
	"golang.org/x/net/context"
	"tap_code_lai/common"
	restaurantlikemodel "tap_code_lai/modules/restaurant_like/model"
)

type ListUserLikeRestaurantStore interface {
	ListUserLikeRestaurant(ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantlikemodel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]common.SimpleUser, error)
}

type listUserLikeRestaurantBiz struct {
	store ListUserLikeRestaurantStore
}

func NewListUserLikeRestaurantBiz(store ListUserLikeRestaurantStore) *listUserLikeRestaurantBiz {
	return &listUserLikeRestaurantBiz{store}
}

func (biz *listUserLikeRestaurantBiz) ListUsers(ctx context.Context,
	filter *restaurantlikemodel.Filter, paging *common.Paging) ([]common.SimpleUser, error) {

	users, err := biz.store.ListUserLikeRestaurant(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantlikemodel.EntityName, err)
	}

	return users, nil
}
