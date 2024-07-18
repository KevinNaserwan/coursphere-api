package book

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinnaserwan/coursphere-api/internal/http/request"
	"github.com/kevinnaserwan/coursphere-api/internal/http/response"
	"github.com/kevinnaserwan/coursphere-api/internal/model"
	bookRepository "github.com/kevinnaserwan/coursphere-api/internal/repository/book"
	errCommon "github.com/kevinnaserwan/coursphere-api/internal/util/error"
)

type bookService struct {
	BookRepository bookRepository.Repository
}

func NewBookService(bookRepository bookRepository.Repository) Service {
	return &bookService{
		BookRepository: bookRepository,
	}
}

// CreateBook creates a new book
func (s *bookService) CreateBook(ctx context.Context, req request.CreateBookRequest) (res response.BookResponse, err error) {
	book := &model.Book{
		Title:          req.Title,
		Language:       req.Language,
		Rank:           req.Rank,
		ReadingTime:    req.ReadingTime,
		Likes:          req.Likes,
		BookFile:       req.BookFile,
		Overview:       req.Overview,
		Writer:         req.Writer,
		CategoryBookID: uuid.MustParse(req.CategoryBookID),
	}

	err = s.BookRepository.Insert(ctx, book)
	if err != nil {
		return res, errCommon.NewBadRequest("Failed to insert book: " + err.Error())
	}

	res = response.BookResponse{
		ID:          book.ID.String(),
		Title:       book.Title,
		Language:    book.Language,
		Rank:        book.Rank,
		ReadingTime: book.ReadingTime,
		Likes:       book.Likes,
		BookFile:    book.BookFile,
		Overview:    book.Overview,
		Writer:      book.Writer,
	}

	return res, nil
}

// GetBookByID retrieves a book by its ID
func (s *bookService) GetBookByID(ctx context.Context, ID string) (res response.BookResponse, err error) {
	bookID := uuid.MustParse(ID)
	book, err := s.BookRepository.GetByID(ctx, bookID)
	if err != nil {
		return res, errCommon.NewBadRequest("Failed to get book: " + err.Error())
	}

	res = response.BookResponse{
		ID:          book.ID.String(),
		Title:       book.Title,
		Language:    book.Language,
		Rank:        book.Rank,
		ReadingTime: book.ReadingTime,
		Likes:       book.Likes,
		BookFile:    book.BookFile,
		Overview:    book.Overview,
		Writer:      book.Writer,
	}

	return res, nil
}

// GetAllBooks retrieves all books
func (s *bookService) GetAllBooks(ctx context.Context) (res []response.BookResponse, err error) {
	books, err := s.BookRepository.GetAll(ctx)
	if err != nil {
		return res, errCommon.NewNotFound("Failed to get books: " + err.Error())
	}

	for _, book := range books {
		res = append(res, response.BookResponse{
			ID:          book.ID.String(),
			Title:       book.Title,
			Language:    book.Language,
			Rank:        book.Rank,
			ReadingTime: book.ReadingTime,
			Likes:       book.Likes,
			BookFile:    book.BookFile,
			Overview:    book.Overview,
			Writer:      book.Writer,
		})
	}

	return res, nil

}

// UpdateBook updates a book
func (s *bookService) UpdateBook(ctx context.Context, ID string, req request.UpdateBookRequest) (res response.BookResponse, err error) {
	bookID := uuid.MustParse(ID)
	book, err := s.BookRepository.GetByID(ctx, bookID)
	if err != nil {
		return res, errCommon.NewBadRequest("Failed to get book: " + err.Error())
	}

	book.Title = req.Title
	book.Language = req.Language
	book.Rank = req.Rank
	book.ReadingTime = req.ReadingTime
	book.Likes = req.Likes
	book.BookFile = req.BookFile
	book.Overview = req.Overview
	book.Writer = req.Writer

	err = s.BookRepository.Update(ctx, book)
	if err != nil {
		return res, errCommon.NewBadRequest("Failed to update book: " + err.Error())
	}

	res = response.BookResponse{
		ID:          book.ID.String(),
		Title:       book.Title,
		Language:    book.Language,
		Rank:        book.Rank,
		ReadingTime: book.ReadingTime,
		Likes:       book.Likes,
		BookFile:    book.BookFile,
		Overview:    book.Overview,
		Writer:      book.Writer,
	}

	return res, nil
}

// DeleteBook deletes a book
func (s *bookService) DeleteBook(ctx context.Context, ID string) (err error) {
	bookID := uuid.MustParse(ID)
	err = s.BookRepository.Delete(ctx, bookID)
	if err != nil {
		return errCommon.NewBadRequest("Failed to delete book: " + err.Error())
	}

	return nil
}
