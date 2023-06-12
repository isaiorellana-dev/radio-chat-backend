package middlewares

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/isaiorellana-dev/radio-api/models"
	"github.com/labstack/echo/v4"
)

func ValidateUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user = new(models.User)

		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Invalid request body",
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
