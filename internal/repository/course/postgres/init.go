package course

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinnaserwan/coursphere-api/internal/model"
	CourseRepository "github.com/kevinnaserwan/coursphere-api/internal/repository/course"
	"gorm.io/gorm"
)

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository.Repository {
	return &courseRepository{
		db: db,
	}
}

// Insert inserts a new course into the database
func (r *courseRepository) Insert(ctx context.Context, course *model.Course) error {
	if err := r.db.Create(course).Error; err != nil {
		return err
	}
	return nil
}

// GetByID retrieves a course by its ID
func (r *courseRepository) GetByID(ctx context.Context, ID uuid.UUID) (*model.Course, error) {
	course := &model.Course{}
	if err := r.db.Where("id = ?", ID).First(course).Error; err != nil {
		return nil, err
	}
	return course, nil
}

// Update updates a course in the database
func (r *courseRepository) Update(ctx context.Context, course *model.Course) error {
	if err := r.db.Save(course).Error; err != nil {
		return err
	}
	return nil
}

// Delete deletes a course from the database
func (r *courseRepository) Delete(ctx context.Context, ID uuid.UUID) error {
	if err := r.db.Where("id = ?", ID).Delete(&model.Course{}).Error; err != nil {
		return err
	}
	return nil
}

// GetAll retrieves all courses from the database
func (r *courseRepository) GetAll(ctx context.Context) ([]model.Course, error) {
	courses := []model.Course{}
	if err := r.db.Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}
