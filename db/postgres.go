package db

import (
	"fmt"
	"github.com/NekruzRakhimov/tg_user_registrator/utils"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var database *gorm.DB

func initDB() *gorm.DB {
	settingParams := utils.AppSettings.PostgresParams

	connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		settingParams.Server, settingParams.Port,
		settingParams.User, settingParams.DataBase,
		settingParams.Password)
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})

	if err != nil {
		log.Fatal("Couldn't connect to database", err.Error())
	}

	return db
}

// StartDbConnection Creates connection to database
func StartDbConnection() {
	database = initDB()
}

// GetDBConn func for getting db conn globally
func GetDBConn() *gorm.DB {
	return database
}
