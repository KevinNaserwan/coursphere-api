package video

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinnaserwan/coursphere-api/internal/model"
)

type Repository interface {
	Insert(ctx context.Context, video *model.Video) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Video, error)
	GetAll(ctx context.Context) ([]model.Video, error)
	Update(ctx context.Context, video *model.Video) error
	Delete(ctx context.Context, id uuid.UUID) error
}
