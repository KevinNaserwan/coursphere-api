package mentor

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinnaserwan/coursphere-api/internal/model"
	MentorRepository "github.com/kevinnaserwan/coursphere-api/internal/repository/mentor"
	errCommon "github.com/kevinnaserwan/coursphere-api/internal/util/error"
	"gorm.io/gorm"
)

type mentorRepository struct {
	db *gorm.DB
}

func NewMentorRepository(db *gorm.DB) MentorRepository.Repository {
	return &mentorRepository{
		db: db,
	}
}

func (r *mentorRepository) CreateMentor(ctx context.Context, mentor *model.Mentor) error {
	if err := r.db.Create(mentor).Error; err != nil {
		return errCommon.NewBadRequest("Failed to insert mentor: " + err.Error())
	}

	return nil
}

func (r *mentorRepository) GetMentorByID(ctx context.Context, id uuid.UUID) (*model.Mentor, error) {
	var mentor model.Mentor
	if err := r.db.First(&mentor, "id = ?", id).Error; err != nil {
		return nil, errCommon.NewBadRequest("Mentor not found: " + err.Error())
	}

	return &mentor, nil
}

func (r *mentorRepository) GetAllMentors(ctx context.Context) ([]model.Mentor, error) {
	var mentors []model.Mentor
	if err := r.db.Find(&mentors).Error; err != nil {
		return nil, errCommon.NewBadRequest("Failed to retrieve mentors: " + err.Error())
	}

	return mentors, nil
}

func (r *mentorRepository) UpdateMentor(ctx context.Context, mentor *model.Mentor) error {
	if err := r.db.Save(mentor).Error; err != nil {
		return errCommon.NewBadRequest("Failed to update mentor: " + err.Error())
	}

	return nil
}

func (r *mentorRepository) DeleteMentor(ctx context.Context, id uuid.UUID) error {
	if err := r.db.Delete(&model.Mentor{}, "id = ?", id).Error; err != nil {
		return errCommon.NewBadRequest("Failed to delete mentor: " + err.Error())
	}

	return nil
}
