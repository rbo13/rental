package controllers

import (
	"log"

	"gopkg.in/gin-gonic/gin.v1"
)

// MaintenanceRequestIndex maintenance index view
func MaintenanceRequestIndex(c *gin.Context) {
	log.Println("MaintenanceRequestIndex")

	RenderHTML(c, gin.H{
		"tenantRequest": "Tenant Requests",
	})
}

// MaintenanceRequestInfoIndex request info
func MaintenanceRequestInfoIndex(c *gin.Context) {
	log.Println("MaintenanceRequestInfoIndex")

	RenderHTML(c, gin.H{
		"tenantRequestInfo": "Tenant Requests Information",
	})
}
