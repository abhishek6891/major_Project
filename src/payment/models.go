package payment

import "time"

type Payment struct {
	ID          string    `json:"id"`
	ClientID    string    `json:"client_id"`
	DeveloperID string    `json:"developer_id"`
	Amount      float64   `json:"amount"`
	Status      string    `json:"status"` // pending, success, failed
	Timestamp   time.Time `json:"timestamp"`
}
