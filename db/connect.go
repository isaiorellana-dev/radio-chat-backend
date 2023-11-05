package database

import (
	"fmt"
	"log"
	"os"

	"github.com/isaiorellana-dev/livechat-backend/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() (*gorm.DB, error) {
	if os.Getenv("ENVIRONMENT") == "development" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	DB_URL := os.Getenv("DB_URL")

	db, err := gorm.Open(postgres.Open(DB_URL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error conectando a la base de datos: %w", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Message{}, &models.Role{}, &models.Permission{})
	if err != nil {
		panic(err)
	}

	return db, nil
}
