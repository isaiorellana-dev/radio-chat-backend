package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	data "github.com/isaiorellana-dev/radio-chat-backend/db"
	m "github.com/isaiorellana-dev/radio-chat-backend/models"

	"github.com/isaiorellana-dev/radio-chat-backend/context"

	// "github.com/isaiorellana-dev/radio-chat-backend/routes"
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

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
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

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
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
	fmt.Println(user)

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
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	fmt.Println(claims, "hellllo")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, objectStr{
			"message": "error en servidor",
		})
	}

	fmt.Println(c.Request().URL.Path, claims)

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

	cc := c.(*context.CustomContext)
	hub := cc.Hub

	var message = c.Get("message").(*m.Message)

	if err := db.Create(&message).Error; err != nil {
		return c.JSON(http.StatusBadRequest, objectStr{
			"error": err.Error(),
		})
	}

	hub.Messages <- message

	// hub := ws.G

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

	var Role = new(m.Role)

	if err := c.Bind(&Role); err != nil {
		return c.JSON(http.StatusBadRequest, objectStr{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if err := db.Create(&Role).Error; err != nil {
		return c.JSON(http.StatusBadRequest, objectStr{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, Role)
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

	var Permission = new(m.Permission)

	if err := c.Bind(&Permission); err != nil {
		return c.JSON(http.StatusBadRequest, objectStr{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if err := db.Create(&Permission).Error; err != nil {
		return c.JSON(http.StatusBadRequest, objectStr{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, Permission)
}
