package handlers

import (
	"net/http"

	data "github.com/isaiorellana-dev/livechat-backend/db"
	"github.com/isaiorellana-dev/livechat-backend/models"
	"github.com/labstack/echo/v4"
)

func UpdateUser(c echo.Context) error {
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

	var user models.User
	var nickname = c.Get("user").(*models.User).Nickname
	userID := c.Param("id")

	if err := db.First(&user, userID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, objectStr{"error": err.Error()})
	}
	if user.Nickname == nickname {
		return c.JSON(http.StatusBadRequest, objectStr{"error": "the nickname has no changes"})
	}
	user.Nickname = nickname

	if err := db.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, objectStr{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func IntiDev(c echo.Context) error {

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

	var user models.User
	if err := db.First(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, objectStr{"error": err.Error()})
	}

	var devRole models.Role

	if err := db.Where("name = ?", DEV).First(&devRole).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, objectStr{"error": err.Error()})
	}

	user.RolID = devRole.ID
	if err := db.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, objectStr{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, objectStr{
		"message": "dev ready",
	})
}
