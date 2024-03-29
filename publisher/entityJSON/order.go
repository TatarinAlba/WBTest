package entityJSON

import (
	"encoding/json"
	"time"
)

type Order struct {
	OrderUID          string      `json:"order_uid"`
	TrackNumber       string      `json:"track_number"`
	Entry             string      `json:"entry"`
	Delivery          Delivery    `json:"delivery"`
	Payment           Payment     `json:"payment"`
	Items             []Item      `json:"items"`
	Locale            string      `json:"locale"`
	InternalSignature string      `json:"internal_signature"`
	CustomerID        string      `json:"customer_id"`
	DeliveryService   string      `json:"delivery_service"`
	ShardKey          string      `json:"shardKey"`
	SmID              json.Number `json:"sm_id" fake:"{number:1,100000000}"`
	DateCreated       time.Time   `json:"date_created"`
	OofShard          string      `json:"oof_shard"`
}
