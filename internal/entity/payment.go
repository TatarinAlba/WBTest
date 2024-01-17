package entity

import (
	"encoding/json"
	"fmt"
)

type Payment struct {
	Transaction  string `json:"transaction" validate:"required"`
	RequestID    string `json:"request_id" validate:"required"`
	Currency     string `json:"currency" validate:"required"`
	Provider     string `json:"provider" validate:"required"`
	Amount       int64  `json:"amount" validate:"required,gte=0"`
	PaymentDt    int64  `json:"payment_dt" validate:"required,gte=0"`
	Bank         string `json:"bank" validate:"required"`
	DeliveryCost int64  `json:"delivery_cost" validate:"required,gte=0"`
	GoodsTotal   int64  `json:"goods_total" validate:"required,gte=0"`
	CustomFee    int64  `json:"custom_fee" validate:"required,gte=0"`
}

func (payment *Payment) Scan(v interface{}) error {
	switch vv := v.(type) {
	case []byte:
		return json.Unmarshal(vv, payment)
	case string:
		return json.Unmarshal([]byte(vv), payment)
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
}
