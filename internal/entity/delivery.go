package entity

import (
	"encoding/json"
	"fmt"
)

type Delivery struct {
	Name    string `json:"name" validate:"required"`
	Phone   string `json:"phone" validate:"required,is_phone"`
	ZIP     string `json:"zip" validate:"required"`
	City    string `json:"city" validate:"required"`
	Address string `json:"address" validate:"required"`
	Region  string `json:"region" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
}

func (delivery *Delivery) Scan(v interface{}) error {
	switch vv := v.(type) {
	case []byte:
		return json.Unmarshal(vv, delivery)
	case string:
		return json.Unmarshal([]byte(vv), delivery)
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
}
