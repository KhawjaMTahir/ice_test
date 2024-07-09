package db

import (
	"fmt"
	"interview/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDatabase() *gorm.DB {
	cfg := config.LoadConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
