package repository

import (
	"github.com/alextotalk/ApiBooks"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user ApiBooks.User) (int, error)
	GetUser(username, password string) (ApiBooks.User, error)
}
type BooksStorage interface {
}

type Repository struct {
	Authorization
	BooksStorage
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
