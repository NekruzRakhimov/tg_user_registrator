package controller

import (
	"github.com/NekruzRakhimov/tg_user_registrator/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//PingPong unconvicted godoc
// @Summary unconvicted
// @Description Роут для проверки работы сервера
// @Accept  json
// @Produce  json
// @Tags url
// @Success 200 {object} models.PingPong
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router / [get]
func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, models.PingPong{Pong: "Server is up and running!"})
}
