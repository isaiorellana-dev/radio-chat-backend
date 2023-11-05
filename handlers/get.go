package handlers

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	data "github.com/isaiorellana-dev/livechat-backend/db"
	m "github.com/isaiorellana-dev/livechat-backend/models"
	"github.com/labstack/echo/v4"
)

func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func GetRoles(c echo.Context) error {

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

	var roles []m.Role

	if err := db.Find(&roles).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, objectStr{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, roles)

}

func GetPermissions(c echo.Context) error {

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

	var permissions []m.Permission

	if err := db.Find(&permissions).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, objectStr{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, permissions)

}

func GetAssociations(c echo.Context) error {

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
	var associations = []associaton{}

	if err := db.Table("role_permissions").Find(&associations).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, objectStr{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, associations)

}

func GetUsers(c echo.Context) error {
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

	var users []m.UserToReturn

	if err := db.Table("users").
		Select("users.id, users.nickname, roles.name as role, users.created_at").
		Joins("JOIN roles ON users.rol_id = roles.id").
		Scan(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, objectStr{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, users)
}

func GetOneUser(c echo.Context) error {
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

	var user m.UserToReturn

	userID := c.Param("id")

	if err := db.Select("users.id, users.nickname, roles.name as role, users.created_at").Joins("JOIN roles ON users.rol_id = roles.id").First(&m.User{}, userID).Scan(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, objectStr{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func GetMessages(c echo.Context) error {
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

	var messages []m.MessageWithUser

	if err := db.Find(&[]m.Message{}).Select("messages.id, messages.body, messages.created_at, users.nickname").Joins("JOIN users ON messages.user_id = users.id").Order("id desc").Scan(&messages).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, objectStr{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, messages)
}

func UserData(c echo.Context) error {
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

	var userData m.UserData

	db.Select("users.nickname, roles.name as role").Joins("JOIN roles ON roles.id = users.rol_id").First(&m.User{}, claims.UserID).Scan(&userData)

	return c.JSON(http.StatusOK, userData)
}
