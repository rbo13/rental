package controllers

import (
	"log"

	"github.com/whaangbuu/home-rental/app/models"
	"gopkg.in/gin-gonic/gin.v1"
)

// ViewRequestIndex handles view of
// tenant request
func ViewRequestIndex(c *gin.Context) {
	log.Println("ViewRequestIndex")

	myOwnerID := GetMyOwnerID(c)

	records, err := models.GetRecordsByOwnerID(myOwnerID)

	if err != nil {
		SetFlashError(c, err.Error())
		Redirect(c, "/admin/home")
		return
	}

	RenderHTML(c, gin.H{
		"requestIndex": "Tenant Requests",
		"myOwnerID":    myOwnerID,
		"records":      records,
	})
}
