package controllers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/whaangbuu/home-rental/app/models"
	"gopkg.in/gin-gonic/gin.v1"
)

// UnitInfoIndex unit info index
func UnitInfoIndex(c *gin.Context) {
	log.Println("UnitInfoIndex")

	if !IsLogin(c) {
		SetFlashError(c, "Authentication is required")
		Redirect(c, "/login")
		return
	}

	unitIDStr := c.Param("unit_id")
	unitID, parseErr := strconv.ParseInt(unitIDStr, 10, 64)
	if parseErr != nil {
		SetFlashError(c, parseErr.Error())
		Redirect(c, "/home")
		return
	}

	unit, err := models.GetUnitByUnitID(unitID)

	if err != nil {
		msg := fmt.Sprintf("ERROR GETTING UNIT: %v", err)
		SetFlashError(c, msg)
		Redirect(c, "/home")
		return
	}
	owner, err := models.GetOwnerByID(unit.OwnerID)
	ownerFullName := owner.FirstName + " " + owner.LastName

	RenderHTML(c, gin.H{
		"unitInfo":      "Unit Information",
		"ownerID":       unit.OwnerID,
		"unit":          unit,
		"owner":         owner,
		"ownerFullName": ownerFullName,
	})
}
