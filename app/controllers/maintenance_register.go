package controllers

import (
	"log"

	"github.com/whaangbuu/home-rental/app/models"
	"gopkg.in/gin-gonic/gin.v1"
)

// MaintenanceRegisterIndex maintenance index view
func MaintenanceRegisterIndex(c *gin.Context) {
	log.Println("MaintenanceRegisterIndex")

	RenderHTML(c, gin.H{
		"register": "Maintenance Register",
	})
}

// MaintenanceRegisterHandler maintenance register handler
func MaintenanceRegisterHandler(c *gin.Context) {
	log.Println("MaintenanceRegisterHandler")

	emailAddress := c.PostForm("emailAddress")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirmPassword")

	if password != confirmPassword {
		msg := "Password did not match"
		SetFlashError(c, msg)
		Redirect(c, "/register")
	}

	maintenance := models.NewMaintenance()
	maintenance.EmailAddress = emailAddress
	maintenance.Password = password

	err := maintenance.Create()

	if err != nil {
		SetFlashError(c, err.Error())
		Redirect(c, "/maintenance/register")
		return
	}

	successMessage := "Successfully registered, please login to continue"
	SetFlashSuccess(c, successMessage)
	Redirect(c, "/maintenance/login")
}
