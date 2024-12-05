package helper

// func ToAccountResponse(account domain.Account) response.AccountResponse {
// 	return response.AccountResponse{
// 		ID:        account.ID,
// 		UserId:    account.UserId,
// 		Balance:   account.Balance,
// 		Currency:  account.Currency,
// 		CreatedAt: account.CreatedAt,
// 		UpdatedAt: account.UpdatedAt,
// 	}
// }

// func ToTranferRespone(transfer domain.Transaction, from domain.Account, to domain.Account) response.TransferResponse {
// 	return response.TransferResponse{
// 		TransactionID: transfer.ID,
// 		From:          ToAccountResponse(from),
// 		To:            ToAccountResponse(to),
// 		Amount:        transfer.Amount,
// 		Currency:      transfer.Currency,
// 		CreatedAt:     transfer.CreatedAt,
// 	}

// }

// func ToDepositRespone(deposit domain.Deposit, account domain.Account) response.DepositResponse {
// 	return response.DepositResponse{
// 		ID:        deposit.ID,
// 		Amount:    deposit.Amount,
// 		Account:   ToAccountResponse(account),
// 		CreatedAt: deposit.CreatedAt,
// 	}

// }

// func ToUserResponse(user domain.User) response.UserResponse {
// 	var listAccount []response.AccountResponse

// 	for _, account := range user.Accounts {
// 		listAccount = append(listAccount, ToAccountResponse(account))
// 	}
// 	return response.UserResponse{
// 		ID:       user.ID,
// 		Username: user.Username,
// 		FullName: user.FullName,
// 		Email:    user.Email,
// 		Role:     user.Role,
// 		Accounts: listAccount,
// 	}
// }
