package entityJSON

import "encoding/json"

type Item struct {
	ChartID     json.Number `json:"chrt_id" fake:"{number:1,100000000}"`
	TrackNumber string      `json:"track_number"`
	Price       json.Number `json:"price" fake:"{number:1,1000}"`
	Rid         string      `json:"rid"`
	Name        string      `json:"name"`
	Sale        json.Number `json:"sale" fake:"{number:1,1000}"`
	Size        string      `json:"size"`
	TotalPrice  json.Number `json:"total_price" fake:"{number:1,1000}"`
	NmId        json.Number `json:"nm_id" fake:"{number:1,1000}"`
	Brand       string      `json:"brand"`
	Status      json.Number `json:"status" fake:"{number:1,1000}"`
}
