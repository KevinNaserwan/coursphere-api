package video

import (
	"context"

	"github.com/kevinnaserwan/coursphere-api/internal/http/request"
	"github.com/kevinnaserwan/coursphere-api/internal/http/response"
)

type Service interface {
	CreateVideo(ctx context.Context, req request.VideoRequest) (err error)
	GetVideoByID(ctx context.Context, id string) (res response.VideoResponse, err error)
	GetAllVideos(ctx context.Context) (res []response.VideoResponse, err error)
	UpdateVideo(ctx context.Context, id string, req request.VideoRequest) (err error)
	DeleteVideo(ctx context.Context, id string) (err error)
}
