package controllers

import (
	"log"

	"gopkg.in/gin-gonic/gin.v1"
)

// MaintenanceIndex maintenance index view
func MaintenanceIndex(c *gin.Context) {
	log.Println("MaintenanceIndex")

	// if IsLogin(c) {
	// 	Redirect(c, "/maintenance/home")
	// 	return
	// }

	RenderHTML(c, gin.H{
		"maintenance": "Maintenance Page",
	})
}
