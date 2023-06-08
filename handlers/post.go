package handlers

import (
	"net/http"

	data "github.com/isaiorellana-dev/radio-api/db"
	"github.com/isaiorellana-dev/radio-api/models"
	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	db, err := data.ConnectToDB()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	defer func() {
		dbSQL, err := db.DB()
		if err != nil {
			return
		}
		dbSQL.Close()
	}()

	var user = c.Get("user").(*models.User)

	if err := db.Create(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, user)
}
