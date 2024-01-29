package postgresql

import (
	"fmt"
	"go-gin-boilerplate/configs"
	"go-gin-boilerplate/internal/models"
	"go-gin-boilerplate/internal/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(config *configs.Postgresql) *gorm.DB {

	logger := logger.LogrusLogger

	address := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		config.Host,
		config.User,
		config.Password,
		config.DataBase,
		config.Port,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: address,
	}), &gorm.Config{})

	if err != nil {
		panic(`😫: Connected failed, check your PostgreSQL with ` + address)
	}

	// Migrate the schema
	migrateErr := db.AutoMigrate(&models.Example{}, &models.User{})
	if migrateErr != nil {
		panic(`😫: Auto migrate failed, check your PostgreSQL with ` + address)
	}

	// export DB
	DB = db

	logger.Printf(`🍟: Successfully connected to PostgreSQL at ` + address)

	return db

}
