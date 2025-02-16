package repository

import "gorm.io/gorm"

type (
	Repository interface {
		CartRepositoryInterface
	}

	repository struct {
		db *gorm.DB
	}
)

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}

}
