package service

import (
	"errors"
	"github.com/NekruzRakhimov/tg_user_registrator/models"
	"github.com/NekruzRakhimov/tg_user_registrator/pkg/repository"
	"github.com/NekruzRakhimov/tg_user_registrator/utils"
)

func GetAllAdmins() (a []models.User, err error) {
	return repository.GetAllAdmins()
}

func CreateNewAdmin(admin models.User) (err error) {
	u, err := repository.GetUserByEmail(admin.Email)
	if err != nil {
		return err
	}

	if u.ID != 0 {
		return errors.New("пользователь с этим email'ом уже существует")
	}

	admin.Password = utils.GenerateHash(admin.Password)
	return repository.CreateNewAdmin(admin)
}

func DeleteAdmin(id int) error {
	return repository.DeleteAdmin(id)
}

func GetAdminsActivity(page, limit int, search string) (activity []models.AdminActivity, lastPage int, err error) {
	return repository.GetAdminsActivity(page, limit, search)
}
