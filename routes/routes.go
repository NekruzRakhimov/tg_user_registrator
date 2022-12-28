package routes

import (
	"fmt"
	"github.com/NekruzRakhimov/tg_user_registrator/docs"
	"github.com/NekruzRakhimov/tg_user_registrator/pkg/controller"
	"github.com/NekruzRakhimov/tg_user_registrator/utils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

func RunAllRoutes() {
	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	//r.Use(cors.Default())

	r.Use(controller.CORSMiddleware())

	// Статус код 500, при любых panic()
	r.Use(gin.Recovery())

	// Запуск роутов
	initAllRoutes(r)

	// Запуск сервера
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = utils.AppSettings.AppParams.PortRun
	}
	_ = r.Run(fmt.Sprintf(":%s", port))
}

func initAllRoutes(r *gin.Engine) {
	r.GET("/", controller.PingPong)
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//r.POST("/image", controller.SaveImage)
	r.GET("/image", controller.GetImage)

	api := r.Group("/api")
	api.GET("/images/:path", controller.GetImage)
	api.POST("/auth/sign-up", controller.SignUp)
	api.POST("/auth/sign-in", controller.SignIn)
	api.GET("/auth/me", controller.UserIdentity, controller.GetMe)
	api.PUT("/auth/me", controller.UserIdentity, controller.EditProfileInfo)
	api.POST("/reference", controller.CreateReference)

	api.GET("/reference", controller.UserIdentity, controller.GetMyReferences)
	api.GET("/reference/:id", controller.UserIdentity, controller.GetReferenceByID)
	api.PUT("/reference/:id", controller.UserIdentity, controller.ChangeReferenceStatus)

	admin := api.Group("/admin")
	admin.GET("/activity", controller.GetAdminsActivity)
	admin.GET("/reference", controller.UserIdentity, controller.GetAllReferences)
	admin.POST("/reference/:id/template", controller.UserIdentity, controller.GetReferenceTemplate)
	admin.GET("/", controller.UserIdentity, controller.GetAllAdmins)
	admin.GET("/users", controller.UserIdentity, controller.GetAllAdmins)

	admin.POST("/", controller.UserIdentity, controller.CreateAdmin)
	admin.POST("/users", controller.UserIdentity, controller.CreateAdmin)

	admin.DELETE("/users/:id", controller.UserIdentity, controller.DeleteAdmin)
}
