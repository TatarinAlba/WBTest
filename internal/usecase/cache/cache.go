package cache

import (
	"github.com/TatarinAlba/WBTest/internal/usecase"
	"github.com/TatarinAlba/WBTest/pkg/cache"
	"github.com/sirupsen/logrus"
	"time"
)

func NewCache(repository usecase.OrderRepository) (*cache.Cache, error) {
	// Can be changed by yourself
	createdCache := cache.New(time.Minute*5, time.Minute*10)
	orderGroup, err := repository.GetAllOrders()
	if err != nil {
		return nil, err
	}
	for _, order := range orderGroup {
		createdCache.Set(order.OrderUID, order, 0)
	}
	logrus.Info("Cache created!")
	return createdCache, err
}
