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

func ToCategoryRespone(c db.Category) response.Category {
	return response.Category{
		ID:        c.ID,
		Name:      c.Name,
		Icon:      c.Icon,
		UpdatedAt: c.UpdatedAt,
		CreatedAt: c.CreatedAt,
	}
}

func ToPaymentMethodRespone(pm db.PaymentMethod) response.PaymentMethodResponse {
	return response.PaymentMethodResponse{
		ID:            pm.ID,
		AccountName:   pm.AccountName,
		AccountNumber: pm.AccountNumber,
		BankName:      pm.BankName,
		BankImages:    pm.BankImages,
		CreatedAt:     pm.CreatedAt,
		UpdatedAt:     pm.UpdatedAt,
	}

}

func ToShippingRespone(sh db.Shipping) response.ShippingResponse {
	return response.ShippingResponse{
		ID:        sh.ID,
		Name:      sh.Name,
		Price:     sh.Price,
		State:     sh.State,
		CreatedAt: sh.CreatedAt,
		UpdatedAt: sh.UpdatedAt,
	}

}

func ToSizeVarianResponse(size db.SizeVarian) response.SizeVarianResponse {
	return response.SizeVarianResponse{
		ID:            size.ID,
		ColorVarianID: size.ColorVarianID,
		Size:          size.Size,
		Stock:         size.Stock,
		CreatedAt:     size.CreatedAt,
		UpdatedAt:     size.UpdatedAt,
	}

}
func ToColorVarianResponse(cv db.ColorVarian, listSize []response.SizeVarianResponse) response.ColorVarianResponse {
	return response.ColorVarianResponse{
		ID:         cv.ID,
		ProductID:  cv.ProductID,
		Name:       cv.Name,
		Color:      cv.Color,
		Images:     cv.Images,
		SizeVarian: listSize,
		CreatedAt:  cv.CreatedAt,
		UpdatedAt:  cv.UpdatedAt,
	}
}

func ToProductDetailResponse(p db.Product, category response.Category, listVarian []response.ColorVarianResponse) response.ProductDetailResponse {
	return response.ProductDetailResponse{
		ID:          p.ID,
		Category:    category,
		Name:        p.Name,
		Description: p.Description,
		Images:      p.Images,
		Rating:      p.Rating,
		Price:       p.Price,
		ColorVarian: listVarian,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}

}

func ToProductResponse(p db.Product, category response.Category) response.ProductResponse {
	return response.ProductResponse{
		ID:          p.ID,
		Category:    category,
		Name:        p.Name,
		Description: p.Description,
		Images:      p.Images,
		Rating:      p.Rating,
		Price:       p.Price,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}

}
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
