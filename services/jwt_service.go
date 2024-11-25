package services

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/exp/rand"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(customerID int) (string, error) {
	claims := jwt.MapClaims{
		"customer_id": customerID,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return jwtSecret, nil
	})
}

func GenerateID() int {
	rand.Seed(uint64(time.Now().UnixNano()))
	return rand.Intn(100000)
}
