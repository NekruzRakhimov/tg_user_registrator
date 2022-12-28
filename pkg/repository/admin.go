package repository

import (
	"errors"
	"fmt"
	"github.com/NekruzRakhimov/tg_user_registrator/db"
	"github.com/NekruzRakhimov/tg_user_registrator/logger"
	"github.com/NekruzRakhimov/tg_user_registrator/models"
	"github.com/NekruzRakhimov/tg_user_registrator/utils"
)

func GetAllAdmins() (a []models.User, err error) {
	sqlQuery := "SELECT * FROM users WHERE is_admin = true and is_removed = false"
	if err = db.GetDBConn().Raw(sqlQuery).Scan(&a).Error; err != nil {
		return nil, err
	}

	return
}

func CreateNewAdmin(u models.User) (err error) {
	u.IsSuperAdmin = false
	u.IsAdmin = true
	if err = db.GetDBConn().Table("users").Omit("old_password").Create(&u).Error; err != nil {
		return err
	}

	return nil
}

func DeleteAdmin(id int) error {
	sqlQuery := "UPDATE users set is_removed = true WHERE id = ?"
	if err := db.GetDBConn().Exec(sqlQuery, id).Error; err != nil {
		return err
	}

	return nil
}

func GetAdminsActivity(page, limit int, search string) (activity []models.AdminActivity, lastPage int, err error) {
	sqlQuery := "SELECT * FROM admins_activity WHERE true"
	if search != "" {
		sqlQuery += " AND full_name like '%" + search + "%'"
	}

	if err = db.GetDBConn().Raw(sqlQuery).Scan(&activity).Error; err != nil {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), err.Error())
		return nil, 0, errors.New("ошибка во время получения данных")
	}

	lastPage = len(activity) / limit
	if len(activity)%limit != 0 {
		lastPage++
		fmt.Println("here", lastPage)
	}

	sqlQuery += fmt.Sprintf(" ORDER BY id DESC OFFSET %d LIMIT %d", (page-1)*limit, limit)

	if err = db.GetDBConn().Raw(sqlQuery).Scan(&activity).Error; err != nil {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), err.Error())
		return nil, 0, errors.New("ошибка во время получения данных")
	}

	return activity, lastPage, nil
}
