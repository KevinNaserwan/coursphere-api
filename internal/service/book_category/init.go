package bookcategory

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinnaserwan/coursphere-api/internal/http/request"
	"github.com/kevinnaserwan/coursphere-api/internal/http/response"
	"github.com/kevinnaserwan/coursphere-api/internal/model"
	bookCategoryRepository "github.com/kevinnaserwan/coursphere-api/internal/repository/book_category"
	errCommon "github.com/kevinnaserwan/coursphere-api/internal/util/error"
)

type bookCategoryService struct {
	BookCategoryRepository bookCategoryRepository.Repository
}

func NewBookCategoryService(bookCategoryRepository bookCategoryRepository.Repository) Service {
	return &bookCategoryService{
		BookCategoryRepository: bookCategoryRepository,
	}
}

// Insert inserts a new book category into the database
func (s *bookCategoryService) Insert(ctx context.Context, req request.CreateBookCategoryRequest) (err error) {
	bookCategory := &model.CategoryBook{
		Name: req.Name,
	}

	if err := s.BookCategoryRepository.Insert(ctx, bookCategory); err != nil {
		return err
	}

	return nil
}

// GetByID retrieves a book category by its ID
func (s *bookCategoryService) GetByID(ctx context.Context, ID string) (bookcategory response.BookCategoryResponse, err error) {
	bookCategory, err := s.BookCategoryRepository.GetByID(ctx, uuid.MustParse(ID))
	if err != nil {
		return bookcategory, errCommon.NewBadRequest("Failed to get user: " + err.Error())
	}

	bookcategory = response.BookCategoryResponse{
		ID:   bookCategory.ID.String(),
		Name: bookCategory.Name,
	}

	return bookcategory, nil
}

// GetAll retrieves all book categories from the database
func (s *bookCategoryService) GetAll(ctx context.Context) (bookcategories []response.BookCategoryResponse, err error) {
	bookCategories, err := s.BookCategoryRepository.GetAll(ctx)
	if err != nil {
		return bookcategories, err
	}

	for _, bookCategory := range bookCategories {
		bookcategories = append(bookcategories, response.BookCategoryResponse{
			ID:   bookCategory.ID.String(),
			Name: bookCategory.Name,
		})
	}

	return bookcategories, nil
}

// Delete deletes a book category from the database
func (s *bookCategoryService) Delete(ctx context.Context, ID string) (err error) {
	if err := s.BookCategoryRepository.Delete(ctx, uuid.MustParse(ID)); err != nil {
		return err
	}

	return nil
}
