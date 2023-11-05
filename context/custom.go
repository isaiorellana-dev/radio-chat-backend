package context

import (
	ws "github.com/isaiorellana-dev/livechat-backend/websocket"
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	Hub *ws.Hub
}
