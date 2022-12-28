package service

import (
	"errors"
	"github.com/NekruzRakhimov/tg_user_registrator/logger"
	"github.com/NekruzRakhimov/tg_user_registrator/models"
	"github.com/NekruzRakhimov/tg_user_registrator/pkg/repository"
	"github.com/NekruzRakhimov/tg_user_registrator/utils"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

func CreateUser(user models.User) error {
	u, err := repository.GetUserByEmail(user.Email)
	if u.ID != 0 {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), "пользователь с такой почтой уже существует")
		return errors.New("пользователь с такой почтой уже существует")
	}

	user.Password = utils.GenerateHash(user.Password)
	if err = repository.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func GetUserByID(id int) (models.User, error) {
	return repository.GetUserByID(id)
}

func GetUserByEmail(email string) (models.User, error) {
	return repository.GetUserByEmail(email)
}

func EditProfileInfo(id int, user models.User) error {
	user.ID = id
	user.Password = utils.GenerateHash(user.Password)
	return repository.EditProfileInfo(user)
}
