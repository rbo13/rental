package controllers

import (
	"log"

	"github.com/whaangbuu/home-rental/app/models"
	"gopkg.in/gin-gonic/gin.v1"
)

// AdminManageIndex ...
func AdminManageIndex(c *gin.Context) {
	log.Println("AdminManageIndex")

	ownerID := GetMyOwnerID(c)
	units, err := models.GetUnitsByOwnerID(ownerID)

	if err != nil {
		SetFlashError(c, "Error getting units")
		Redirect(c, "/admin")
		return
	}

	RenderHTML(c, gin.H{
		"manageUnit": "Manage a Units",
		"units":      units,
	})
}
