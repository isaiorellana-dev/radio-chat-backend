package routes

import (
	h "github.com/isaiorellana-dev/radio-chat-backend/handlers"
	m "github.com/isaiorellana-dev/radio-chat-backend/middlewares"
	ws "github.com/isaiorellana-dev/radio-chat-backend/websocket"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	hub := ws.NewHub()
	go hub.Run()

	e.Use(m.CustomContextMiddleware(hub), m.CheckJWT)

	e.GET("/ws", func(c echo.Context) error {
		return ws.ServeWs(hub, c)
	})

	// Get methods
	e.GET("/api/v1/hello", h.HelloWorld)
	e.GET("/api/v1/users", h.GetUsers, m.CheckPermissions(h.GetUsersPerms))
	e.GET("/api/v1/users/:id", h.GetOneUser)
	e.GET("/api/v1/messages", h.GetMessages)

	// Post methods
	e.POST("/api/v1/signup", h.Register, m.ValidateUser)
	e.POST("/api/v1/message", h.CreateMessage,
		m.CheckPermissions(h.CreateMessagePerms),
		m.ValidateUserByID,
		m.ValidateMessage)
	e.POST("/api/v1/login", h.Login)
	e.POST("/api/v1/role", h.CreateRole, m.CheckPermissions(h.CreateRolePerms))
	e.POST("/api/v1/permission", h.CreatePermission, m.CheckPermissions(h.CreatePermissionsPerms))

	// Put methods
	e.PUT("/api/v1/users/:id", h.UpdateUser, m.ValidateUser)

	// Delete methods
	e.DELETE("api/v1/users/:id", h.DeleteUser, m.CheckPermissions(h.DeleteUserPerms))
	e.DELETE("/api/v1/message/:id", h.DeleteMessage, m.CheckPermissions(h.DeleteMessagePerms))
}
