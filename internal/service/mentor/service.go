package mentor

import (
	"context"

	"github.com/kevinnaserwan/coursphere-api/internal/http/request"
	"github.com/kevinnaserwan/coursphere-api/internal/http/response"
)

type Service interface {
	CreateMentor(ctx context.Context, req request.CreateMentorRequest) (err error)
	GetMentorByID(ctx context.Context, id string) (res response.MentorResponse, err error)
	GetAllMentors(ctx context.Context) (res []response.MentorResponse, err error)
	UpdateMentor(ctx context.Context, id string, req request.UpdateMentorRequest) (err error)
	DeleteMentor(ctx context.Context, id string) (err error)
}
