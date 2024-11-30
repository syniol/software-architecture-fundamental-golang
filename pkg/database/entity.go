package database

import "time"

type Entity[T any] struct {
	ID               int
	Data             T
	CreatedTimestamp time.Time
}
