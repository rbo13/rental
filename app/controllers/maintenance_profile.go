package controllers

import (
	"log"

	"github.com/whaangbuu/home-rental/app/models"
	"gopkg.in/gin-gonic/gin.v1"
)

// MaintenanceProfileIndex maintenance index view
func MaintenanceProfileIndex(c *gin.Context) {
	log.Println("MaintenanceProfileIndex")

	maintenanceID := GetMyMaintenanceID(c)
	maintenance, err := models.GetMaintenanceByID(maintenanceID)

	firstName := "First Name"
	lastName := "Last Name"
	birthDate := "1 January, 1970"
	gender := ""
	phoneNumber := ""

	if err != nil {
		SetFlashError(c, "Error getting tenant")
		Redirect(c, "/home")
		return
	}

	if maintenance.FirstName != "" && maintenance.LastName != "" && maintenance.Birthdate != "" {
		firstName = maintenance.FirstName
		lastName = maintenance.LastName
		birthDate = maintenance.Birthdate
		gender = maintenance.Gender
		phoneNumber = maintenance.PhoneNumber
	}

	RenderHTML(c, gin.H{
		"profile":       "Maintenance Profile Page",
		"firstName":     firstName,
		"lastName":      lastName,
		"gender":        gender,
		"phoneNumber":   phoneNumber,
		"birthDate":     birthDate,
		"maintenanceID": maintenanceID,
	})
}
