package controllers

import (
	"log"

	"github.com/whaangbuu/home-rental/app/models"
	"gopkg.in/gin-gonic/gin.v1"
)

// MyRentControllerIndex my rent controller
// index handler
func MyRentControllerIndex(c *gin.Context) {
	log.Println("MyRentControllerIndex")
	tenantID := GetMyTenantID(c)
	rent, err := models.GetUnitByTenantID(tenantID)
	if err != nil {
		SetFlashError(c, "Error getting rent")
		Redirect(c, "/profile")
		return
	}
	log.Print(rent)
	RenderHTML(c, gin.H{
		"rentPage": "My Rent Page",
		"rent":     rent,
	})
}
