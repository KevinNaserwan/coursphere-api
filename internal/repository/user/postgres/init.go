package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinnaserwan/coursphere-api/internal/model"
	UserRepository "github.com/kevinnaserwan/coursphere-api/internal/repository/user"
	errCommon "github.com/kevinnaserwan/coursphere-api/internal/util/error"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository.Repository {
	return &userRepository{
		db: db,
	}
}

// Insert inserts a new user into the database
func (r *userRepository) Insert(ctx context.Context, user *model.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return errCommon.NewBadRequest("Failed to insert user: " + err.Error())
	}
	return nil
}

// GetByID retrieves a user by its ID
func (r *userRepository) GetByID(ctx context.Context, ID uuid.UUID) (*model.User, error) {
	user := &model.User{}
	if err := r.db.Where("id = ?", ID).First(user).Error; err != nil {
		return nil, errCommon.NewNotFound("User not found: " + err.Error())
	}
	return user, nil
}

// Update updates a user in the database
func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return errCommon.NewBadRequest("Failed to update user: " + err.Error())
	}
	return nil
}

// Delete deletes a user from the database
func (r *userRepository) Delete(ctx context.Context, ID uuid.UUID) error {
	if err := r.db.Where("id = ?", ID).Delete(&model.User{}).Error; err != nil {
		return errCommon.NewBadRequest("Failed to delete user: " + err.Error())
	}
	return nil
}

// GetAll retrieves all users from the database
func (r *userRepository) GetAll(ctx context.Context) ([]model.User, error) {
	users := []model.User{}
	if err := r.db.Find(&users).Error; err != nil {
		return nil, errCommon.NewNotFound("No users found: " + err.Error())
	}
	return users, nil
}

// GetByEmail retrieves a user by its email
func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}
	if err := r.db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, errCommon.NewNotFound("User not found: " + err.Error())
	}
	return user, nil
}
