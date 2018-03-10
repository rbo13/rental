package controllers

import (
	"log"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/whaangbuu/home-rental/app/models"
	"gopkg.in/gin-gonic/gin.v1"
)

// RequestRentHandler unit info index
func RequestRentHandler(c *gin.Context) {
	log.Println("RequestRentHandler")
	session := sessions.Default(c)
	var startDate interface{}
	var endDate interface{}
	var unitType interface{}

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
	if startDate = session.Get("startDate"); startDate != nil {
		val, ok := startDate.(int)

		if ok && val == 1 {
			log.Println(val)
		}
	}
	if endDate = session.Get("endDate"); endDate != nil {
		val, ok := endDate.(int)

		if ok && val == 1 {
			log.Println(val)
		}
	}
	if unitType = session.Get("unitType"); unitType != nil {
		val, ok := unitType.(int)

		if ok && val == 1 {
			log.Println(val)
		}
	}

	tenantFirstName := GetMyFirstName(c)
	tenantLastName := GetMyLastName(c)
	tenantID := GetMyAccountID(c)

	unit, err := models.GetUnitByUnitID(unitID)
	owner, getErr := models.GetOwnerByID(unit.OwnerID)

	if getErr != nil {
		SetFlashError(c, getErr.Error())
		Redirect(c, "/home")
		return
	}

	if err != nil {
		SetFlashError(c, err.Error())
		Redirect(c, "/home")
		return
	}

	if tenantFirstName == "" && tenantLastName == "" {
		SetFlashError(c, "You need to update your profile")
		Redirect(c, "/profile")
		return
	}

	tenantRecord := models.NewTenantRecord()
	tenantRecord.FirstName = tenantFirstName
	tenantRecord.LastName = tenantLastName
	tenantRecord.TenantID = tenantID
	tenantRecord.UnitID = unitID
	tenantRecord.OwnerID = owner.ID
	tenantRecord.StartDate = startDate.(string)
	tenantRecord.EndDate = endDate.(string)
	tenantRecord.UnitType = unitType.(string)
	tenantRecord.TenantStatus = false
	tenantRecord.PaymentStatus = false

	createErr := tenantRecord.Create()

	if createErr != nil {
		SetFlashError(c, "Error inserting new tenant")
		Redirect(c, "/home")
		return
	}

	session.Delete("startDate")
	session.Delete("endDate")
	session.Delete("unitType")

	SetFlashSuccess(c, "Request sent")
	Redirect(c, "/home")
}
