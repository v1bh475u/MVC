package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/v1bh475u/LibMan_MVC/pkg/types"
)

func CreateToken(user types.User) (string, error) {
	secret_key := os.Getenv("SECRET_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": user.Username,
			"role":     user.Role,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString([]byte(secret_key))
	if err != nil {
		return "", err
	} else {
		return tokenString, nil
	}
}

func VerifyToken(tokenString string) (string, string, error) {
	secret_key := os.Getenv("SECRET_KEY")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret_key), nil
	})

	if err != nil {
		return "", "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		role := claims["role"].(string)
		return username, role, nil
	} else {
		return "", "", err
	}
}
