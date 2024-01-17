package usecase

import "github.com/TatarinAlba/WBTest/internal/entity"

type (
	OrderRepository interface {
		CreateOrder(order entity.Order) error
		GetOrderByUID(orderUid string) (entity.Order, error)
		GetAllOrders() ([]entity.Order, error)
		RemoveOrder(order entity.Order) error
		UpdateOrder(order entity.Order) error
	}
)
