package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/isaiorellana-dev/livechat-backend/context"
	data "github.com/isaiorellana-dev/livechat-backend/db"
	m "github.com/isaiorellana-dev/livechat-backend/models"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) error {
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

	if os.Getenv("ENVIRONMENT") == "development" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	HASH := os.Getenv("HASH_COST")
	HASH_COST, err := strconv.Atoi(HASH)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, objectStr{
			"eror": "Invalid Hash Cost",
		})
	}

	var user = c.Get("user").(*m.User)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Pin), HASH_COST)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, objectStr{
			"error": err.Error(),
		})
	}
	user.Pin = string(hashedPassword)

	userRegistered := m.UserRegister{}

	if err := db.Create(&user).Scan(&userRegistered).Error; err != nil {
		return c.JSON(http.StatusBadRequest, objectStr{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, userRegistered)
}

func Login(c echo.Context) error {
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

	if os.Getenv("ENVIRONMENT") == "development" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	JWT_SECRET := os.Getenv("JWT_SECRET")

	var login = new(m.UserLogin)

	if err := c.Bind(login); err != nil {
		return c.JSON(http.StatusBadRequest, objectStr{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	var user = new(m.User)

	if err := db.Table("users").
		Select("users.id, users.nickname, users.pin, users.rol_id").Where("nickname = ?", login.Nickname).
		Scan(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, objectStr{
			"message": "invalid credentials",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Pin), []byte(login.Pin)); err != nil {
		return c.JSON(http.StatusUnauthorized, objectStr{
			"message": "invalid credentials",
		})
	}

	claims := m.AppClaims{
		UserID:   user.ID,
		Nickname: user.Nickname,
		RolID:    user.RolID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * (24 * 30)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, objectStr{
			"message": "error en servidor",
		})
	}

	return c.JSON(http.StatusOK, objectStr{
		"token": tokenString,
	})
}

func CreateMessage(c echo.Context) error {
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

	claims, _ := token.Claims.(*m.AppClaims)

	cc := c.(*context.CustomContext)
	hub := cc.Hub

	var message = c.Get("message").(*m.Message)

	message.UserID = claims.UserID

	if err := db.Create(&message).Error; err != nil {
		return c.JSON(http.StatusBadRequest, objectStr{
			"error": err.Error(),
		})
	}

	var messageWithUser = m.MessageWithUser{
		ID:        message.ID,
		Nickname:  claims.Nickname,
		Body:      message.Body,
		CreatedAt: message.CreatedAt,
	}

	hub.Messages <- &messageWithUser

	return c.JSON(http.StatusOK, message)
}

func CreateRole(c echo.Context) error {
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

	var role = new(m.Role)

	if err := c.Bind(&role); err != nil {
		return c.JSON(http.StatusBadRequest, objectStr{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if err := db.Create(&role).Error; err != nil {
		return c.JSON(http.StatusBadRequest, objectStr{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, role)
}

func CreatePermission(c echo.Context) error {
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

	var permission = new(m.Permission)

	if err := c.Bind(&permission); err != nil {
		return c.JSON(http.StatusBadRequest, objectStr{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if err := db.Create(&permission).Error; err != nil {
		return c.JSON(http.StatusBadRequest, objectStr{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, permission)
}

func CreateRolePermission(c echo.Context) error {
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

	type associaton struct {
		Role_id       int `json:"role_id"`
		Permission_id int `json:"permission_id"`
	}

	var ass = new(associaton)

	if err := c.Bind(&ass); err != nil {
		return c.JSON(http.StatusBadRequest, objectStr{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	fmt.Println("Role permission:", ass)

	if err := db.Table("role_permissions").Create(ass).Error; err != nil {
		return c.JSON(http.StatusBadRequest, objectStr{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, ass)
}

func InitScript(c echo.Context) error {
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

	permissions := []*m.Permission{
		&postMessages,
		&deleteMessages,
		&deleteUsers,
		&viewListOfUsers,
		&viewListOfRoles,
		&viewListOfPermissions,
		&assignRoles,
		&createRoles,
		&createPermissions,
	}
	for _, p := range permissions {
		if err := db.Create(&p).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
	}

	roles := []*m.Role{
		&guest,
		&admin,
		&dev,
	}
	for _, r := range roles {
		if err := db.Create(&r).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
	}

	if err := db.Model(&admin).Association("Permissions").Append(
		&postMessages,
		&deleteMessages,
		&deleteUsers,
		&viewListOfUsers,
	); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := db.Model(&guest).Association("Permissions").Append(
		&postMessages,
	); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := db.Model(&dev).Association("Permissions").Append(
		&postMessages,
		&deleteMessages,
		&deleteUsers,
		&viewListOfUsers,
		&viewListOfRoles,
		&assignRoles,
		&createRoles,
		&createPermissions,
	); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, objectStr{
		"message": "script done",
	})
}
