package helper

import (
	db "github.com/adityaputra42/e-commerce_backend/db/sqlc"
	"github.com/adityaputra42/e-commerce_backend/dto/response"
)

func ToAddressResponse(address db.Address) response.AddressResponse {
	return response.AddressResponse{
		ID:                   address.ID,
		RecipientName:        address.RecipientName,
		RecipientPhoneNumber: address.RecipientPhoneNumber,
		Province:             address.Province,
		City:                 address.City,
		District:             address.District,
		Village:              address.Village,
		PostalCode:           address.PostalCode,
		FullAddress:          address.FullAddress,
		CreatedAt:            address.CreatedAt,
		UpdatedAt:            address.UpdatedAt,
	}
}

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

func ToUserResponse(user db.User) response.UserResponse {

	return response.UserResponse{
		UID:       user.Uid,
		Username:  user.Username,
		FullName:  user.FullName,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
