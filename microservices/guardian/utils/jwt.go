package utils

import (
	"errors"
	"guardian/api/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId int64, displayName string, email string, roles string) (string, error) {
	//generate token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":      userId,
		"displayName": displayName,
		"email":       email,
		"roles":       roles,
		"exp":         time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(config.Config["jwtSecretKey"]))
}

func VerifyToken(token string) (userId int64, roles string, Error error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // 187 course how was the token signed

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(config.Config["jwtSecretKey"]), nil
	})

	if err != nil {
		return 0, "", errors.New("could not parse token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, "", errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, "", errors.New("invalid token claims")
	}

	// email := claims["email"].(string)
	userId = int64(claims["userId"].(float64))
	roles = string(claims["roles"].(string))

	return userId, roles, nil
}
