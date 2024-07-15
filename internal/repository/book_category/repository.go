package bookcategory

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinnaserwan/coursphere-api/internal/model"
)

type Repository interface {
	Insert(ctx context.Context, bookCategory *model.CategoryBook) error
	GetByID(ctx context.Context, ID uuid.UUID) (*model.CategoryBook, error)
	GetAll(ctx context.Context) ([]model.CategoryBook, error)
	Update(ctx context.Context, bookCategory *model.CategoryBook) error
	Delete(ctx context.Context, ID uuid.UUID) error
}
