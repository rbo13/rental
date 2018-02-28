package controllers

import (
	"log"

	"gopkg.in/gin-gonic/gin.v1"
)

// DashboardIndex handles index page
func DashboardIndex(c *gin.Context) {
	log.Println("DashboardIndex")

	if !IsLogin(c) {
		errorMessage := "Login is required"
		SetFlashError(c, errorMessage)
		Redirect(c, "/login")
		return
	}

	RenderHTML(c, gin.H{
		"dashboard": "Dashboard Page",
	})
}
