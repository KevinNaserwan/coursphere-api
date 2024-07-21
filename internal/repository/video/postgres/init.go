package video

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinnaserwan/coursphere-api/internal/model"
	VideoRepository "github.com/kevinnaserwan/coursphere-api/internal/repository/video"
	errCommon "github.com/kevinnaserwan/coursphere-api/internal/util/error"
	"gorm.io/gorm"
)

type videoRepository struct {
	db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) VideoRepository.Repository {
	return &videoRepository{
		db: db,
	}
}

// Insert inserts a new video record into the database.
func (r *videoRepository) Insert(ctx context.Context, video *model.Video) error {
	if err := r.db.Create(video).Error; err != nil {
		return errCommon.NewBadRequest("Failed to insert video: " + err.Error())
	}

	return nil
}

// GetByID retrieves a video by its ID.
func (r *videoRepository) GetByID(ctx context.Context, ID uuid.UUID) (*model.Video, error) {
	video := &model.Video{}
	if err := r.db.Where("id = ?", ID).First(video).Error; err != nil {
		return nil, errCommon.NewNotFound("Video not found: " + err.Error())
	}

	return video, nil
}

// Update updates a video record in the database.
func (r *videoRepository) Update(ctx context.Context, video *model.Video) error {
	if err := r.db.Save(video).Error; err != nil {
		return errCommon.NewBadRequest("Failed to update video: " + err.Error())
	}

	return nil
}

// Delete deletes a video record from the database.
func (r *videoRepository) Delete(ctx context.Context, ID uuid.UUID) error {
	if err := r.db.Where("id = ?", ID).Delete(&model.Video{}).Error; err != nil {
		return errCommon.NewBadRequest("Failed to delete video: " + err.Error())
	}

	return nil
}

// GetAll retrieves all video records from the database.
func (r *videoRepository) GetAll(ctx context.Context) ([]model.Video, error) {
	videos := []model.Video{}
	if err := r.db.Find(&videos).Error; err != nil {
		return nil, errCommon.NewNotFound("No videos found: " + err.Error())
	}

	return videos, nil
}
