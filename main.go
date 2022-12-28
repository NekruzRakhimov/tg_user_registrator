package main

import (
	"github.com/NekruzRakhimov/tg_user_registrator/db"
	"github.com/NekruzRakhimov/tg_user_registrator/logger"
	"github.com/NekruzRakhimov/tg_user_registrator/routes"
	"github.com/NekruzRakhimov/tg_user_registrator/utils"
)

// @title Gin Swagger Unconvicted Api
// @version 1.0
// @description Проверка несудимости.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email nekruzrakhimov@icloud.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host https://unconvicted.herokuapp.com/
// @BasePath /
// @schemes http
func main() {
	utils.ReadSettings()
	utils.PutAdditionalSettings()
	logger.Init()
	db.StartDbConnection()
	//go jobs.StartJobs()
	routes.RunAllRoutes()
}
