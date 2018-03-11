package controllers

import (
	"log"

	"gopkg.in/gin-gonic/gin.v1"
)

// MaintenancePayslipIndex admin index view
func MaintenancePayslipIndex(c *gin.Context) {
	log.Println("MaintenancePayslipIndex")

	if !IsLogin(c) {
		Redirect(c, "/maintenance/login")
		return
	}

	RenderHTML(c, gin.H{
		"payslip": "My Payslip",
	})
}
