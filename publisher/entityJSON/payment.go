package entityJSON

import "encoding/json"

type Payment struct {
	Transaction  string      `json:"transaction"`
	RequestID    string      `json:"request_id"`
	Currency     string      `json:"currency"`
	Provider     string      `json:"provider"`
	Amount       json.Number `json:"amount" fake:"{number:1,100000000}"`
	PaymentDt    json.Number `json:"payment_dt" fake:"{number:1,100000000}"`
	Bank         string      `json:"bank"`
	DeliveryCost json.Number `json:"delivery_cost" fake:"{number:1,100000000}"`
	GoodsTotal   json.Number `json:"goods_total" fake:"{number:1,100000000}"`
	CustomFee    json.Number `json:"custom_fee" fake:"{number:1,100000000}"`
}
