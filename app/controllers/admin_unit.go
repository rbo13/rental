package controllers

import (
	"log"
	"strconv"

	"github.com/whaangbuu/home-rental/app/models"
	"gopkg.in/gin-gonic/gin.v1"
)

// AdminUnitIndex ...
func AdminUnitIndex(c *gin.Context) {
	log.Println("AdminUnitIndex")

	RenderHTML(c, gin.H{
		"addUnit": "Add a Unit",
	})
}

// AdminUnitAddHandler ...
func AdminUnitAddHandler(c *gin.Context) {
	log.Println("AdminUnitAddHandler")

	unitType := c.PostForm("unitType")
	unitNumberStr := c.PostForm("unitNumber")
	unitPriceStr := c.PostForm("unitPrice")
	ownerID := GetMyAccountID(c)

	unitPrice, parseErr := strconv.ParseInt(unitPriceStr, 10, 64)
	unitNumber, err := strconv.ParseInt(unitNumberStr, 10, 64)
	if parseErr != nil && err != nil {
		SetFlashError(c, "ERROR PARSING")
		Redirect(c, "/admin")
		return
	}

	unit := models.NewUnit(ownerID, unitPrice, unitNumber, unitType, true)
	unit.Create()

	SetFlashSuccess(c, "Unit successfully added")
	Redirect(c, "/admin/home")
}
