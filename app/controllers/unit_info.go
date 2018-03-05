package controllers

import (
	"log"

	"gopkg.in/gin-gonic/gin.v1"
)

// UnitInfoIndex unit info index
func UnitInfoIndex(c *gin.Context) {
	log.Println("UnitInfoIndex")

	RenderHTML(c, gin.H{
		"unitInfo": "Unit Information",
	})
}
