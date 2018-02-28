package controllers

import (
	"log"

	"gopkg.in/gin-gonic/gin.v1"
)

// WelcomeIndex handles the welcome page
func WelcomeIndex(c *gin.Context) {
	log.Println("WelcomeIndex")

	RenderHTML(c, gin.H{
		"greetings": "Welcome Page",
	})
}
