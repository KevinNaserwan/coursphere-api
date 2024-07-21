package coursecategory

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinnaserwan/coursphere-api/internal/model"
	CourseCategoryRepository "github.com/kevinnaserwan/coursphere-api/internal/repository/course_category"
	errCommon "github.com/kevinnaserwan/coursphere-api/internal/util/error"
	"gorm.io/gorm"
)

type courseCategoryRepository struct {
	db *gorm.DB
}

func NewCourseCategoryRepository(db *gorm.DB) CourseCategoryRepository.Repository {
	return &courseCategoryRepository{
		db: db,
	}
}

// Insert inserts a new course category into the database
func (r *courseCategoryRepository) Insert(ctx context.Context, courseCategory *model.CategoryCourse) error {
	if err := r.db.Create(courseCategory).Error; err != nil {
		return errCommon.NewBadRequest("Failed to insert course category: " + err.Error())
	}
	return nil
}

// GetByID retrieves a course category by its ID
func (r *courseCategoryRepository) GetByID(ctx context.Context, ID uuid.UUID) (*model.CategoryCourse, error) {
	courseCategory := &model.CategoryCourse{}
	if err := r.db.Where("id = ?", ID).First(courseCategory).Error; err != nil {
		return nil, errCommon.NewNotFound("User not found: " + err.Error())
	}
	return courseCategory, nil
}

// Update updates a course category in the database
func (r *courseCategoryRepository) Update(ctx context.Context, courseCategory *model.CategoryCourse) error {
	if err := r.db.Save(courseCategory).Error; err != nil {
		return errCommon.NewBadRequest("Failed to update course category: " + err.Error())
	}
	return nil
}

// Delete deletes a course category from the database
func (r *courseCategoryRepository) Delete(ctx context.Context, ID uuid.UUID) error {
	if err := r.db.Where("id = ?", ID).Delete(&model.CategoryCourse{}).Error; err != nil {
		return errCommon.NewBadRequest("Failed to delete course category: " + err.Error())
	}
	return nil
}

// GetAll retrieves all course categories from the database
func (r *courseCategoryRepository) GetAll(ctx context.Context) ([]model.CategoryCourse, error) {
	courseCategories := []model.CategoryCourse{}
	if err := r.db.Find(&courseCategories).Error; err != nil {
		return nil, errCommon.NewNotFound("No course categories found: " + err.Error())
	}
	return courseCategories, nil
}
