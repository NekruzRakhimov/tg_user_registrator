package service

import (
	"errors"
	"github.com/NekruzRakhimov/tg_user_registrator/logger"
	"github.com/NekruzRakhimov/tg_user_registrator/pkg/repository"
	"github.com/NekruzRakhimov/tg_user_registrator/utils"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(email, password string) (string, error) {
	user, err := repository.GetUser(email, utils.GenerateHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Server",
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), "invalid signing method")
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), "token claims are not type of *tokenClaims")
		return 0, errors.New("token claims are not type of *tokenClaims")
	}

	return claims.UserID, nil
}
