package middlewares

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/isaiorellana-dev/radio-chat-backend/models"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// var (
// 	NO_AUTH_NEEDED = []string{
// 		"login", "register",
// 	}
// )

// func shouldCheckToken(route string) bool {
// 	for _, path := range NO_AUTH_NEEDED {
// 		if strings.Contains(route, path) {
// 			return false
// 		}
// 	}
// 	return true
// }

func extractTokenFromAuthHeader(authHeader string) string {
	const prefix = "Bearer "
	if len(authHeader) > len(prefix) && authHeader[:len(prefix)] == prefix {
		return authHeader[len(prefix):]
	}
	return ""
}

func CheckJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal("Error loading .env file")
		}
		JWT_SECRET := os.Getenv("JWT_SECRET")

		authHeader := strings.TrimSpace(c.Request().Header.Get("Authorization"))
		tokenStr := extractTokenFromAuthHeader(authHeader)

		token, err := jwt.ParseWithClaims(tokenStr, &models.AppClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(JWT_SECRET), nil
		})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "invalid token",
			})
		}

		if !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "invalid token",
			})
		}

		return next(c)
	}
}
