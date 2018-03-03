package controllers

import (
	"log"

	"github.com/whaangbuu/home-rental/app/models"
	"gopkg.in/gin-gonic/gin.v1"
)

// SearchUnitIndex ...
func SearchUnitIndex(c *gin.Context) {
	log.Println("SearchUnitIndex")

	unitType := c.Request.URL.Query().Get("unitType")
	log.Println(unitType)

	if (unitType != "Room") && (unitType != "House") {
		SetFlashError(c, "Unit type not found")
		Redirect(c, "/home")
		return
	}

	units, err := models.GetUnitsByUnitType(unitType)

	if err != nil {
		SetFlashError(c, "Unit not found")
		Redirect(c, "/home")
		return
	}
	RenderHTML(c, gin.H{
		"units": units,
	})
}
