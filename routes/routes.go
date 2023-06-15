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
	e.GET("/api/v1/users/:id", h.GetOneUser)
	e.GET("/api/v1/messages", h.GetMessages)

	// Post methods
	e.POST("/api/v1/users", h.CreateUser, m.ValidateUser)
	e.POST("/api/v1/messages", h.CreateMessage, m.ValidateUserExist, m.ValidateMessage)

	// Put methods
	e.PUT("/api/v1/users/:id", h.UpdateUser, m.ValidateUser)

	// Delete methods
	e.DELETE("api/v1/users/:id", h.DeleteUser)
	e.DELETE("/api/v1/messages/:id", h.DeleteMessage)
}
