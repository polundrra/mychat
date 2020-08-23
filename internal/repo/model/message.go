package model

import "time"

type Message struct {
	ID uint64
	Chat uint64
	Author uint64
	Text string
	CreatedAt time.Time
}
