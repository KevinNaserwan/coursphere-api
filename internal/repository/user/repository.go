package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinnaserwan/coursphere-api/internal/model"
)

type Repository interface {
	Insert(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, ID uuid.UUID) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetAll(ctx context.Context) ([]model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, ID uuid.UUID) error
}
