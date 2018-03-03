package controllers

import (
	"log"

	"gopkg.in/gin-gonic/gin.v1"
)

// AdminUnitIndex ...
func AdminUnitIndex(c *gin.Context) {
	log.Println("AdminUnitIndex")

	RenderHTML(c, gin.H{
		"addUnit": "Add a Unit",
	})
}

// AdminUnitAddHandler ...
func AdminUnitAddHandler(c *gin.Context) {
	log.Println("AdminUnitAddHandler")

	unitType := c.PostForm("unitType")
	unitNumber := c.PostForm("unitNumber")
	unitPrice := c.PostForm("unitPrice")

	log.Println(unitType)
	log.Println(unitNumber)
	log.Println(unitPrice)

	SetFlashSuccess(c, "Unit successfully added")
	Redirect(c, "/admin/home")
}
