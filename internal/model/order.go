package model

import "time"

type Order struct {
	ID           int64     `json:"id"`
	UserID       int64     `json:"user_id"`
	ExtendedCost float64   `json:"extended_cost"`
	CreatedAt    time.Time `json:"created_at"`
}

type Orders struct {
	Orders []Order `json:"orders"`
}
