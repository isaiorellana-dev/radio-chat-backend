package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDB() (*gorm.DB, error) {
	user := "root"
	password := "root"
	port := "3306"
	ip := "db-server"
	// Change the name of the schema
	schema := "my-schema"

	dsn := user + ":" + password + "@tcp(" + ip + ":" + port + ")/" + schema + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error conectando a la base de datos: %w", err)
	}

	return db, nil
}
