package controllers

import (
	"log"

	"github.com/whaangbuu/home-rental/app/models"
	"gopkg.in/gin-gonic/gin.v1"
)

// UnitIndex unit handler
func UnitIndex(c *gin.Context) {
	log.Println("UnitIndex")

	units, err := models.GetUnits()

	if err != nil {
		SetFlashError(c, err.Error())
		Redirect(c, "/home")
		return
	}

	RenderHTML(c, gin.H{
		"unitViewer": "View Units",
		"units":      units,
	})
}
