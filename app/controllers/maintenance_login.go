package controllers

import (
	"log"

	"gopkg.in/gin-gonic/gin.v1"
)

// MaintenanceLoginIndex maintenance index view
func MaintenanceLoginIndex(c *gin.Context) {
	log.Println("MaintenanceLoginIndex")

	RenderHTML(c, gin.H{
		"login": "Maintenance Login",
	})
}
