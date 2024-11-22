package entities

import "time"

type History struct {
	ID         int       `json:"id"`
	CustomerID int       `json:"customer_id"`
	Action     string    `json:"action"`
	Amount     float64   `json:"amount"`
	Timestamp  time.Time `json:"timestamp"`
}
