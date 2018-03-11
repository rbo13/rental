package controllers

import (
	"log"

	"gopkg.in/gin-gonic/gin.v1"
)

// MaintenanceRegisterIndex maintenance index view
func MaintenanceRegisterIndex(c *gin.Context) {
	log.Println("MaintenanceRegisterIndex")

	RenderHTML(c, gin.H{
		"register": "Maintenance Register",
	})
}
