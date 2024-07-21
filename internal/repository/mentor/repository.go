package mentor

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinnaserwan/coursphere-api/internal/model"
)

type Repository interface {
	CreateMentor(ctx context.Context, mentor *model.Mentor) error
	GetMentorByID(ctx context.Context, id uuid.UUID) (*model.Mentor, error)
	GetAllMentors(ctx context.Context) ([]model.Mentor, error)
	UpdateMentor(ctx context.Context, mentor *model.Mentor) error
	DeleteMentor(ctx context.Context, id uuid.UUID) error
}
