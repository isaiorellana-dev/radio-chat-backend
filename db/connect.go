package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDB() (*gorm.DB, error) {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	DB_URL := os.Getenv("DB_URL")

	fmt.Println(DB_URL)

	db, err := gorm.Open(mysql.Open(DB_URL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error conectando a la base de datos: %w", err)
	}

	// err = db.AutoMigrate(&models.User{}, &models.Message{}, &models.Role{}, &models.Permission{})
	// if err != nil {
	// 	panic(err)
	// }

	return db, nil
}
