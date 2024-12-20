package request

type CategoriesParam struct {
	Name string `form:"name"`
	Icon string `form:"icon"`
}

type UpdateCategory struct {
	ID   int64  `form:"id"`
	Name string `form:"name"`
	Icon string `form:"icon"`
}
