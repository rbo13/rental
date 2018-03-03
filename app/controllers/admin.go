package controllers

import (
	"log"

	"github.com/whaangbuu/home-rental/app/models"
	"gopkg.in/gin-gonic/gin.v1"
)

// AdminIndex admin index view
func AdminIndex(c *gin.Context) {
	log.Println("AdminIndex")

	RenderHTML(c, gin.H{
		"admin": "Admin Page",
	})
}

// AdminRegisterIndex register admin
func AdminRegisterIndex(c *gin.Context) {
	log.Println("AdminRegisterIndex")

	RenderHTML(c, gin.H{
		"adminRegister": "Register Admin",
	})
}

// AdminRegisterHandler handles registering of admin/owner
func AdminRegisterHandler(c *gin.Context) {
	log.Println("AdminRegisterIndex")

	emailAddress := c.PostForm("emailAddress")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirmPassword")

	if emailAddress == "" {
		SetFlashError(c, "Email is required")
		Redirect(c, "/admin/register")
		return
	}

	if password != confirmPassword {
		SetFlashError(c, "Password dont match")
		Redirect(c, "/admin/register")
		return
	}

	account := models.NewAccount()
	account.EmailAddress = emailAddress
	account.Password = password
	account.IsAdmin = true

	err := account.Create()

	if err != nil {
		SetFlashError(c, err.Error())
		Redirect(c, "/admin/register")
		return
	}

	SetFlashSuccess(c, "Successfully registered")
	Redirect(c, "/admin")
}

// AdminLoginHandler admin login handler
func AdminLoginHandler(c *gin.Context) {
	log.Println("AdminLoginHandler")

	if IsLogin(c) {
		Redirect(c, "/home")
	}

	emailAddress := c.PostForm("email")
	password := c.PostForm("password")
	account, err := models.LoginUser(emailAddress, password)

	if err != nil {
		SetFlashError(c, err.Error())
		Redirect(c, "/admin/register")
		return
	}

	SetAuth(c, *account)
	SetFlashSuccess(c, "Successfully loggedd in")
	Redirect(c, "/admin/home")
}
