package main

import (
	"log"
	"os"

	"github.com/isaiorellana-dev/radio-chat-backend/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	routes.RegisterRoutes(e)

	if os.Getenv("ENVIRONMENT") == "development" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	PORT := os.Getenv("PORT")

	if err := e.Start(":" + PORT); err != nil {
		e.Logger.Fatal(err)
	}
}
