package handlers_test

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	utils_test "github.com/isaiorellana-dev/radio-chat-backend/__tests__/utils"
	"github.com/isaiorellana-dev/radio-chat-backend/handlers"
	"github.com/isaiorellana-dev/radio-chat-backend/models"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func init() {
	// Cargar las variables de entorno de .env.test
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatal("Error loading .env.test file")
	}
}
func TestGetMessagesHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/messages", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handlers.GetMessages(c)

	var messagesWithUser []models.MessageWithUser

	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)

	assert.Equal(t, "application/json; charset=UTF-8", rec.Header().Get("Content-Type"))
	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &messagesWithUser))

}

func TestGetUsers(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handlers.GetUsers(c)

	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUserData(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	user := utils_test.CreateUser()

	c.Set("user", user)

	err := handlers.UserData(c)

	assert.NoError(t, err)

}
