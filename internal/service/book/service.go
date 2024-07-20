package book

import (
	"context"

	"github.com/kevinnaserwan/coursphere-api/internal/http/request"
	"github.com/kevinnaserwan/coursphere-api/internal/http/response"
)

type Service interface {
	CreateBook(ctx context.Context, req request.CreateBookRequest) (res response.BookResponse, err error)
	GetBookByID(ctx context.Context, ID string) (res response.BookResponse, err error)
	GetAllBooks(ctx context.Context, categoryName string) ([]response.BookResponse, error)
	UpdateBook(ctx context.Context, ID string, req request.UpdateBookRequest) (res response.BookResponse, err error)
	DeleteBook(ctx context.Context, ID string) (err error)
}
