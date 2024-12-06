package request

type CreatePayment struct {
	TransactionID string  `json:"transaction_id"`
	TotalPayment  float64 `json:"total_payment"`
}
