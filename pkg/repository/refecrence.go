package repository

import (
	"errors"
	"fmt"
	"github.com/NekruzRakhimov/tg_user_registrator/db"
	"github.com/NekruzRakhimov/tg_user_registrator/logger"
	"github.com/NekruzRakhimov/tg_user_registrator/models"
	"github.com/NekruzRakhimov/tg_user_registrator/utils"
)

func CreateReference(r models.Reference) error {
	if err := db.GetDBConn().Table("references").Omit("created_at").Create(&r).Error; err != nil {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), err.Error())
		return errors.New("ошибка во время создания записи")
	}

	return nil
}

func GetMyReferences(email string) (r []models.Reference, err error) {
	fmt.Println(email)
	sqlQuery := "SELECT * FROM \"references\" WHERE email = ? ORDER BY id DESC"
	if err = db.GetDBConn().Raw(sqlQuery, email).Scan(&r).Error; err != nil {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), err.Error())
		return nil, errors.New("ошибка во время получения данных")
	}

	return r, nil
}

func GetAllReferences(page, limit int, search, status, tariff string) (r []models.Reference, lastPage int, err error) {
	sqlQuery := "SELECT * FROM \"references\" WHERE true "
	if search != "" {
		sqlQuery += " AND full_name like '%" + search + "%'"
	}

	if status != "" {
		sqlQuery += fmt.Sprintf(" AND status_type = '%s'", status)
	}

	if tariff != "" {
		sqlQuery += fmt.Sprintf(" AND reference_tariff = '%s'", tariff)
	}
	if err = db.GetDBConn().Raw(sqlQuery).Scan(&r).Error; err != nil {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), err.Error())
		return nil, 0, errors.New("ошибка во время получения данных")
	}

	fmt.Println(len(r)%limit != 0)
	lastPage = len(r) / limit
	if len(r)%limit != 0 {
		lastPage++
		fmt.Println("here", lastPage)
	}

	sqlQuery += fmt.Sprintf(" ORDER BY id DESC OFFSET %d LIMIT %d", (page-1)*limit, limit)

	if err = db.GetDBConn().Raw(sqlQuery).Scan(&r).Error; err != nil {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), err.Error())
		return nil, 0, errors.New("ошибка во время получения данных")
	}
	return r, lastPage, nil
}

func GetReferenceByID(id int) (r models.Reference, err error) {
	sqlQuery := "SELECT * FROM \"references\" WHERE id = ?"
	if err = db.GetDBConn().Raw(sqlQuery, id).Scan(&r).Error; err != nil {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), err.Error())
		return models.Reference{}, errors.New("ошибка во время получения данных")
	}

	return r, nil
}

func ChangeReferenceStatus(id int, comment, statusType string) error {
	var status string
	switch statusType {
	case "on_consideration":
		status = "На рассмотрение"
	case "canceled":
		status = "Отказано"
	case "ready":
		status = "Принято"
	}

	sqlQuery := "UPDATE \"references\" set \"status\" = ?, \"comment\" = ?, \"status_type\" = ? WHERE id = ?"
	if err := db.GetDBConn().Exec(sqlQuery, status, comment, statusType, id).Error; err != nil {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), err.Error())
		return errors.New("ошибка во время получения данных")
	}

	return nil
}
