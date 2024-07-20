package course

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinnaserwan/coursphere-api/internal/model"
)

type Repository interface {
	Insert(ctx context.Context, course *model.Course) error
	GetByID(ctx context.Context, ID uuid.UUID) (*model.Course, error)
	GetAll(ctx context.Context) ([]model.Course, error)
	Update(ctx context.Context, course *model.Course) error
	Delete(ctx context.Context, ID uuid.UUID) error
}
