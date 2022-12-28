package controller

import (
	"encoding/json"
	"github.com/NekruzRakhimov/tg_user_registrator/logger"
	"github.com/NekruzRakhimov/tg_user_registrator/models"
	"github.com/NekruzRakhimov/tg_user_registrator/pkg/service"
	"github.com/NekruzRakhimov/tg_user_registrator/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateReference(c *gin.Context) {
	var (
		err       error
		reference models.Reference
	)
	str, ok := c.GetPostForm("json")
	if !ok {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), "json field not found1")
		c.JSON(http.StatusBadRequest, gin.H{"reason": "json field not found"})
		return
	}

	if err = json.Unmarshal([]byte(str), &reference); err != nil {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": "error while unmarshalling"})
		return
	}

	reference.PassportFront, err = SaveImage(c, "passport_front")
	if err != nil {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": "passport_front field not found"})
		return
	}

	reference.PassportBack, err = SaveImage(c, "passport_back")
	if err != nil {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": "passport_back field not found"})
		return
	}

	reference.PassportSelfie, err = SaveImage(c, "passport_selfie")
	if err != nil {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": "passport_selfie field not found"})
		return
	}

	reference.PaymentReceipt, err = SaveImage(c, "payment_receipt")
	if err != nil {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": "payment_receipt field not found"})
		return
	}

	if err = service.CreateReference(reference); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "Запись успешно создана"})
}

func GetMyReferences(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": err.Error()})
		return
	}

	logger.Debug.Println(userID)

	r, err := service.GetMyReference(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, r)
}

func GetReferenceByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "id not found"})
		return
	}

	r, err := service.GetReferenceByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, r)
}

func ChangeReferenceStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "id not found"})
		return
	}

	var refStatus models.ReferenceStatus
	if err = c.BindJSON(&refStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err = service.ChangeReferenceStatus(id, refStatus.Status, refStatus.Comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "статус успешно изменен"})
}

func GetReferenceTemplate(c *gin.Context) {
	_, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "id not found"})
		return
	}

	c.File("./files/template.doc")
}

func GetAllReferences(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": err.Error()})
		return
	}

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
	status := c.Query("status")
	tariff := c.Query("tariff")

	user, err := service.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	if !user.IsAdmin || !user.IsSuperAdmin {
		c.JSON(http.StatusForbidden, gin.H{"reason": "у вас нет необходимых прав"})
		return
	}

	logger.Debug.Println(userID)

	r, lastPage, err := service.GetAllReferences(page, limit, search, status, tariff)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"references": r,
		"last_page":  lastPage,
	})
}
