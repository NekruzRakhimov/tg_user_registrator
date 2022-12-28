package controller

import (
	"github.com/NekruzRakhimov/tg_user_registrator/models"
	"github.com/NekruzRakhimov/tg_user_registrator/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllAdmins(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": err.Error()})
		return
	}

	user, err := service.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	if !user.IsSuperAdmin {
		c.JSON(http.StatusForbidden, gin.H{"reason": "у вас нет необходимых прав"})
		return
	}

	a, err := service.GetAllAdmins()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, a)
}

func CreateAdmin(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": err.Error()})
		return
	}

	sAdmin, err := service.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	if !sAdmin.IsSuperAdmin {
		c.JSON(http.StatusForbidden, gin.H{"reason": "у вас нет необходимых прав"})
		return
	}

	var u models.User
	if err = c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}
	u.Password = "admin"

	if err = service.CreateNewAdmin(u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "новый модератор успешно создан"})
}

func DeleteAdmin(c *gin.Context) {
	_, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "id not found"})
		return
	}

	if err = service.DeleteAdmin(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "пользователь успешно удален"})
}

func GetAdminsActivity(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	limit, err := strconv.Atoi(c.Query("per_page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	search := c.Query("search")

	activity, lastPage, err := service.GetAdminsActivity(page, limit, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"activity":  activity,
		"last_page": lastPage,
	},
	)
}
