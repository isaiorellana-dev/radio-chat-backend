package database

import (
	"fmt"

	"github.com/isaiorellana-dev/radio-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDB() (*gorm.DB, error) {
	user := "root"
	password := "root"
	port := "3306"
	// para probar el dockerfile cambiar localhost por "db-server"
	ip := "localhost"
	// Change the name of the schema
	schema := "my-schema"

	dsn := user + ":" + password + "@tcp(" + ip + ":" + port + ")/" + schema + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error conectando a la base de datos: %w", err)
	}

	// Ejecuta la migración automática para crear las tablas y el esquema
	err = db.AutoMigrate(&models.User{}, &models.Message{})
	if err != nil {
		panic(err)
	}

	return db, nil
}
