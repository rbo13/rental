package controllers

import (
	"log"

	"gopkg.in/gin-gonic/gin.v1"
)

// MaintenanceContactIndex admin index view
func MaintenanceContactIndex(c *gin.Context) {
	log.Println("MaintenanceContactIndex")

	RenderHTML(c, gin.H{
		"contact": "Contact Us",
	})
}
