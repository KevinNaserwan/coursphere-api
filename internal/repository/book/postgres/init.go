package book

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinnaserwan/coursphere-api/internal/model"
	BookRepository "github.com/kevinnaserwan/coursphere-api/internal/repository/book"
	errCommon "github.com/kevinnaserwan/coursphere-api/internal/util/error"
	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository.Repository {
	return &bookRepository{
		db: db,
	}
}

// Insert inserts a new book into the database
func (r *bookRepository) Insert(ctx context.Context, book *model.Book) error {
	if err := r.db.Create(book).Error; err != nil {
		return errCommon.NewBadRequest("Failed to insert book: " + err.Error())
	}
	return nil
}

// GetByID retrieves a book by its ID
func (r *bookRepository) GetByID(ctx context.Context, ID uuid.UUID) (*model.Book, error) {
	book := &model.Book{}
	if err := r.db.Where("id = ?", ID).First(book).Error; err != nil {
		return nil, errCommon.NewBadRequest("Book not found: " + err.Error())
	}
	return book, nil
}

// GetAll retrieves all books from the database
func (r *bookRepository) GetAll(ctx context.Context) ([]model.Book, error) {
	books := []model.Book{}
	if err := r.db.Find(&books).Error; err != nil {
		return nil, errCommon.NewNotFound("No books found: " + err.Error())
	}
	return books, nil
}

// Update updates a book in the database
func (r *bookRepository) Update(ctx context.Context, book *model.Book) error {
	if err := r.db.Save(book).Error; err != nil {
		return errCommon.NewBadRequest("Failed to update book: " + err.Error())
	}
	return nil
}

// Delete deletes a book from the database
func (r *bookRepository) Delete(ctx context.Context, ID uuid.UUID) error {
	if err := r.db.Where("id = ?", ID).Delete(&model.Book{}).Error; err != nil {
		return errCommon.NewBadRequest("Failed to delete book: " + err.Error())
	}
	return nil
}
