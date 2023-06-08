package routes

import (
	h "github.com/isaiorellana-dev/radio-api/handlers"
	m "github.com/isaiorellana-dev/radio-api/middlewares"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	// Get methods
	e.GET("/api/v1/", h.HelloWorld)
	e.GET("/api/v1/users", h.GetUsers)

	// Post methods
	e.POST("/api/v1/users", h.CreateUser, m.ValidateUser)

	// Put methods
	e.PUT("/api/v1/users/:id", h.UpdateUser, m.ValidateUser)
}
