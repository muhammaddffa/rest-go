package domain

import (
	"context"
	"database/sql"
)

type BookStock struct {
	Code       string       `db:"code"`
	BookId     string       `db:"book_id"`
	Status     string       `db:"status"`
	BorrowedId sql.NullTime `db:"borrowed_id"`
	BorrowedAt sql.NullTime `db:"boroowed_at"`
}

type BookStockRepository interface {
	FindByBookId(ctx context.Context, id string) ([]BookStock, error)
	FindByBookAndCode(ctx context.Context, id string, code string) (BookStock, error)
	Save(ctx context.Context, data []BookStock) error
	Update(ctx context.Context, data *BookStock) error
	DeleteByBookId(ctx context.Context, id string) error
	DeleteByCodes(ctx context.Context, codes []string) error
}
