package controllers

import (
	"log"
	"strconv"

	"github.com/whaangbuu/home-rental/app/models"
	"gopkg.in/gin-gonic/gin.v1"
)

// AcceptRequestIndex admin home index
func AcceptRequestIndex(c *gin.Context) {
	log.Println("AcceptRequestIndex")

	requestIDStr := c.Param("request_id")
	requestID, parseErr := strconv.ParseInt(requestIDStr, 10, 64)

	if parseErr != nil {
		log.Println(parseErr)
		Redirect(c, "/admin/home")
		return
	}

	request, err := models.GetTenantRecordByID(requestID)

	if err != nil {
		SetFlashError(c, err.Error())
		Redirect(c, "/admin/home")
		return
	}
	request.TenantStatus = true
	request.Update()

	SetFlashSuccess(c, "Request Accepted")
	Redirect(c, "/admin/home")
}
