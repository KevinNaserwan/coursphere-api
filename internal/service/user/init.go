package user

import (
	"context"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/kevinnaserwan/coursphere-api/internal/http/request"
	"github.com/kevinnaserwan/coursphere-api/internal/http/response"
	"github.com/kevinnaserwan/coursphere-api/internal/model"
	userRepository "github.com/kevinnaserwan/coursphere-api/internal/repository/user"
	errCommon "github.com/kevinnaserwan/coursphere-api/internal/util/error"
	"github.com/kevinnaserwan/coursphere-api/internal/util/jwt"
	"github.com/kevinnaserwan/coursphere-api/internal/util/mail"
	PasswordUtil "github.com/kevinnaserwan/coursphere-api/internal/util/password"
	"gopkg.in/gomail.v2"
)

type userService struct {
	UserRepository userRepository.Repository
	JwtManager     *jwt.JWTManager
	MailDialer     *gomail.Dialer
}

func NewUserService(userRepository userRepository.Repository, jwtManager *jwt.JWTManager, mailDialer *gomail.Dialer) Service {
	return &userService{
		UserRepository: userRepository,
		JwtManager:     jwtManager,
		MailDialer:     mailDialer,
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

	// Generate OTP
	otp := generateOTP()

	user := model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: passwordHash,
		AuthCode: otp, // Save OTP to AuthCode
	}

	if err := s.UserRepository.Insert(ctx, &user); err != nil {
		return err
	}

	// Render email verification template
	emailData := mail.EmailVerification{
		OTP: otp,
	}
	emailBody, err := mail.RenderEmailVerificationTemplate(emailData)
	if err != nil {
		return err
	}

	// Send OTP to user's email
	msg := mail.NewMessage(os.Getenv("EMAIL_FROM"), req.Email, "Coursphere - Email Verification", emailBody)
	if err := s.MailDialer.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}

func generateOTP() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const otpLength = 4

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	otp := strings.Builder{}
	otp.Grow(otpLength)
	for i := 0; i < otpLength; i++ {
		randomChar := charset[seededRand.Intn(len(charset))]
		otp.WriteByte(randomChar)
	}

	return otp.String()
}

// ResendOTPByEmail sends a new OTP to the user's email based on the email address
func (s *userService) ResendOTPByEmail(ctx context.Context, email string) error {
	// Get user by email
	user, err := s.UserRepository.GetByEmail(ctx, email)
	if err != nil {
		return errCommon.NewBadRequest("User not found: " + err.Error())
	}

	// Generate a new OTP
	newOTP := generateOTP()

	// Update user's AuthCode with new OTP
	user.AuthCode = newOTP
	if err := s.UserRepository.Update(ctx, user); err != nil {
		return errCommon.NewBadRequest("Failed to update user: " + err.Error())
	}

	// Render email verification template
	emailData := mail.EmailVerification{
		OTP: newOTP,
	}
	emailBody, err := mail.RenderEmailVerificationTemplate(emailData)
	if err != nil {
		return errCommon.NewBadRequest("Failed to render email template: " + err.Error())
	}

	// Send OTP to user's email
	msg := mail.NewMessage(os.Getenv("EMAIL_FROM"), user.Email, "Resend Email Verification", emailBody)
	if err := s.MailDialer.DialAndSend(msg); err != nil {
		return errCommon.NewBadRequest("Failed to send email: " + err.Error())
	}

	return nil
}

// VerifyAuthToken verifies the OTP sent to the user's email
func (s *userService) VerifyAuthToken(ctx context.Context, req request.UserOTPVerifyRequest) (err error) {
	user, err := s.UserRepository.GetByEmail(ctx, req.Email)
	if err != nil {
		return errCommon.NewBadRequest("User not found: " + err.Error())
	}

	if user.AuthCode != req.AuthCode {
		return errCommon.NewBadRequest("Invalid OTP")
	}

	user.IsVerified = true
	user.AuthCode = "" // Clear AuthCode upon successful verification

	if err := s.UserRepository.Update(ctx, user); err != nil {
		return errCommon.NewBadRequest("Failed to update user: " + err.Error())
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
