package controllers

import (
	"log"

	"github.com/whaangbuu/home-rental/app/models"
	"gopkg.in/gin-gonic/gin.v1"
)

// LoginIndex ...
func LoginIndex(c *gin.Context) {
	log.Println("LoginIndex")

	RenderHTML(c, gin.H{
		"login": "Login Page",
	})
}

// LoginPostHandler handles logging in
func LoginPostHandler(c *gin.Context) {
	log.Println("LoginPostHandler")

	if IsLogin(c) {
		Redirect(c, "/home")
	}

	email := c.PostForm("email")
	password := c.PostForm("password")

	account, err := models.LoginUser(email, password)

	if err != nil {
		SetFlashError(c, err.Error())
		Redirect(c, "/register")
		return
	}

	SetAuth(c, *account)
	SetFlashSuccess(c, "Login successful")
	Redirect(c, "/home")
}
