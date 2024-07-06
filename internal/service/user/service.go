package user

import (
	"context"

	"github.com/kevinnaserwan/coursphere-api/internal/http/request"
	"github.com/kevinnaserwan/coursphere-api/internal/http/response"
)

type Service interface {
	Register(ctx context.Context, req request.UserRegisterRequest) (err error)
	Login(ctx context.Context, req request.UserLoginRequest) (JwtToken string, err error)
	GetByID(ctx context.Context, ID string) (user response.UserResponse, err error)
	GetAll(ctx context.Context) (user []response.UserResponse, err error)
	Update(ctx context.Context, ID string, req request.UserUpdateRequest) (err error)
	Delete(ctx context.Context, ID string) (err error)
}
