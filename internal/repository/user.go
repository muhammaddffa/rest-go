package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"shellrean.id/belajar-golang-rest-api/domain"
)

type userRepository struct {
	db *goqu.Database
}

func NewUser(con *sql.DB) domain.UserRepository {
	return &userRepository{
		db: goqu.New("deafult", con),
	}
}

func (u userRepository) FindByEmail(ctx context.Context, email string) (usr domain.User, err error) {
	dataset := u.db.From("users").Where(goqu.C("email").Eq(email))
	_, err = dataset.ScanStructContext(ctx, &usr)
	return
}

