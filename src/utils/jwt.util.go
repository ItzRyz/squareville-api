package utils

import (
	"fmt"

	"github.com/ItzRyz/squareville-api/src/lib"
	"github.com/golang-jwt/jwt"
)

var SecretKey = lib.GetEnvVar("JWT_SECRET")

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	// Create Handler
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webtoken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	return webtoken, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	// Parsing token string
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}
	claims, isOk := token.Claims.(jwt.MapClaims)
	if isOk && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
