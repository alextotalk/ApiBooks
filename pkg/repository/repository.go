package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}
type BooksStorage interface {
}

type Repository struct {
	Authorization
	BooksStorage
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
