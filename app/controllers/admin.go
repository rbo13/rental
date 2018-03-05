package controllers

import (
	"log"
	"strconv"

	"github.com/whaangbuu/home-rental/app/models"
	"gopkg.in/gin-gonic/gin.v1"
)

// AdminIndex admin index view
func AdminIndex(c *gin.Context) {
	log.Println("AdminIndex")

	if IsLogin(c) {
		Redirect(c, "/admin/home")
		return
	}

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

	// create to owner
	owner := models.NewOwner("", "", emailAddress, "", "")
	owner.Create()

	err := account.Create()

	if err != nil {
		SetFlashError(c, err.Error())
		Redirect(c, "/admin/register")
		return
	}

	SetFlashSuccess(c, "Successfully registered")
	Redirect(c, "/admin")
}

// AdminProfileIndex admin profile setting
func AdminProfileIndex(c *gin.Context) {
	log.Println("AdminProfileIndex")

	accountID := GetMyAccountID(c)
	account, err := models.GetAccountByID(accountID)

	if err != nil {
		SetFlashError(c, err.Error())
		Redirect(c, "/")
		return
	}

	RenderHTML(c, gin.H{
		"adminProfile": "Update Profile as Admin",
		"accountID":    accountID,
		"account":      account,
	})
}

// AdminUpdateProfileHandler handles updating
// of admin profile
func AdminUpdateProfileHandler(c *gin.Context) {
	log.Printf("AdminUpdateProfileHandler")

	emailAddress := c.PostForm("email")
	lastName := c.PostForm("lastName")
	firstName := c.PostForm("firstName")
	phoneNumber := c.PostForm("phoneNumber")
	unitType := c.PostForm("unitType")
	numberOfPropertyStr := c.PostForm("numberOfProperty")

	numberOfProperty, parseErr := strconv.ParseInt(numberOfPropertyStr, 10, 64)
	if parseErr != nil {
		SetFlashError(c, parseErr.Error())
		Redirect(c, "/admin/profile")
		return
	}

	owner, err := models.GetOwnerByEmailAddress(emailAddress)

	if err != nil {
		SetFlashError(c, err.Error())
		Redirect(c, "/admin/profile")
		return
	}

	owner.EmailAddress = emailAddress
	owner.LastName = lastName
	owner.FirstName = firstName
	owner.PhoneNumber = phoneNumber
	owner.TypeOfPropertyOwned = unitType
	owner.NoOfPropertyOwned = numberOfProperty
	owner.Update()

	SetFlashSuccess(c, "Update successfully")
	Redirect(c, "/admin/profile")
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
