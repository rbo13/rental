package controllers

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/whaangbuu/home-rental/app/models"
	"gopkg.in/gin-gonic/gin.v1"
)

// ProfileIndex profile index
func ProfileIndex(c *gin.Context) {
	log.Println("ProfileIndexHandler")

	accountID := GetMyAccountID(c)
	tenant, err := models.GetTenantByAccountID(accountID)

	log.Println("ACCOUNT ID")
	log.Println(accountID)

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

	if tenant.FirstName != "" && tenant.LastName != "" && tenant.Birthdate != "" {
		firstName = tenant.FirstName
		lastName = tenant.LastName
		birthDate = tenant.Birthdate
		gender = tenant.Gender
		phoneNumber = tenant.PhoneNumber
	}

	RenderHTML(c, gin.H{
		"profile":     "Profile Page",
		"firstName":   firstName,
		"lastName":    lastName,
		"gender":      gender,
		"phoneNumber": phoneNumber,
		"birthDate":   birthDate,
	})
}

// ProfileUpdateInfoHandler handles updating
// of profile information
func ProfileUpdateInfoHandler(c *gin.Context) {
	log.Println("ProfileUpdateInfoHandler")

	firstName := c.PostForm("firstName")
	lastName := c.PostForm("lastName")
	gender := c.PostForm("gender")
	birthDate := c.PostForm("birthDate")

	accountID := GetMyAccountID(c)
	account, err := models.GetAccountByID(accountID)
	if err != nil {
		SetFlashError(c, err.Error())
		Redirect(c, "/profile")
		return
	}

	tenant := models.NewTenant(account.ID, lastName, firstName, "", gender, birthDate)
	updateTenantErr := tenant.Update()
	if updateTenantErr != nil {
		SetFlashError(c, "Error creating new tenant")
		Redirect(c, "/profile")
		return
	}

	successMessage := "Profile updated successfully"
	SetFlashSuccess(c, successMessage)
	Redirect(c, "/profile")
}

// ProfileUpdatePhoneNumberHandler updates phone number
func ProfileUpdatePhoneNumberHandler(c *gin.Context) {
	log.Println("ProfileUpdatePhoneNumberHandler")

	phoneNumber := c.PostForm("phoneNumber")

	accountID := GetMyAccountID(c)
	account, err := models.GetAccountByID(accountID)

	if err != nil {
		SetFlashError(c, err.Error())
		Redirect(c, "/profile")
		return
	}

	updateTenant, getErr := models.GetTenantByAccountID(account.ID)

	if getErr != nil {
		SetFlashError(c, getErr.Error())
		Redirect(c, "/profile")
		return
	}

	updateTenant.PhoneNumber = phoneNumber
	updateTenant.Update()

	successMessage := "Phone number updated successfully"
	SetFlashSuccess(c, successMessage)
	Redirect(c, "/profile")
}

// ProfileUpdateEmailHandler ...
func ProfileUpdateEmailHandler(c *gin.Context) {
	log.Println("ProfileUpdateEmailHandler")
	session := sessions.Default(c)

	newEmail := c.PostForm("email")
	accountID := GetMyAccountID(c)

	account, err := models.GetAccountByID(accountID)
	if err != nil {
		SetFlashError(c, err.Error())
		Redirect(c, "/profile")
		return
	}

	account.EmailAddress = newEmail
	account.Update()
	session.Set("email", newEmail)
	session.Save()

	successMessage := "Email updated succesfully"
	SetFlashSuccess(c, successMessage)
	Redirect(c, "/profile")
}
