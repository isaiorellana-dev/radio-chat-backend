package handlers

import (
	"net/http"

	data "github.com/isaiorellana-dev/radio-api/db"
	"github.com/isaiorellana-dev/radio-api/models"
	"github.com/labstack/echo/v4"
)

func DeleteUser(c echo.Context) error {
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

	userId := c.Param("id")

	if err := db.Delete(&models.User{}, userId).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "User deleted",
	})
}
