package middlewares

import (
	"github.com/isaiorellana-dev/radio-chat-backend/context"
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
