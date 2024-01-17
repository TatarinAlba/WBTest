package entity

import (
	"encoding/json"
	"fmt"
)

type Items []Item

type Item struct {
	ChartID     int64   `json:"chrt_id" validate:"required,gt=0"`
	TrackNumber string  `json:"track_number" validate:"required"`
	Price       float64 `json:"price" validate:"required,gte=0"`
	Rid         string  `json:"rid" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Sale        float64 `json:"sale" validate:"required,gte=0"`
	Size        string  `json:"size" validate:"required"`
	TotalPrice  float64 `json:"total_price" validate:"required,gte=0"`
	NmId        int64   `json:"nm_id" validate:"required,gt=0"`
	Brand       string  `json:"brand" validate:"required"`
	Status      int64   `json:"status" validate:"required,gt=0"`
}

func (item *Items) Scan(v interface{}) error {
	switch vv := v.(type) {
	case []byte:
		return json.Unmarshal(vv, item)
	case string:
		return json.Unmarshal([]byte(vv), item)
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
}
