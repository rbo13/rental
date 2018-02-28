package controllers

import (
	"log"

	"gopkg.in/gin-gonic/gin.v1"
)

// AdminIndex admin index view
func AdminIndex(c *gin.Context) {
	log.Println("AdminIndex")

	RenderHTML(c, gin.H{
		"admin": "Admin Page",
	})
}
