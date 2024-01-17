package entity

import "time"

type Order struct {
	OrderUID          string    `json:"order_uid" validate:"required" db:"order_uid"`
	TrackNumber       string    `json:"track_number" validate:"required" db:"track_number"`
	Entry             string    `json:"entry" validate:"required" db:"entry"`
	Delivery          Delivery  `json:"delivery" validate:"required" db:"delivery"`
	Payment           Payment   `json:"payment" validate:"required" db:"payment"`
	Items             Items     `json:"items" validate:"required" db:"items"`
	Locale            string    `json:"locale" validate:"required" db:"locale"`
	InternalSignature string    `json:"internal_signature" db:"internal_signature"`
	CustomerID        string    `json:"customer_id" validate:"required" db:"customer_id"`
	DeliveryService   string    `json:"delivery_service" validate:"required" db:"delivery_service"`
	ShardKey          string    `json:"shardKey" validate:"required" db:"shard_key"`
	SmID              int64     `json:"sm_id" validate:"required,gt=0" db:"sm_id"`
	DateCreated       time.Time `json:"date_created" validate:"required" db:"date_created"`
	OofShard          string    `json:"oof_shard" validate:"required" db:"oof_shard"`
}
