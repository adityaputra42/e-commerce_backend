package request

import "mime/multipart"

type CreateProduct struct {
	CategoryID  int64                 `form:"category_id"`
	Name        string                `form:"name"`
	Description string                `form:"description"`
	Images      *multipart.FileHeader `form:"images"`
	Rating      float32               `form:"rating"`
	Price       float64               `form:"price"`
	ColorVarian string                `form:"color_varians"`
}

type CreateColorVarianProduct struct {
	ProductId int64                 `json:"product_id"`
	Name      string                `json:"name"`
	Color     string                `json:"color"`
	Images    *multipart.FileHeader `json:"-"`
	Sizes     string                `json:"sizes"`
}

type CreateSizeVarianProduct struct {
	ColorVarianId int64  `json:"color_varian_id"`
	Size          string `json:"size"`
	Stock         int64  `json:"stock"`
}

type UpdateSizeVarianProduct struct {
	ID    int64  `json:"id"`
	Size  string `json:"size"`
	Stock int64  `json:"stock"`
}
