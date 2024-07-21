package coursecategory

import (
	"context"

	"github.com/kevinnaserwan/coursphere-api/internal/http/request"
	"github.com/kevinnaserwan/coursphere-api/internal/http/response"
)

type Service interface {
	CreateCourseCategory(ctx context.Context, req request.CreateCategoryCourseRequest) (err error)
	GetCourseCategoryByID(ctx context.Context, id string) (res response.CourseCategoryResponse, err error)
	GetAllCourseCategories(ctx context.Context) (res []response.CourseCategoryResponse, err error)
	UpdateCourseCategory(ctx context.Context, id string, req request.UpdateCategoryCourseRequest) (err error)
	DeleteCourseCategory(ctx context.Context, id string) (err error)
}
