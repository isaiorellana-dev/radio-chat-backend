package handlers

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	data "github.com/isaiorellana-dev/radio-chat-backend/db"
	m "github.com/isaiorellana-dev/radio-chat-backend/models"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) error {
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

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	HASH := os.Getenv("HASH_COST")
	HASH_COST, err := strconv.Atoi(HASH)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"eror": "Invalid Hash Cost",
		})
	}

	var user = c.Get("user").(*m.User)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Pin), HASH_COST)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	user.Pin = string(hashedPassword)

	userRegistered := m.UserRegister{}

	if err := db.Create(&user).Scan(&userRegistered).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, userRegistered)
}

func Login(c echo.Context) error {
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

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	JWT_SECRET := os.Getenv("JWT_SECRET")

	var login = new(m.UserLogin)

	if err := c.Bind(login); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	var user = new(m.User)

	if err := db.First(&m.User{}, "nickname = ?", login.Nickname).Scan(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "invalid credentials",
		})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Pin), []byte(login.Pin)); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "invalid credentials",
		})
	}

	claims := m.AppClaims{
		UserID:   user.ID,
		Nickname: user.Nickname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "error en servidor",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})
}

func CreateMessage(c echo.Context) error {
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

	var message = c.Get("message").(*m.Message)

	if err := db.Create(&message).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, message)
}
