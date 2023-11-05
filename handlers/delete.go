package handlers

import (
	"net/http"

	data "github.com/isaiorellana-dev/livechat-backend/db"
	"github.com/isaiorellana-dev/livechat-backend/models"
	"github.com/labstack/echo/v4"
)

func DeleteUser(c echo.Context) error {
	db, err := data.ConnectToDB()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, objectStr{"error": err.Error()})
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
		return c.JSON(http.StatusInternalServerError, objectStr{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, objectStr{
		"message": "User deleted",
	})
}

func DeleteMessage(c echo.Context) error {
	db, err := data.ConnectToDB()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, objectStr{"error": err.Error()})
	}

	defer func() {
		dbSQL, err := db.DB()
		if err != nil {
			return
		}
		dbSQL.Close()
	}()

	messageId := c.Param("id")

	if err := db.Delete(&models.Message{}, messageId).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, objectStr{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, objectStr{
		"message": "Message deleted",
	})
}
