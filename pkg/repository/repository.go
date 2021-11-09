package repository

type Authorization interface {
}
type BooksStorage interface {
}

type Repository struct {
	Authorization
	BooksStorage
}

func NewRepository() *Repository {
	return &Repository{}
}
