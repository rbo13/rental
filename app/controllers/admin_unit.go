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

	emailAddress := GetMyEmail(c)
	owner, err := models.GetOwnerByEmailAddress(emailAddress)

	if err != nil {
		SetFlashError(c, "Admin not found")
		Redirect(c, "/admin/home")
		return
	}

	if owner.FirstName == "" && owner.LastName == "" {
		SetFlashError(c, "You need to update your profile")
		Redirect(c, "/admin/profile")
		return
	}

	log.Println("#######")
	log.Println(owner.TypeOfPropertyOwned)

	RenderHTML(c, gin.H{
		"addUnit":   "Add a Unit",
		"unitOwned": owner.TypeOfPropertyOwned,
	})
}

// AdminUnitAddHandler ...
func AdminUnitAddHandler(c *gin.Context) {
	log.Println("AdminUnitAddHandler")

	unitType := c.PostForm("unitType")
	unitNumberStr := c.PostForm("unitNumber")
	unitPriceStr := c.PostForm("unitPrice")
	ownerID := GetMyOwnerID(c)

	unitPrice, parseErr := strconv.ParseInt(unitPriceStr, 10, 64)
	unitNumber, err := strconv.ParseInt(unitNumberStr, 10, 64)
	if parseErr != nil && err != nil {
		SetFlashError(c, "ERROR PARSING")
		Redirect(c, "/admin")
		return
	}

	if unitType == "" {
		SetFlashError(c, "Unit is required")
		Redirect(c, "/admin/unit")
		return
	}

	unit := models.NewUnit(ownerID, unitPrice, unitNumber, unitType, true)
	unit.Create()

	SetFlashSuccess(c, "Unit successfully added")
	Redirect(c, "/admin/home")
}
