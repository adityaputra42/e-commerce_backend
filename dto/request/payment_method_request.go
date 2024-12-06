package request

type CreatePaymentMethod struct {
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
	BankName      string `json:"bank_name"`
	BankImages    string `json:"bank_images"`
}
