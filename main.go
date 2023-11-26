package main

import (
	"log"
	"os"

	"github.com/isaiorellana-dev/livechat-backend/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	if os.Getenv("ENVIRONMENT") == "development" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	PORT := os.Getenv("PORT")
	FRONTEND_APP := os.Getenv("FRONTEND_APP")

	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{FRONTEND_APP},
		AllowHeaders: middleware.DefaultCORSConfig.AllowHeaders,
		AllowMethods: middleware.DefaultCORSConfig.AllowMethods,
	}))

	routes.RegisterRoutes(e)

	if os.Getenv("ENVIRONMENT") == "development" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	if err := e.Start(":" + PORT); err != nil {
		e.Logger.Fatal(err)
	}
}
