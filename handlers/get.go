package handlers

import (
	"net/http"

	data "github.com/isaiorellana-dev/radio-api/db"
	"github.com/isaiorellana-dev/radio-api/models"
	m "github.com/isaiorellana-dev/radio-api/models"
	"github.com/labstack/echo/v4"
)

func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func GetUsers(c echo.Context) error {
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

	var users []m.User

	if err := db.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, users)
}

func GetOneUser(c echo.Context) error {
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

	userID := c.Param("id")

	if err := db.First(&user, userID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func GetMessages(c echo.Context) error {
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

	var messages []m.MessageWithUser

	if err := db.Model(&m.Message{}).Select("messages.id, messages.body, messages.user_id, users.nickname").Joins("JOIN users ON messages.user_id = users.id").Scan(&messages).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, messages)
}
