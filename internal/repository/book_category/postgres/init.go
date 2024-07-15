package bookcategory

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinnaserwan/coursphere-api/internal/model"
	BookCategoryRepository "github.com/kevinnaserwan/coursphere-api/internal/repository/book_category"
	errCommon "github.com/kevinnaserwan/coursphere-api/internal/util/error"
	"gorm.io/gorm"
)

type bookCategoryRepository struct {
	db *gorm.DB
}

func NewBookCategoryRepository(db *gorm.DB) BookCategoryRepository.Repository {
	return &bookCategoryRepository{
		db: db,
	}
}

// Insert inserts a new book category into the database
func (r *bookCategoryRepository) Insert(ctx context.Context, bookCategory *model.CategoryBook) error {
	if err := r.db.Create(bookCategory).Error; err != nil {
		return err
	}
	return nil
}

// GetByID retrieves a book category by its ID
func (r *bookCategoryRepository) GetByID(ctx context.Context, ID uuid.UUID) (*model.CategoryBook, error) {
	bookCategory := &model.CategoryBook{}
	if err := r.db.Where("id = ?", ID).First(bookCategory).Error; err != nil {
		return nil, errCommon.NewNotFound("User not found: " + err.Error())
	}
	return bookCategory, nil
}

// Update updates a book category in the database
func (r *bookCategoryRepository) Update(ctx context.Context, bookCategory *model.CategoryBook) error {
	if err := r.db.Save(bookCategory).Error; err != nil {
		return err
	}
	return nil
}

// Delete deletes a book category from the database
func (r *bookCategoryRepository) Delete(ctx context.Context, ID uuid.UUID) error {
	if err := r.db.Where("id = ?", ID).Delete(&model.CategoryBook{}).Error; err != nil {
		return err
	}
	return nil
}

// GetAll retrieves all book categories from the database
func (r *bookCategoryRepository) GetAll(ctx context.Context) ([]model.CategoryBook, error) {
	bookCategories := []model.CategoryBook{}
	if err := r.db.Find(&bookCategories).Error; err != nil {
		return nil, err
	}
	return bookCategories, nil
}
