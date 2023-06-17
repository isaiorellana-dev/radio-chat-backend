package middlewares

import (
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	data "github.com/isaiorellana-dev/radio-chat-backend/db"
	"github.com/isaiorellana-dev/radio-chat-backend/models"
	"github.com/labstack/echo/v4"
)

func ValidateUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user = new(models.User)

		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Invalid request body",
				"error":   err.Error(),
			})
		}

		v := validator.New()
		if err := v.Struct(user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Invalid request body",
				"error":   err.Error(),
			})
		}

		c.Set("user", user)

		return next(c)
	}
}

func ValidateUserByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var message = new(models.Message)

		if err := c.Bind(message); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

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

		if err := db.First(&models.User{}, message.UserID).Error; err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
		}

		c.Set("message", message)

		return next(c)
	}
}

func ValidateMessage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var message = c.Get("message").(*models.Message)

		trimmedStr := strings.TrimSpace(message.Body)
		message.Body = trimmedStr

		v := validator.New()
		if err := v.Struct(message); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Invalid request body",
				"error":   err.Error(),
			})
		}

		c.Set("message", message)

		return next(c)
	}
}
