package course

import (
	"context"

	"github.com/kevinnaserwan/coursphere-api/internal/http/request"
	"github.com/kevinnaserwan/coursphere-api/internal/http/response"
)

type Service interface {
	CreateCourse(ctx context.Context, req request.CreateCourseRequest) (err error)
	GetCourseByID(ctx context.Context, id string) (res response.CourseResponse, err error)
	GetAllCourses(ctx context.Context) (res []response.CourseResponse, err error)
	UpdateCourse(ctx context.Context, id string, req request.UpdateCourseRequest) (err error)
	DeleteCourse(ctx context.Context, id string) (err error)
	AddCourseVideo(ctx context.Context, courseID string, videoID string) error
}
