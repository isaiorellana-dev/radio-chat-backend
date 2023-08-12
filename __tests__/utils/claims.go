package utils_test

import (
	"github.com/golang-jwt/jwt"
	"github.com/isaiorellana-dev/radio-chat-backend/models"
)

func CreateUser() (t string) {
	claims := models.AppClaims{
		UserID:   4,
		Nickname: "usergeneric",
		RolID:    2,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodNone, claims)

	tokenString, _ := token.SignedString(jwt.UnsafeAllowNoneSignatureType)

	return tokenString
}
