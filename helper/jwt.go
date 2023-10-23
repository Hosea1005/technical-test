package helper

import (
	"fmt"
	"service-auth/internal/config"

	"github.com/golang-jwt/jwt/v4"
)

func ParseJWTToken(tokenString string) (*config.JWTClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &config.JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return config.JWT_KEY, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*config.JWTClaim)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	return claims, nil
}
