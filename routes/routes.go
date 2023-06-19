package routes

import (
	h "github.com/isaiorellana-dev/radio-chat-backend/handlers"
	m "github.com/isaiorellana-dev/radio-chat-backend/middlewares"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	// Get methods
	e.GET("/api/v1/", h.HelloWorld)
	e.GET("/api/v1/users", h.GetUsers)
	e.GET("/api/v1/users/:id", h.GetOneUser)
	e.GET("/api/v1/messages", h.GetMessages, m.CheckJWT)

	// Post methods
	e.POST("/api/v1/register", h.Register, m.ValidateUser)
	e.POST("/api/v1/messages", h.CreateMessage, m.ValidateUserByID, m.ValidateMessage)
	e.POST("/api/v1/login", h.Login)
	e.POST("/api/v1/role", h.CreateRole)
	e.POST("/api/v1/permission", h.CreatePermission)

	// Put methods
	e.PUT("/api/v1/users/:id", h.UpdateUser, m.ValidateUser)

	// Delete methods
	e.DELETE("api/v1/users/:id", h.DeleteUser)
	e.DELETE("/api/v1/messages/:id", h.DeleteMessage)
}
