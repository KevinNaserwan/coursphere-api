package mentor

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinnaserwan/coursphere-api/internal/http/request"
	"github.com/kevinnaserwan/coursphere-api/internal/http/response"
	"github.com/kevinnaserwan/coursphere-api/internal/model"
	mentorRepository "github.com/kevinnaserwan/coursphere-api/internal/repository/mentor"
	errCommon "github.com/kevinnaserwan/coursphere-api/internal/util/error"
)

type mentorService struct {
	MentorRepository mentorRepository.Repository
}

func NewMentorService(mentorRepository mentorRepository.Repository) Service {
	return &mentorService{
		MentorRepository: mentorRepository,
	}
}

// CreateMentor creates a new mentor
func (s *mentorService) CreateMentor(ctx context.Context, req request.CreateMentorRequest) (err error) {
	mentor := model.Mentor{
		Name:       req.Name,
		Image:      req.Image,
		Experience: req.Experience,
	}

	if err := s.MentorRepository.CreateMentor(ctx, &mentor); err != nil {
		return errCommon.NewBadRequest("Failed to insert mentor: " + err.Error())
	}

	return nil
}

// GetMentorByID retrieves a mentor by its ID
func (s *mentorService) GetMentorByID(ctx context.Context, id string) (res response.MentorResponse, err error) {
	mentor, err := s.MentorRepository.GetMentorByID(ctx, uuid.MustParse(id))
	if err != nil {
		return res, errCommon.NewNotFound("Mentor not found: " + err.Error())
	}

	res = response.MentorResponse{
		ID:         mentor.ID.String(),
		Name:       mentor.Name,
		Image:      mentor.Image,
		Experience: mentor.Experience,
	}

	return res, nil

}

// GetAllMentors retrieves all mentors
func (s *mentorService) GetAllMentors(ctx context.Context) (res []response.MentorResponse, err error) {
	mentors, err := s.MentorRepository.GetAllMentors(ctx)
	if err != nil {
		return res, errCommon.NewNotFound("Failed to retrieve mentors: " + err.Error())
	}

	for _, mentor := range mentors {
		res = append(res, response.MentorResponse{
			ID:         mentor.ID.String(),
			Name:       mentor.Name,
			Image:      mentor.Image,
			Experience: mentor.Experience,
		})
	}

	return res, nil
}

// UpdateMentor updates a mentor
func (s *mentorService) UpdateMentor(ctx context.Context, id string, req request.UpdateMentorRequest) (err error) {
	mentor, err := s.MentorRepository.GetMentorByID(ctx, uuid.MustParse(id))
	if err != nil {
		return err
	}

	mentor.Name = req.Name
	mentor.Image = req.Image
	mentor.Experience = req.Experience

	if err := s.MentorRepository.UpdateMentor(ctx, mentor); err != nil {
		return errCommon.NewBadRequest("Failed to update mentor: " + err.Error())
	}

	return nil
}

// DeleteMentor deletes a mentor
func (s *mentorService) DeleteMentor(ctx context.Context, id string) (err error) {
	if err := s.MentorRepository.DeleteMentor(ctx, uuid.MustParse(id)); err != nil {
		return errCommon.NewNotFound("Failed to delete mentor: " + err.Error())
	}

	return nil
}
