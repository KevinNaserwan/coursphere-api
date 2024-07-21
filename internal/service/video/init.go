package video

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinnaserwan/coursphere-api/internal/http/request"
	"github.com/kevinnaserwan/coursphere-api/internal/http/response"
	"github.com/kevinnaserwan/coursphere-api/internal/model"
	videoRepository "github.com/kevinnaserwan/coursphere-api/internal/repository/video"
)

type videoService struct {
	VideoRepository videoRepository.Repository
}

func NewVideoService(videoRepository videoRepository.Repository) Service {
	return &videoService{
		VideoRepository: videoRepository,
	}
}

// CreateVideo is a service that creates a new video.
func (s *videoService) CreateVideo(ctx context.Context, req request.VideoRequest) (err error) {
	video := &model.Video{
		Name: req.Name,
		URL:  req.URL,
		Time: req.Time,
	}

	if err := s.VideoRepository.Insert(ctx, video); err != nil {
		return err
	}

	return nil
}

// GetVideoByID is a service that retrieves a video by its ID.
func (s *videoService) GetVideoByID(ctx context.Context, id string) (res response.VideoResponse, err error) {
	video, err := s.VideoRepository.GetByID(ctx, uuid.MustParse(id))
	if err != nil {
		return res, err
	}

	res = response.VideoResponse{
		ID:   video.ID.String(),
		Name: video.Name,
		URL:  video.URL,
		Time: video.Time,
	}

	return res, nil
}

// GetAllVideos is a service that retrieves all videos.
func (s *videoService) GetAllVideos(ctx context.Context) (res []response.VideoResponse, err error) {
	videos, err := s.VideoRepository.GetAll(ctx)
	if err != nil {
		return res, err
	}

	for _, video := range videos {
		res = append(res, response.VideoResponse{
			ID:   video.ID.String(),
			Name: video.Name,
			URL:  video.URL,
			Time: video.Time,
		})
	}

	return res, nil
}

// UpdateVideo is a service that updates a video.
func (s *videoService) UpdateVideo(ctx context.Context, id string, req request.VideoRequest) (err error) {
	video, err := s.VideoRepository.GetByID(ctx, uuid.MustParse(id))
	if err != nil {
		return err
	}

	video.Name = req.Name
	video.URL = req.URL
	video.Time = req.Time

	if err := s.VideoRepository.Update(ctx, video); err != nil {
		return err
	}

	return nil
}

// DeleteVideo is a service that deletes a video.
func (s *videoService) DeleteVideo(ctx context.Context, id string) (err error) {
	if err := s.VideoRepository.Delete(ctx, uuid.MustParse(id)); err != nil {
		return err
	}

	return nil
}
