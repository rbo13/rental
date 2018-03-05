package controllers

import (
	"log"

	"github.com/whaangbuu/home-rental/app/models"
	"gopkg.in/gin-gonic/gin.v1"
)

// AdminGenerateUser generates fake data
func AdminGenerateUser(c *gin.Context) {
	log.Println("AdminGenerate")
	for index := 0; index < 5; index++ {
		// models.InsertFakeData()
	}

	SetFlashInfo(c, "User generated")
	Redirect(c, "/admin")
}

// AdminShowListOfOwnerByType shows the list
// of owner by property type. TODO :: enhance
func AdminShowListOfOwnerByType(c *gin.Context) {
	log.Println("AdminShowListOfOwnerByType")

	owners, err := models.GetOwnersByPropertyType("Room")
	if err != nil {
		SetFlashError(c, "ERROR GETTING LIST OF OWNERS:")
		log.Print(err)
		Redirect(c, "/admin")
		return
	}
	log.Println(owners)
}
