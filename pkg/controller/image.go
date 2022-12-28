package controller

import (
	"fmt"
	"github.com/NekruzRakhimov/tg_user_registrator/logger"
	"github.com/NekruzRakhimov/tg_user_registrator/utils"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func SaveImage(c *gin.Context, key string) (string, error) {
	file, err := c.FormFile(key)
	if err != nil {
		logger.Error.Printf("1. [%s] Error is: %s\n", utils.FuncName(), err.Error())
		return "", err
	}

	filename := fmt.Sprintf("images/%s_%s", utils.RandomString(5), file.Filename)
	err = c.SaveUploadedFile(file, filename)
	if err != nil {
		logger.Error.Printf("2. [%s] Error is: %s\n", utils.FuncName(), err.Error())
		return "", err
	}

	return filename, nil
}

func GetImage(c *gin.Context) {
	f := c.Param("path")
	fmt.Println(f)
	extension := filepath.Ext(f)
	fmt.Println(1)
	switch extension {
	case ".jpg":
		c.Writer.Header().Set("Content-Type", "image/jpeg")
	case ".png":
		c.Writer.Header().Set("Content-Type", "image/png")
	case ".pdf":
		c.Writer.Header().Set("Content-Type", "application/pdf")
	}
	c.File(fmt.Sprintf("images/%s", f))
}
