package service

import "github.com/alextotalk/ApiBooks/pkg/repository"

type Authorization interface {
}
type BooksStorage interface {
}

type Service struct {
	Authorization
	BooksStorage
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
