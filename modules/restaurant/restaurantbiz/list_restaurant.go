package restaurantbiz

import (
	"errors"
	"golang.org/x/net/context"
	"tap_code_lai/common"
	"tap_code_lai/modules/restaurant/resraurantmodel"
)

type ListRestaurantRepo interface {
	ListRestaurant(ctx context.Context,
		conditions map[string]interface{},
		filter *resraurantmodel.Filter, paging *common.Paging) ([]resraurantmodel.Restaurant, error)
}

type listRestaurantBiz struct {
	repo ListRestaurantRepo
}

func NewListRestaurantBiz(repo ListRestaurantRepo) *listRestaurantBiz {
	return &listRestaurantBiz{repo}
}

func (biz *listRestaurantBiz) ListRestaurant(ctx context.Context,
	conditions map[string]interface{},
	filter *resraurantmodel.Filter, paging *common.Paging) ([]resraurantmodel.Restaurant, error) {

	if filter.CityId < 0 {
		return nil, errors.New("City_id must > 0")
	}

	result, err := biz.repo.ListRestaurant(ctx, conditions, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(resraurantmodel.EntityName, err)
	}

	return result, nil

}
