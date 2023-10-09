package helpers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(email string, id int) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(1 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return err.Error()
	}
	return tokenStr

}
