package routes

import (
	"fmt"

	// "github.com/isaiorellana-dev/radio-chat-backend/handlers"
	"github.com/isaiorellana-dev/radio-chat-backend/context"
	"github.com/isaiorellana-dev/radio-chat-backend/handlers"
	m "github.com/isaiorellana-dev/radio-chat-backend/middlewares"
	ws "github.com/isaiorellana-dev/radio-chat-backend/websocket"
	"github.com/labstack/echo/v4"
)

func CustomContextMiddleware(hub *ws.Hub) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &context.CustomContext{
				Context: c,
				Hub:     hub,
			}
			return next(cc)
		}
	}
}

func RegisterRoutes(e *echo.Echo) {

	hub := ws.NewHub()
	go hub.Run()

	e.Use(m.CheckJWT)

	e.Use(CustomContextMiddleware(hub))

	e.GET("/ws", func(c echo.Context) error {
		fmt.Println("Hola desde el hanlder")
		return ws.ServeWs(hub, c)
	})

	// Get methods
	e.GET("/api/v1/hello", handlers.HelloWorld)
	e.GET("/api/v1/users", handlers.GetUsers, m.CheckPermissions(handlers.GetUsersPerms))
	e.GET("/api/v1/users/:id", handlers.GetOneUser)
	e.GET("/api/v1/messages", handlers.GetMessages)

	// Post methods
	e.POST("/api/v1/signup", handlers.Register, m.ValidateUser)
	e.POST("/api/v1/messages", handlers.CreateMessage,
		// m.CheckPermissions(handlers.CreateMessagePerms),
		m.ValidateUserByID,
		m.ValidateMessage)
	e.POST("/api/v1/login", handlers.Login)
	e.POST("/api/v1/role", handlers.CreateRole, m.CheckPermissions(handlers.CreateRolePerms))
	e.POST("/api/v1/permission", handlers.CreatePermission, m.CheckPermissions(handlers.CreatePermissionsPerms))

	// Put methods
	e.PUT("/api/v1/users/:id", handlers.UpdateUser, m.ValidateUser)

	// Delete methods
	e.DELETE("api/v1/users/:id", handlers.DeleteUser, m.CheckPermissions(handlers.DeleteUserPerms))
	e.DELETE("/api/v1/messages/:id", handlers.DeleteMessage, m.CheckPermissions(handlers.DeleteMessagePerms))
}
