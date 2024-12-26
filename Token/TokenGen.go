package token

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(ClientID uint) (string, error) {

	claims := jwt.MapClaims{
		"ClientID": ClientID,
		"exp":      time.Now().Add(time.Minute * 20).Unix(),
	}
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return "", fmt.Errorf("secretKey empty")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("token signing failed")
	}
	return signedToken, nil
}
