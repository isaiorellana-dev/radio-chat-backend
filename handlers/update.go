package handlers

import (
	"net/http"

	data "github.com/isaiorellana-dev/radio-chat-backend/db"
	"github.com/labstack/echo/v4"
)

func SetAdmin(c echo.Context) error {
	db, err := data.ConnectToDB()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, objectStr{
			"error": err.Error(),
		})
	}

	defer func() {
		dbSQL, err := db.DB()
		if err != nil {
			return
		}
		dbSQL.Close()
	}()

	// var user models.User

	if err := db.Exec("UPDATE users SET rol_id = 2 WHERE id=6").Error; err != nil {
		return c.JSON(http.StatusInternalServerError, objectStr{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, objectStr{
		"message": "admin set",
	})

}
