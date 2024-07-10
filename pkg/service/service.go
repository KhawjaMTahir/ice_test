package services

import "interview/pkg/repository"

type Service interface {
	CartServiceInterface
}

type service struct {
	repository repository.CartRepositoryInterface
}

func NewService(repo repository.CartRepositoryInterface) *service {
	return &service{
		repository: repo,
	}
}
