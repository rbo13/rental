package controllers

import (
	"log"

	"github.com/whaangbuu/home-rental/app/models"
	"gopkg.in/gin-gonic/gin.v1"
)

// MaintenanceLoginIndex maintenance index view
func MaintenanceLoginIndex(c *gin.Context) {
	log.Println("MaintenanceLoginIndex")

	RenderHTML(c, gin.H{
		"login": "Maintenance Login",
	})
}

// MaintenanceLoginPostHandler handles logging in
func MaintenanceLoginPostHandler(c *gin.Context) {
	log.Println("MaintenanceLoginPostHandler")

	if IsLogin(c) {
		Redirect(c, "/home")
	}

	email := c.PostForm("email")
	password := c.PostForm("password")

	maintenance, err := models.LoginMaintenance(email, password)

	if err != nil {
		SetFlashError(c, err.Error())
		Redirect(c, "/maintenance/register")
		return
	}

	SetAuthMaintenance(c, *maintenance)
	SetFlashSuccess(c, "Login successful")
	Redirect(c, "/maintenance")
}
