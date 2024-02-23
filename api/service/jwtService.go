package service

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kylerequez/go-gin-dependency-injection-v1/types"
)

type JwtService struct {
}

func NewJwtService() *JwtService {
	return &JwtService{}
}

func (js *JwtService) GenerateJWT(user *types.User) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	claims["authority"] = user.Authority
	claims["id"] = user.ID
	claims["email"] = user.Email

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (js *JwtService) ValidateJWT(t string) (token *jwt.Token, err error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	resultToken, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}
	return resultToken, nil
}

func (js *JwtService) ExtractJWTClaims(t *jwt.Token) (c jwt.MapClaims, err error) {
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("error in extracting claims")
	}
	return claims, nil
}
