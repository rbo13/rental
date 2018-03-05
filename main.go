package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/justinas/nosurf"
	"github.com/leekchan/gtf"
	"github.com/whaangbuu/home-rental/app/controllers"
	"github.com/whaangbuu/home-rental/app/libs/ezgintemplate"
	"github.com/whaangbuu/home-rental/app/models"
	"github.com/whaangbuu/home-rental/tmplfunc"
	"gopkg.in/gin-gonic/gin.v1"
)

func init() {
	log.SetFlags(log.Lshortfile)
	models.Setup()
}

func main() {
	router := gin.Default()

	store := sessions.NewCookieStore([]byte("secret1233"))
	router.Use(sessions.Sessions("mysession", store))
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	// Static Routes
	router.Static("/assets", "./assets")
	// router.StaticFile("/favicon.ico", "./assets/favicon.ico")
	router.Use(gin.Recovery())

	render := ezgintemplate.New()
	render.TemplatesDir = "app/views/"
	render.Layout = "layouts/base"
	render.Ext = ".tmpl"
	render.Debug = true
	// TODO:: implet template functions here
	funcMap := template.FuncMap{
		"convertStatusToString": tmplfunc.ConvertStatusToString,
	}

	// Inject our template func
	gtf.Inject(funcMap)

	render.TemplateFuncMap = funcMap
	router.HTMLRender = render.Init()

	initializeRoutes(router)
	router.Run(":8080")
}

func csrfFailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", nosurf.Reason(r))
}

func initializeRoutes(origRouter *gin.Engine) {
	router := origRouter.Group("")

	router.GET("/", controllers.WelcomeIndex)

	// Login
	router.GET("/login", controllers.LoginIndex)
	// Register
	router.GET("/register", controllers.RegisterIndex)
	// Dashboard
	router.GET("/home", controllers.DashboardIndex)

	// Logout
	router.GET("/logout", controllers.LogoutHandler)
	// Profile
	router.GET("/profile", controllers.ProfileIndex)

	// Rent
	router.GET("/rent", controllers.MyRentControllerIndex)

	// Search Result
	router.GET("/search/unit", controllers.SearchUnitIndex)
	router.GET("/unit/info/:unit_id", controllers.UnitInfoIndex)
	// Post Request
	router.POST("/register", controllers.RegisterPostHandler)
	router.POST("/login", controllers.LoginPostHandler)
	router.POST("/update/info", controllers.ProfileUpdateInfoHandler)
	router.POST("/update/phone", controllers.ProfileUpdatePhoneNumberHandler)
	router.POST("/update/email", controllers.ProfileUpdateEmailHandler)

	// Admin
	admin := origRouter.Group("/admin")
	{
		admin.GET("/", controllers.AdminIndex)
		admin.GET("/register", controllers.AdminRegisterIndex)
		admin.POST("/register", controllers.AdminRegisterHandler)
		admin.POST("/login", controllers.AdminLoginHandler)
		admin.GET("/home", controllers.AdminHomeIndex)
		admin.GET("/profile", controllers.AdminProfileIndex)
		admin.GET("/generate/user", controllers.AdminGenerateUser)
		admin.GET("/list/owner", controllers.AdminShowListOfOwnerByType)
		admin.GET("/unit", controllers.AdminUnitIndex)
		admin.GET("/manage", controllers.AdminManageIndex)
		admin.POST("/add/unit", controllers.AdminUnitAddHandler)
		admin.POST("/update/profile", controllers.AdminUpdateProfileHandler)
	}
}
