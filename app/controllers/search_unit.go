package controllers

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/whaangbuu/home-rental/app/models"
	"gopkg.in/gin-gonic/gin.v1"
)

// SearchUnitIndex ...
func SearchUnitIndex(c *gin.Context) {
	log.Println("SearchUnitIndex")

	unitType := c.Request.URL.Query().Get("type")
	startDate := c.Request.URL.Query().Get("startDate")
	endDate := c.Request.URL.Query().Get("endDate")
	session := sessions.Default(c)
	session.Set("startDate", startDate)
	session.Set("endDate", endDate)
	session.Set("unitType", unitType)
	session.Save()

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
		"units":    units,
		"unitType": unitType,
	})
}
