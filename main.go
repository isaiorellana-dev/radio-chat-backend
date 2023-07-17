package main

import (
	"fmt"
	"os"

	"github.com/isaiorellana-dev/radio-chat-backend/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	routes.RegisterRoutes(e)

	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	PORT := os.Getenv("PORT")

	fmt.Println(os.Getenv("PORT"))

	if err := e.Start(":" + PORT); err != nil {
		e.Logger.Fatal(err)
	}
}
