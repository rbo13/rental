package controllers

import (
	"log"

	"gopkg.in/gin-gonic/gin.v1"
)

// AdminHomeIndex admin home index
func AdminHomeIndex(c *gin.Context) {
	log.Println("AdminHomeIndex")

	RenderHTML(c, gin.H{
		"adminHome": "Admin Home",
	})
}
