package book

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinnaserwan/coursphere-api/internal/model"
)

type Repository interface {
	Insert(ctx context.Context, book *model.Book) error
	GetByID(ctx context.Context, ID uuid.UUID) (*model.Book, error)
	GetAll(ctx context.Context, categoryName string) ([]model.Book, error)
	Update(ctx context.Context, book *model.Book) error
	Delete(ctx context.Context, ID uuid.UUID) error
}
