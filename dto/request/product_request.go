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
	Name   string `json:"name"`
	Color  string `json:"color"`
	Images *multipart.FileHeader
	Sizes  string `json:"sizes"`
}

type CreateSizeVarianProduct struct {
	Size  string `json:"size"`
	Stock int64  `json:"stock"`
}
