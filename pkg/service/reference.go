package service

import (
	"github.com/NekruzRakhimov/tg_user_registrator/logger"
	"github.com/NekruzRakhimov/tg_user_registrator/models"
	"github.com/NekruzRakhimov/tg_user_registrator/pkg/repository"
)

func CreateReference(r models.Reference) error {
	r.Status = "На рассмотрении"
	r.StatusType = "on_consideration"
	return repository.CreateReference(r)
}

func GetMyReference(userID int) (r []models.Reference, err error) {
	u, err := repository.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	logger.Debug.Println(u)

	r, err = repository.GetMyReferences(u.Email)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func GetAllReferences(page, limit int, search, status, tariff string) (r []models.Reference, lastPage int, err error) {
	r, lastPage, err = repository.GetAllReferences(page, limit, search, status, tariff)
	if err != nil {
		return nil, 0, err
	}

	return r, lastPage, nil
}

func GetReferenceByID(id int) (r models.Reference, err error) {
	return repository.GetReferenceByID(id)
}

func ChangeReferenceStatus(id int, status string, comment string) error {
	return repository.ChangeReferenceStatus(id, comment, status)
}
