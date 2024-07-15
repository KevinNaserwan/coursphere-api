package request

type CreateBookCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}
