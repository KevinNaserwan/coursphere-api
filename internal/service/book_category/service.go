package bookcategory

import (
	"context"

	"github.com/kevinnaserwan/coursphere-api/internal/http/request"
	"github.com/kevinnaserwan/coursphere-api/internal/http/response"
)

type Service interface {
	Insert(ctx context.Context, req request.CreateBookCategoryRequest) (err error)
	GetByID(ctx context.Context, ID string) (bookcategory response.BookCategoryResponse, err error)
	GetAll(ctx context.Context) (bookcategorys []response.BookCategoryResponse, err error)
	Delete(ctx context.Context, ID string) (err error)
}
