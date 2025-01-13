package utils

import (
	"time"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(id uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": id,
		"exp" : time.Now().Add(time.Hour * 72).Unix(),
		"role" : role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseJwt(tokenStr string)(uint, string, error){
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return 0, "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
		userId := uint(claims["user_id"].(float64))
		role := claims["role"].(string)
		return userId, role, nil
	}
	return 0, "", nil
}