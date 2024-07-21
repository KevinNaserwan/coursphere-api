package coursecategory

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinnaserwan/coursphere-api/internal/model"
)

type Repository interface {
	Insert(ctx context.Context, courseCategory *model.CategoryCourse) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.CategoryCourse, error)
	GetAll(ctx context.Context) ([]model.CategoryCourse, error)
	Update(ctx context.Context, courseCategory *model.CategoryCourse) error
	Delete(ctx context.Context, id uuid.UUID) error
}
