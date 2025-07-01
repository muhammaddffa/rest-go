package domain

import (
	"context"
	"database/sql"

	"shellrean.id/belajar-golang-rest-api/dto"
)

type Customer struct {
	ID   		string 		`db:"id"`
	Code 		string 		`db:"code"`
	Name 		string 		`db:"name"`
	CreatedAt 	sql.NullTime	`db:"created_at"`
	UpdatedAt 	sql.NullTime	`db:"updated_at"`
	DeletedAt 	sql.NullTime	`db:"deleted_at"`
}

type CustomerRepository interface {
	FindAll(ctx context.Context) ([]Customer, error)
	FindById(ctx context.Context, id string) (Customer, error)
	Save(ctx context.Context, c *Customer) error
	Update(ctx context.Context, c *Customer) error
	Delete(ctx context.Context, id string) error
}

type CustomerService interface {
	Index(ctx context.Context) ([]dto.CustomerData, error)
}