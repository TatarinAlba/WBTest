package cache

import "time"

type Item struct {
	Value      any
	Created    time.Time
	Expiration int64
}
