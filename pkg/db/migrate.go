package db

import (
	"interview/pkg/entity"

	"gorm.io/gorm"
)

func MigrateDatabase(db *gorm.DB) {

	// AutoMigrate will create or update the tables based on the models
	err := db.AutoMigrate(&entity.CartEntity{}, &entity.CartItem{})
	if err != nil {
		panic(err)
	}
}
