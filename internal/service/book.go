package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"shellrean.id/belajar-golang-rest-api/domain"
	"shellrean.id/belajar-golang-rest-api/dto"
)

type bookService struct {
	bookRepository      domain.BookRepository
	bookStockRepository domain.BookStockRepository
}

func NewBook(bookRepository domain.BookRepository,
	bookStockRepository domain.BookStockRepository) domain.BookService {
	return &bookService{
		bookRepository:      bookRepository,
		bookStockRepository: bookStockRepository,
	}
}

// Create implements domain.BookService.
func (b *bookService) Create(ctx context.Context, req dto.CreateBookRequest) error {
	book := domain.Book{
		Id:          uuid.NewString(),
		Isbn:        req.Isbn,
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   sql.NullTime{Valid: true, Time: time.Now()},
	}

	return b.bookRepository.Save(ctx, &book)
}

// Delete implements domain.BookService.
func (b *bookService) Delete(ctx context.Context, id string) error {
	persisted, err := b.bookRepository.FindById(ctx, id)
	if err != nil {
		return err
	}

	if persisted.Id == "" {
		return errors.New("data buku tidak ditemukan")
	}

	err = b.bookRepository.Delete(ctx, persisted.Id)

	if err != nil {
		return err
	}
	return b.bookStockRepository.DeleteByBookId(ctx, persisted.Id)
}

// Index implements domain.BookService.
func (b *bookService) Index(ctx context.Context) ([]dto.BookData, error) {
	books, err := b.bookRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var booksData []dto.BookData
	for _, v := range books {
		booksData = append(booksData, dto.BookData{
			Id:          v.Id,
			Isbn:        v.Isbn,
			Title:       v.Title,
			Description: v.Description,
		})
	}
	return booksData, nil
}

// Show implements domain.BookService.
func (b *bookService) Show(ctx context.Context, id string) (dto.BookData, error) {
	data, err := b.bookRepository.FindById(ctx, id)
	if err != nil {
		return dto.BookData{}, err
	}

	if data.Id == "" {
		return dto.BookData{}, errors.New("data buku tidak ditemukan")
	}

	return dto.BookData{
		Id:          data.Id,
		Isbn:        data.Isbn,
		Title:       data.Title,
		Description: data.Description,
	}, nil
}

// Update implements domain.BookService.
func (b *bookService) Update(ctx context.Context, req dto.UpdateBookRequest) error {
	persisted, err := b.bookRepository.FindById(ctx, req.Id)
	if err != nil {
		return err
	}

	if persisted.Id == "" {
		return errors.New("data buku tidak ditemukan")
	}

	persisted.Isbn = req.Isbn
	persisted.Title = req.Title
	persisted.Description = req.Description
	persisted.UpdateAt = sql.NullTime{Valid: true, Time: time.Now()}
	return b.bookRepository.Update(ctx, &persisted)
}
