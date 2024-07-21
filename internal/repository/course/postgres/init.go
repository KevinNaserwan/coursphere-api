package course

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinnaserwan/coursphere-api/internal/model"
	CourseRepository "github.com/kevinnaserwan/coursphere-api/internal/repository/course"
	errCommon "github.com/kevinnaserwan/coursphere-api/internal/util/error"
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
		return errCommon.NewBadRequest("Failed to insert course: " + err.Error())
	}
	return nil
}

// GetByID retrieves a course by its ID including Mentor and CategoryCourse
func (r *courseRepository) GetByID(ctx context.Context, ID uuid.UUID) (*model.Course, error) {
	course := &model.Course{}
	if err := r.db.Preload("Mentor").Preload("Category").Where("id = ?", ID).First(course).Error; err != nil {
		return nil, errCommon.NewNotFound("Course not found: " + err.Error())
	}
	return course, nil
}

// Update updates a course in the database
func (r *courseRepository) Update(ctx context.Context, course *model.Course) error {
	if err := r.db.Save(course).Error; err != nil {
		return errCommon.NewBadRequest("Failed to update course: " + err.Error())
	}
	return nil
}

// Delete deletes a course from the database
func (r *courseRepository) Delete(ctx context.Context, ID uuid.UUID) error {
	if err := r.db.Where("id = ?", ID).Delete(&model.Course{}).Error; err != nil {
		return errCommon.NewNotFound("Failed to delete course: " + err.Error())
	}
	return nil
}

func (r *courseRepository) GetAll(ctx context.Context) ([]model.Course, error) {
	courses := []model.Course{}
	query := r.db.WithContext(ctx).
		Preload("Mentor").
		Preload("Category").
		Preload("Videos") // Preload videos

	if err := query.Find(&courses).Error; err != nil {
		return nil, errCommon.NewBadRequest("No courses found: " + err.Error())
	}
	return courses, nil
}

func (r *courseRepository) InsertCourseVideo(ctx context.Context, courseID uuid.UUID, videoID uuid.UUID) error {
	if err := r.db.WithContext(ctx).Exec("INSERT INTO video (course_id, video_id) VALUES (?, ?)", courseID, videoID).Error; err != nil {
		return errCommon.NewBadRequest("Failed to insert course video: " + err.Error())
	}
	return nil
}
