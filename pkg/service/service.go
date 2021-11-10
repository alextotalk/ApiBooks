package service

import (
	"github.com/alextotalk/ApiBooks"
	"github.com/alextotalk/ApiBooks/pkg/repository"
)

type Authorization interface {
	CreateUser(user ApiBooks.User) (int, error)
}
type BooksStorage interface {
}

type Service struct {
	Authorization
	BooksStorage
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
