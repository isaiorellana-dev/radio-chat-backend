package handlers

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	data "github.com/isaiorellana-dev/radio-chat-backend/db"
	"github.com/isaiorellana-dev/radio-chat-backend/models"
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

	var token = c.Get("token").(*jwt.Token)
	claims, _ := token.Claims.(*models.AppClaims)

	fmt.Println(claims)

	var role models.Role
	var permissions []models.Permission

	// Buscar el rol por ID y cargar los permisos relacionados
	db.Preload("Permissions").First(&role, claims.RolID)

	// Obtener los permisos del rol
	fmt.Println(permissions)

	var permission = new(models.Permission)
	if err := db.First(&permission).Where("").Error; err != nil {
		return c.JSON(http.StatusInternalServerError, objectStr{
			"error": err.Error(),
		})
	}

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
