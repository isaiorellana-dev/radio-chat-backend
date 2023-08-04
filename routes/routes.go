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

	prefix := "/api/v1"

	// Hello world
	e.GET(prefix+"/hi", h.HelloWorld)

	// Websocket
	e.GET("/ws", func(c echo.Context) error {
		return ws.ServeWs(hub, c)
	})

	// Users
	e.GET(prefix+"/users", h.GetUsers, m.CheckPermissions(h.GetUsersPerms))
	e.GET(prefix+"/users/:id", h.GetOneUser)
	e.DELETE("api/v1/users/:id", h.DeleteUser, m.CheckPermissions(h.DeleteUserPerms))
	e.PUT(prefix+"/users/:id", h.UpdateUser, m.ValidateUser)
	e.POST(prefix+"/signup", h.Register, m.ValidateUser)
	e.POST(prefix+"/login", h.Login)
	e.GET(prefix+"/auth/user", h.UserData)

	// Messages
	e.GET(prefix+"/messages", h.GetMessages)
	e.POST(prefix+"/messages", h.CreateMessage,
		m.CheckPermissions(h.CreateMessagePerms),
		m.ValidateUserByID,
		m.ValidateMessage)
	e.DELETE(prefix+"/messages/:id", h.DeleteMessage, m.CheckPermissions(h.DeleteMessagePerms))

	// Roles and permissions
	e.POST(prefix+"/role", h.CreateRole, m.CheckPermissions(h.CreateRolePerms))
	e.POST(prefix+"/permission", h.CreatePermission, m.CheckPermissions(h.CreatePermissionsPerms))

}
