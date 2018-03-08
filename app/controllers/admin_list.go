package controllers

import (
	"log"

	"github.com/whaangbuu/home-rental/app/models"
	"gopkg.in/gin-gonic/gin.v1"
)

// ListTenantsIndex list of all tenants
func ListTenantsIndex(c *gin.Context) {
	log.Println("ListTenantsIndex")

	tenants, err := models.GetTenantsUsingStoredProcedure()

	if err != nil {
		SetFlashError(c, "CANT GET FROM STORED PROCEDURE")
		Redirect(c, "/admin/manage")
		return
	}

	RenderHTML(c, gin.H{
		"tenantsList": "List of Tenants",
		"tenants":     tenants,
	})
}
