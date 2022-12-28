package controller

import (
	"github.com/NekruzRakhimov/tg_user_registrator/logger"
	"github.com/NekruzRakhimov/tg_user_registrator/models"
	"github.com/NekruzRakhimov/tg_user_registrator/pkg/service"
	"github.com/NekruzRakhimov/tg_user_registrator/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SignUp(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": "ошибка во время парсинга данных"})
		return
	}

	if err := service.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	token, err := service.GenerateToken(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": token})
}

func SignIn(c *gin.Context) {
	adminPanel, _ := strconv.ParseBool(c.Query("admin_panel"))

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": "ошибка во время парсинга данных"})
		return
	}

	if adminPanel {
		u, err := service.GetUserByEmail(user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		if !u.IsAdmin && !u.IsSuperAdmin {
			c.JSON(http.StatusForbidden, gin.H{"reason": "неправильный логин или пароль"})
			return
		}
	}

	token, err := service.GenerateToken(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": token})
}
