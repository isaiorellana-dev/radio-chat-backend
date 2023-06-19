package handlers

import (
	"net/http"

	data "github.com/isaiorellana-dev/radio-chat-backend/db"
	m "github.com/isaiorellana-dev/radio-chat-backend/models"
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

	var users []m.UserToReturn

	if err := db.Table("users").
		Select("users.id, users.nickname, roles.name as role, users.created_at").
		Joins("JOIN roles ON users.rol_id = roles.id").
		Scan(&users).Error; err != nil {
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

	var user m.UserToReturn

	userID := c.Param("id")

	if err := db.Select("id, nickname, created_at, updated_at").First(&m.User{}, userID).Scan(&user).Error; err != nil {
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

	if err := db.Find(&[]m.Message{}).Select("messages.id, messages.body, messages.created_at, users.nickname").Joins("JOIN users ON messages.user_id = users.id").Scan(&messages).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, messages)
}
