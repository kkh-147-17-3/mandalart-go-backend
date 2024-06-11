package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}


func CreateToken(userId int) (*Tokens, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.Itoa(userId),
		"token_type": "access_token",
		"iss": "mandalart.com",
		"exp": time.Now().Add(time.Hour * 100000).Unix(),
		"iat": time.Now().Unix(),
	})

	accessTokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	claims = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.Itoa(userId),
		"token_type": "refresh_token",
		"iss": "mandalart.com",
		"exp": time.Now().Add(time.Hour * 100000).Unix(),
		"iat": time.Now().Unix(),
	})

	refreshTokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return &Tokens{
		AccessToken: accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	// Check for verification errors
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Return the verified token
	return token, nil
}

func GetUserIdFromToken(tokenString string) (string, error) {
	token, err := verifyToken(tokenString)
	if err != nil {
		return "", err
	}
	sub, err := token.Claims.GetSubject()
	if err != nil {
		return "", err
	}
	return sub, nil
}
