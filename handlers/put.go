package handlers

import (
	"net/http"

	data "github.com/isaiorellana-dev/radio-chat-backend/db"
	"github.com/isaiorellana-dev/radio-chat-backend/models"
	"github.com/labstack/echo/v4"
)

func UpdateUser(c echo.Context) error {
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

	var user models.User
	var nickname = c.Get("user").(*models.User).Nickname
	userID := c.Param("id")

	if err := db.First(&user, userID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if user.Nickname == nickname {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "the nickname has no changes"})
	}
	user.Nickname = nickname

	if err := db.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}
