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

type GetLikeRestaurantStore interface {
	GetRestaurantLike(ctx context.Context, ids []int) (map[int]int, error)
}

type listRestaurantStore struct {
	store ListRestaurantStore
	like  GetLikeRestaurantStore
}

func NewListRestaurantStore(store ListRestaurantStore, like GetLikeRestaurantStore) *listRestaurantStore {
	return &listRestaurantStore{store, like}
}

func (biz *listRestaurantStore) ListRestaurant(ctx context.Context,
	conditions map[string]interface{},
	filter *resraurantmodel.Filter, paging *common.Paging,
	moreKeys ...string) ([]resraurantmodel.Restaurant, error) {

	if filter.CityId < 0 {
		return nil, errors.New("City_id must > 0")
	}

	result, err := biz.store.ListByConditions(ctx, nil, filter, paging)
	if err != nil {
		return nil, err
	}

	ids := make([]int, len(result))

	for i := range result {
		ids[i] = result[i].Id
	}

	mapLike, err := biz.like.GetRestaurantLike(ctx, ids)

	if err != nil {
		return nil, common.ErrCannotListEntity(resraurantmodel.EntityName, err)
	}

	for i := range result {
		result[i].LikeCount = mapLike[result[i].Id]
	}

	return result, nil
}
