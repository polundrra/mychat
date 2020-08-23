package model

import "time"

type User struct {
	ID        uint64
	Username  string
	CreatedAt time.Time
}
