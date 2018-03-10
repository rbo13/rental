package controllers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/whaangbuu/home-rental/app/models"
	"gopkg.in/gin-gonic/gin.v1"
)

// UnitInfoIndex unit info index
func UnitInfoIndex(c *gin.Context) {
	log.Println("UnitInfoIndex")
	var startDate interface{}
	var endDate interface{}
	var startDateView string
	var endDateView string
	session := sessions.Default(c)

	if !IsLogin(c) {
		SetFlashError(c, "Authentication is required")
		Redirect(c, "/login")
		return
	}

	if startDate = session.Get("startDate"); startDate != nil {
		val, ok := startDate.(int)
		startDateView = startDate.(string)
		if ok && val == 1 {
			log.Println(val)
		}
	}
	if endDate = session.Get("endDate"); endDate != nil {
		val, ok := endDate.(int)
		endDateView = endDate.(string)
		if ok && val == 1 {
			log.Println(val)
		}
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

	session.Delete("startDate")
	session.Delete("endDate")
	session.Save()

	RenderHTML(c, gin.H{
		"unitInfo":      "Unit Information",
		"ownerID":       unit.OwnerID,
		"unit":          unit,
		"owner":         owner,
		"ownerFullName": ownerFullName,
		"startDate":     startDateView,
		"endDate":       endDateView,
	})
}
