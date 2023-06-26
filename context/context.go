package context

import (
	ws "github.com/isaiorellana-dev/radio-chat-backend/websocket"
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	Hub *ws.Hub
}
