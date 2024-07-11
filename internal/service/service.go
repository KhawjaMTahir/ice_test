package services

import "interview/internal/repository"

type Service interface {
	CartServiceInterface
}

type service struct {
	repository repository.CartRepositoryInterface
}

func NewService(repo repository.CartRepositoryInterface) Service {
	return &service{
		repository: repo,
	}
}
