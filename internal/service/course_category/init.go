package coursecategory

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinnaserwan/coursphere-api/internal/http/request"
	"github.com/kevinnaserwan/coursphere-api/internal/http/response"
	"github.com/kevinnaserwan/coursphere-api/internal/model"
	courseCategoryRepository "github.com/kevinnaserwan/coursphere-api/internal/repository/course_category"
)

type courseCategoryService struct {
	CourseCategoryRepository courseCategoryRepository.Repository
}

func NewCourseCategoryService(courseCategoryRepository courseCategoryRepository.Repository) Service {
	return &courseCategoryService{
		CourseCategoryRepository: courseCategoryRepository,
	}
}

// CreateCourseCategory creates a new course category
func (s *courseCategoryService) CreateCourseCategory(ctx context.Context, req request.CreateCategoryCourseRequest) (err error) {
	courseCategory := &model.CategoryCourse{
		Name: req.Name,
	}

	if err := s.CourseCategoryRepository.Insert(ctx, courseCategory); err != nil {
		return err
	}

	return nil
}

// GetCourseCategoryByID retrieves a
func (s *courseCategoryService) GetCourseCategoryByID(ctx context.Context, id string) (res response.CourseCategoryResponse, err error) {
	courseCategory, err := s.CourseCategoryRepository.GetByID(ctx, uuid.MustParse(id))
	if err != nil {
		return res, err
	}

	res = response.CourseCategoryResponse{
		ID:   courseCategory.ID.String(),
		Name: courseCategory.Name,
	}

	return res, nil
}

// GetAllCourseCategories retrieves all course categories
func (s *courseCategoryService) GetAllCourseCategories(ctx context.Context) (res []response.CourseCategoryResponse, err error) {
	courseCategories, err := s.CourseCategoryRepository.GetAll(ctx)
	if err != nil {
		return res, err
	}

	for _, courseCategory := range courseCategories {
		res = append(res, response.CourseCategoryResponse{
			ID:   courseCategory.ID.String(),
			Name: courseCategory.Name,
		})
	}

	return res, nil
}

// UpdateCourseCategory updates a course category
func (s *courseCategoryService) UpdateCourseCategory(ctx context.Context, id string, req request.UpdateCategoryCourseRequest) (err error) {
	courseCategory, err := s.CourseCategoryRepository.GetByID(ctx, uuid.MustParse(id))
	if err != nil {
		return err
	}

	courseCategory.Name = req.Name

	if err := s.CourseCategoryRepository.Update(ctx, courseCategory); err != nil {
		return err
	}

	return nil
}

// DeleteCourseCategory deletes a course category
func (s *courseCategoryService) DeleteCourseCategory(ctx context.Context, id string) (err error) {
	if err := s.CourseCategoryRepository.Delete(ctx, uuid.MustParse(id)); err != nil {
		return err
	}

	return nil
}
