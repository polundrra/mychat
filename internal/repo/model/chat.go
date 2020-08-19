package model

import (
	"time"
)

type Chat struct {
	ID uint64
	Name string
	Users []uint32
	CreatedAt time.Time
}
