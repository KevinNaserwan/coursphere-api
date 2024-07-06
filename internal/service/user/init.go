package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/kevinnaserwan/coursphere-api/internal/http/request"
	"github.com/kevinnaserwan/coursphere-api/internal/http/response"
	"github.com/kevinnaserwan/coursphere-api/internal/model"
	userRepository "github.com/kevinnaserwan/coursphere-api/internal/repository/user"
	errCommon "github.com/kevinnaserwan/coursphere-api/internal/util/error"
	"github.com/kevinnaserwan/coursphere-api/internal/util/jwt"
	PasswordUtil "github.com/kevinnaserwan/coursphere-api/internal/util/password"
)

type userService struct {
	UserRepository userRepository.Repository
	JwtManager     *jwt.JWTManager
}

func NewUserService(userRepository userRepository.Repository, jwtManager *jwt.JWTManager) Service {
	return &userService{
		UserRepository: userRepository,
		JwtManager:     jwtManager,
	}
}

// Register registers a new user
func (s *userService) Register(ctx context.Context, req request.UserRegisterRequest) (err error) {
	passwordHash, err := PasswordUtil.HashPassword(req.Password)
	if err != nil {
		return err
	}

	allUser, err := s.UserRepository.GetAll(ctx)
	if err != nil {
		return err
	}

	for _, user := range allUser {
		if user.Email == req.Email {
			return errCommon.NewBadRequest("Email already exists")
		}
	}

	user := model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: passwordHash,
	}

	if err := s.UserRepository.Insert(ctx, &user); err != nil {
		return err
	}

	return nil
}

// Login logs in a user
func (s *userService) Login(ctx context.Context, req request.UserLoginRequest) (JwtToken string, err error) {
	user, err := s.UserRepository.GetByEmail(ctx, req.Email)
	if err != nil {
		return "", err
	}

	err = PasswordUtil.CheckPasswordHash(req.Password, user.Password)
	if err != nil {
		return "", err
	}

	token, err := s.JwtManager.GenerateAuthToken(
		user.ID.String(),
		user.Email,
		"mahasiswa",
		24*time.Hour,
	)
	if err != nil {
		return "", errCommon.NewBadRequest("Failed to generate token: " + err.Error())
	}
	return token, nil
}

// GetByID retrieves a user by its ID
func (s *userService) GetByID(ctx context.Context, ID string) (user response.UserResponse, err error) {
	userModel, err := s.UserRepository.GetByID(ctx, uuid.MustParse(ID))
	if err != nil {
		return user, errCommon.NewBadRequest("Failed to get user: " + err.Error())
	}

	var userAchievements []response.UserAchievementResponse
	for _, ua := range userModel.UserAchievement {
		var achievements []response.AchievementResponse
		for _, achievement := range ua.Achievement {
			achievements = append(achievements, response.AchievementResponse{
				ID:          achievement.ID.String(),
				Name:        achievement.Name,
				Description: achievement.Description,
			})
		}

		userAchievements = append(userAchievements, response.UserAchievementResponse{
			ID:          ua.ID.String(),
			UserID:      ua.UserID.String(),
			Achievement: achievements,
		})
	}

	user = response.UserResponse{
		ID:              userModel.ID.String(),
		Username:        userModel.Username,
		Email:           userModel.Email,
		Profession:      userModel.Profession,
		UserAchievement: userAchievements,
	}

	return user, nil
}

// GetAll retrieves all users
func (s *userService) GetAll(ctx context.Context) (users []response.UserResponse, err error) {
	userModels, err := s.UserRepository.GetAll(ctx)
	if err != nil {
		return users, errCommon.NewBadRequest("Failed to get all users: " + err.Error())
	}

	for _, userModel := range userModels {
		var userAchievements []response.UserAchievementResponse
		for _, ua := range userModel.UserAchievement {
			var achievements []response.AchievementResponse
			for _, achievement := range ua.Achievement {
				achievements = append(achievements, response.AchievementResponse{
					ID:          achievement.ID.String(),
					Name:        achievement.Name,
					Description: achievement.Description,
				})
			}

			userAchievements = append(userAchievements, response.UserAchievementResponse{
				ID:          ua.ID.String(),
				UserID:      ua.UserID.String(),
				Achievement: achievements,
			})
		}

		user := response.UserResponse{
			ID:              userModel.ID.String(),
			Username:        userModel.Username,
			Email:           userModel.Email,
			Profession:      userModel.Profession,
			UserAchievement: userAchievements,
		}

		users = append(users, user)
	}

	return users, nil
}

// Update updates a user
func (s *userService) Update(ctx context.Context, ID string, req request.UserUpdateRequest) (err error) {
	user, err := s.UserRepository.GetByID(ctx, uuid.MustParse(ID))
	if err != nil {
		return errCommon.NewBadRequest("Failed to update user: " + err.Error())
	}

	user.Username = req.Username

	if err := s.UserRepository.Update(ctx, user); err != nil {
		return errCommon.NewBadRequest("Failed to update user: " + err.Error())
	}

	return nil
}

// Delete deletes a user
func (s *userService) Delete(ctx context.Context, ID string) (err error) {
	if err := s.UserRepository.Delete(ctx, uuid.MustParse(ID)); err != nil {
		return errCommon.NewBadRequest("Failed to delete user: " + err.Error())
	}

	return nil
}
