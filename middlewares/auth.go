package middlewares

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	data "github.com/isaiorellana-dev/radio-chat-backend/db"
	"github.com/isaiorellana-dev/radio-chat-backend/models"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var (
	NO_AUTH_NEEDED = []string{
		"/login",
		"/signup",
	}
)

type objectStr map[string]string

func getPermissions(rolID int) ([]string, error) {
	db, err := data.ConnectToDB()
	if err != nil {
		return nil, err
	}

	defer func() {
		dbSQL, err := db.DB()
		if err != nil {
			return
		}
		dbSQL.Close()
	}()

	var permissions []string

	err = db.Table("roles").
		Select("permissions.name").
		Joins("JOIN role_permissions ON roles.id = role_permissions.role_id").
		Joins("JOIN permissions ON role_permissions.permission_id = permissions.id").
		Where("roles.id = ?", rolID).
		Scan(&permissions).Error
	if err != nil {
		return nil, err
	}

	return permissions, nil
}

func shouldCheckToken(route string) bool {
	for _, path := range NO_AUTH_NEEDED {
		if strings.Contains(route, path) {
			return false
		}
	}
	return true
}

func extractTokenFromAuthHeader(authHeader string) string {
	const prefix = "Bearer "
	if len(authHeader) > len(prefix) && authHeader[:len(prefix)] == prefix {
		return authHeader[len(prefix):]
	}
	return ""
}

func CheckJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		if !shouldCheckToken(c.Request().URL.Path) {
			return next(c)
		}

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

		c.Set("token", token)

		return next(c)
	}
}

func CheckPermissions(requiredPerms []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			var token = c.Get("token").(*jwt.Token)
			claims, _ := token.Claims.(*models.AppClaims)

			permissions, err := getPermissions(claims.RolID)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, objectStr{"error": err.Error()})
			}

			hasPermissions := false
			for _, rp := range requiredPerms {
				hasPerm := false
				for _, p := range permissions {
					if rp == p {
						hasPerm = true
						break
					}
				}
				if !hasPerm {
					hasPermissions = hasPerm
					break
				}
				hasPermissions = true
			}

			if !hasPermissions {
				return c.JSON(http.StatusForbidden, objectStr{
					"message": "you don't have the required permissions",
				})
			}

			return next(c)
		}
	}
}
