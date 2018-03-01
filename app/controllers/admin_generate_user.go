package controllers

import (
	"log"

	"github.com/whaangbuu/home-rental/app/models"
	"gopkg.in/gin-gonic/gin.v1"
)

// AdminGenerateUser generates fake data
func AdminGenerateUser(c *gin.Context) {
	log.Println("AdminGenerate")
	for index := 0; index < 5; index++ {
		models.InsertFakeData()
	}

	SetFlashInfo(c, "User generated")
	Redirect(c, "/admin")
}
