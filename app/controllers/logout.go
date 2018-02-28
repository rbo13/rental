package controllers

import "gopkg.in/gin-gonic/gin.v1"

// LogoutHandler handles logging out of user
func LogoutHandler(c *gin.Context) {
	ClearAuth(c)

	SetFlashSuccess(c, "You have logged out")
	Redirect(c, "/?logout=1")
}
