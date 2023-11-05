package routes

import (
	h "github.com/isaiorellana-dev/livechat-backend/handlers"
	m "github.com/isaiorellana-dev/livechat-backend/middlewares"
	ws "github.com/isaiorellana-dev/livechat-backend/websocket"
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
	e.GET(prefix+"/users/:id", h.GetOneUser, m.CheckPermissions(h.GetUsersPerms))
	e.DELETE(prefix+"/users/:id", h.DeleteUser, m.CheckPermissions(h.DeleteUserPerms))
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
	e.GET(prefix+"/roles", h.GetRoles, m.CheckPermissions(h.GetRolesPerms))
	e.GET(prefix+"/permissions", h.GetPermissions, m.CheckPermissions(h.GetPermissionsPerms))
	e.POST(prefix+"/roles", h.CreateRole, m.CheckPermissions(h.CreateRolePerms))
	e.POST(prefix+"/permissions", h.CreatePermission, m.CheckPermissions(h.CreatePermissionsPerms))
	e.POST(prefix+"/associations", h.CreateRolePermission, m.CheckPermissions(h.CreateRolePerms))
	e.GET(prefix+"/associations", h.GetAssociations, m.CheckPermissions(h.CreateRolePerms))

	// temporal routes
	e.POST(prefix+"/init_script", h.InitScript)
	e.PUT(prefix+"/dev", h.IntiDev)

}
