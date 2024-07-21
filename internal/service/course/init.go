package course

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinnaserwan/coursphere-api/internal/http/request"
	"github.com/kevinnaserwan/coursphere-api/internal/http/response"
	"github.com/kevinnaserwan/coursphere-api/internal/model"
	courseRepository "github.com/kevinnaserwan/coursphere-api/internal/repository/course"
	courseCategoryRepository "github.com/kevinnaserwan/coursphere-api/internal/repository/course_category"
	mentorRepository "github.com/kevinnaserwan/coursphere-api/internal/repository/mentor"
)

type courseService struct {
	CourseRepository         courseRepository.Repository
	MentorRepository         mentorRepository.Repository
	CourseCategoryRepository courseCategoryRepository.Repository
}

func NewCourseService(courseRepository courseRepository.Repository, mentorRepository mentorRepository.Repository, courseCategoryRepository courseCategoryRepository.Repository) Service {
	return &courseService{
		CourseRepository:         courseRepository,
		MentorRepository:         mentorRepository,
		CourseCategoryRepository: courseCategoryRepository,
	}
}

// CreateCourse creates a new course.
func (s *courseService) CreateCourse(ctx context.Context, req request.CreateCourseRequest) (err error) {
	course := &model.Course{
		BannerImage: req.BannerImage,
		Title:       req.Title,
		Description: req.Description,
		MentorID:    uuid.MustParse(req.MentorID),
		CategoryID:  uuid.MustParse(req.CategoryID),
		Star:        req.Star,
		Price:       req.Price,
		Lessons:     req.Lessons,
	}

	err = s.CourseRepository.Insert(ctx, course)
	if err != nil {
		return err
	}

	return nil
}

// GetCourseByID retrieves a course by its ID.
func (s *courseService) GetCourseByID(ctx context.Context, id string) (res response.CourseResponse, err error) {
	course, err := s.CourseRepository.GetByID(ctx, uuid.MustParse(id))
	if err != nil {
		return res, err
	}

	res = response.CourseResponse{
		BannerImage: course.BannerImage,
		Title:       course.Title,
		Description: course.Description,
		Star:        course.Star,
		Price:       course.Price,
		Lessons:     course.Lessons,
		Mentor: response.MentorResponse{
			ID:         course.Mentor.ID.String(),
			Name:       course.Mentor.Name,
			Image:      course.Mentor.Image,
			Experience: course.Mentor.Experience,
		},
		Category: response.CourseCategoryResponse{
			ID:   course.Category.ID.String(),
			Name: course.Category.Name,
		},
	}

	return res, nil
}

// GetAllCourses retrieves all courses.
func (s *courseService) GetAllCourses(ctx context.Context) (res []response.CourseResponse, err error) {
	courses, err := s.CourseRepository.GetAll(ctx)
	if err != nil {
		return res, err
	}

	for _, course := range courses {
		var videoResponses []response.VideoResponse
		for _, video := range course.Videos {
			videoResponses = append(videoResponses, response.VideoResponse{
				ID:   video.ID.String(),
				Name: video.Name,
				URL:  video.URL,
				Time: video.Time,
			})
		}

		res = append(res, response.CourseResponse{
			ID:          course.ID.String(),
			BannerImage: course.BannerImage,
			Title:       course.Title,
			Description: course.Description,
			Star:        course.Star,
			Price:       course.Price,
			Lessons:     course.Lessons,
			Mentor: response.MentorResponse{
				ID:         course.Mentor.ID.String(),
				Name:       course.Mentor.Name,
				Image:      course.Mentor.Image,
				Experience: course.Mentor.Experience,
			},
			Category: response.CourseCategoryResponse{
				ID:   course.Category.ID.String(),
				Name: course.Category.Name,
			},
			Videos: videoResponses,
		})
	}

	return res, nil
}

// UpdateCourse updates a course.
func (s *courseService) UpdateCourse(ctx context.Context, id string, req request.UpdateCourseRequest) (err error) {
	course, err := s.CourseRepository.GetByID(ctx, uuid.MustParse(id))
	if err != nil {
		return err
	}

	course.BannerImage = req.BannerImage
	course.Title = req.Title
	course.Description = req.Description
	course.MentorID = uuid.MustParse(req.MentorID)
	course.CategoryID = uuid.MustParse(req.CategoryID)
	course.Star = req.Star
	course.Price = req.Price
	course.Lessons = req.Lessons

	if err := s.CourseRepository.Update(ctx, course); err != nil {
		return err
	}

	return nil
}

// DeleteCourse deletes a course.
func (s *courseService) DeleteCourse(ctx context.Context, id string) (err error) {
	err = s.CourseRepository.Delete(ctx, uuid.MustParse(id))
	if err != nil {
		return err
	}

	return nil
}

func (s *courseService) AddCourseVideo(ctx context.Context, courseID string, videoID string) error {
	parsedCourseID := uuid.MustParse(courseID)
	parsedVideoID := uuid.MustParse(videoID)

	return s.CourseRepository.InsertCourseVideo(ctx, parsedCourseID, parsedVideoID)
}
