package usecase

import (
	"github.com/TatarinAlba/WBTest/internal/entity"
	"github.com/TatarinAlba/WBTest/pkg/cache"
	"github.com/sirupsen/logrus"
)

type OrderUsecase struct {
	repository OrderRepository
	cache      *cache.Cache
}

func NewOrderUsecase(repository OrderRepository, ordersCache *cache.Cache) *OrderUsecase {
	return &OrderUsecase{repository, ordersCache}
}

func (usecase *OrderUsecase) CreateOrder(order entity.Order) error {
	err := usecase.repository.CreateOrder(order)
	if err != nil {
		return err
	}
	usecase.cache.Set(order.OrderUID, order, 0)
	logrus.Debugf("Added to cache [%s]", order.OrderUID)
	return nil
}

func (usecase *OrderUsecase) GetOrder(orderUid string) (entity.Order, error) {
	if order, found := usecase.cache.Get(orderUid); found {
		logrus.Debugf("Cache HIT! [%s]", orderUid)
		return order.(entity.Order), nil
	}
	order, err := usecase.repository.GetOrderByUID(orderUid)
	if err != nil {
		return order, err
	}
	usecase.cache.Set(order.OrderUID, order, 0)
	logrus.Debugf("Cache MISS! [%s]", order.OrderUID)
	return order, err
}
