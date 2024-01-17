package repo

import (
	"encoding/json"
	"github.com/TatarinAlba/WBTest/internal/entity"
	"github.com/jmoiron/sqlx"
)

type OrderRepositoryPostgres struct {
	connection *sqlx.DB
}

func (repository *OrderRepositoryPostgres) CreateOrder(order entity.Order) error {
	deliveryBin, err := json.Marshal(order.Delivery)
	if err != nil {
		return err
	}
	paymentBin, err := json.Marshal(order.Payment)
	if err != nil {
		return err
	}
	itemsBin, err := json.Marshal(order.Items)
	if err != nil {
		return err
	}
	_, err = repository.connection.Exec(
		"INSERT INTO orders VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)",
		order.OrderUID, order.TrackNumber, order.Entry, deliveryBin, paymentBin, itemsBin, order.Locale,
		order.InternalSignature, order.CustomerID, order.DeliveryService, order.ShardKey, order.SmID, order.DateCreated,
		order.OofShard)
	if err != nil {
		return err
	}
	return nil
}

func (repository *OrderRepositoryPostgres) GetOrderByUID(orderUid string) (entity.Order, error) {
	gotOrder := entity.Order{}
	err := repository.connection.Get(&gotOrder, "SELECT * FROM orders WHERE order_uid=$1", orderUid)
	if err != nil {
		return gotOrder, err
	}
	return gotOrder, nil
}

func (repository *OrderRepositoryPostgres) RemoveOrder(order entity.Order) error {
	_, err := repository.connection.Exec("DELETE FROM orders WHERE order_uid=$1", order.OrderUID)
	if err != nil {
		return err
	}
	return nil
}

func (repository *OrderRepositoryPostgres) UpdateOrder(order entity.Order) error {
	deliveryBin, err := json.Marshal(order.Delivery)
	if err != nil {
		return err
	}
	paymentBin, err := json.Marshal(order.Payment)
	if err != nil {
		return err
	}
	itemsBin, err := json.Marshal(order.Items)
	if err != nil {
		return err
	}
	_, err = repository.connection.Exec(
		`UPDATE orders SET track_number=$1, entry=$2, delivery=$3, payment=$4, items=$5, locale=$6, 
                  internal_signature=$7, customer_id=$8, delivery_service=$9, shard_key=$10, sm_id=$11, 
                  date_created=$12, oof_shard=$13 WHERE order_uid=$14`,
		order.TrackNumber, order.Entry, deliveryBin, paymentBin, itemsBin, order.Locale,
		order.InternalSignature, order.CustomerID, order.DeliveryService, order.ShardKey, order.SmID, order.DateCreated,
		order.OofShard, order.OrderUID)
	if err != nil {
		return err
	}
	return nil
}

func (repository *OrderRepositoryPostgres) GetAllOrders() ([]entity.Order, error) {
	var orders []entity.Order
	err := repository.connection.Select(&orders, "SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func NewOrderRepositoryPostgres(database *sqlx.DB) *OrderRepositoryPostgres {
	return &OrderRepositoryPostgres{database}
}
