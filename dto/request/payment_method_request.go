package request

type CreatePaymentMethod struct {
	AccountName   string `form:"account_name"`
	AccountNumber string `form:"account_number"`
	BankName      string `form:"bank_name"`
	BankImages    string `form:"bank_images"`
}

type UpdatePaymentMethod struct {
	ID            int64  `form:"id"`
	AccountName   string `form:"account_name"`
	AccountNumber string `form:"account_number"`
	BankName      string `form:"bank_name"`
	BankImages    string `form:"bank_images"`
}
