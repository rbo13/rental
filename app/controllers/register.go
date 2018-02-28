package controllers

import (
	"log"

	"github.com/whaangbuu/home-rental/app/models"
	"gopkg.in/gin-gonic/gin.v1"
)

// RegisterIndex handles the index page
func RegisterIndex(c *gin.Context) {
	log.Println("LoginIndex")

	RenderHTML(c, gin.H{
		"register": "Register Page",
	})
}

// RegisterPostHandler handles registration of user
func RegisterPostHandler(c *gin.Context) {
	log.Println("RegisterPostHandler")

	emailAddress := c.PostForm("emailAddress")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirmPassword")

	if password != confirmPassword {
		msg := "Password did not match"
		SetFlashError(c, msg)
		Redirect(c, "/register")
	}

	account := models.NewAccount()
	account.EmailAddress = emailAddress
	account.Password = password

	err := account.Create()

	if err != nil {
		SetFlashError(c, err.Error())
		Redirect(c, "/register")
		return
	}

	// Create record of tenant at the same time
	tenant := models.NewTenant(account.ID, "", "", "", "", "")
	tenant.Create()

	successMessage := "Successfully registered, please login to continue"
	SetFlashSuccess(c, successMessage)
	Redirect(c, "/login")
}
