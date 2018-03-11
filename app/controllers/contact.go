package controllers

import (
	"log"

	"gopkg.in/gin-gonic/gin.v1"
)

// ContactIndex admin index view
func ContactIndex(c *gin.Context) {
	log.Println("ContactIndex")

	RenderHTML(c, gin.H{
		"contact": "Contact Us",
	})
}
