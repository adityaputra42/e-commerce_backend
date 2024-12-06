package response

import "time"

type Payment struct {
	ID           int64               `json:"id"`
	Transaction  TransactionResponse `json:"transaction"`
	TotalPayment float64             `json:"total_payment"`
	Status       string              `json:"status"`
	UpdatedAt    time.Time           `json:"updated_at"`
	CreatedAt    time.Time           `json:"created_at"`
}
