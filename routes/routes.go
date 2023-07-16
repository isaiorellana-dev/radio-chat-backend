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

	// Hello world
	e.GET("/api/v1/hello", h.HelloWorld)

	// Websocket
	e.GET("/ws", func(c echo.Context) error {
		return ws.ServeWs(hub, c)
	})

	// Users
	e.GET("/api/v1/users", h.GetUsers, m.CheckPermissions(h.GetUsersPerms))
	e.GET("/api/v1/users/:id", h.GetOneUser)
	e.DELETE("api/v1/users/:id", h.DeleteUser, m.CheckPermissions(h.DeleteUserPerms))
	e.PUT("/api/v1/users/:id", h.UpdateUser, m.ValidateUser)
	e.POST("/api/v1/signup", h.Register, m.ValidateUser)
	e.POST("/api/v1/login", h.Login)

	// Messages
	e.GET("/api/v1/messages", h.GetMessages)
	e.POST("/api/v1/messages", h.CreateMessage,
		m.CheckPermissions(h.CreateMessagePerms),
		m.ValidateUserByID,
		m.ValidateMessage)
	e.DELETE("/api/v1/messages/:id", h.DeleteMessage, m.CheckPermissions(h.DeleteMessagePerms))

	// Roles and permissions
	e.POST("/api/v1/role", h.CreateRole, m.CheckPermissions(h.CreateRolePerms))
	e.POST("/api/v1/permission", h.CreatePermission, m.CheckPermissions(h.CreatePermissionsPerms))

}
