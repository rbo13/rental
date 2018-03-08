package controllers

import (
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/justinas/nosurf"
	"github.com/whaangbuu/home-rental/app/libs/tmplname"
	"github.com/whaangbuu/home-rental/app/models"
	"gopkg.in/gin-gonic/gin.v1"
)

// RenderTemplate ...
func RenderTemplate(c *gin.Context, tmpl string, data gin.H, statusCode int) {

	// setFlash
	data["host"] = c.Request.Host
	data["flash_error"] = GetFlashError(c)
	data["flash_warning"] = GetFlashWarning(c)
	data["flash_info"] = GetFlashInfo(c)
	data["flash_success"] = GetFlashSuccess(c)
	data["csrf_token"] = nosurf.Token(c.Request)
	data["is_login"] = IsLogin(c)
	data["current_uri"] = strings.ToLower(c.Request.URL.Path)
	// data["my_username"] = GetMyUsername(c)
	// data["profile_picture"] = GetMyProfilePic(c)
	data["my_account_id"] = GetMyAccountID(c)
	data["email"] = GetMyEmail(c)
	data["fullName"] = GetMyFullname(c)
	data["tenantID"] = GetMyTenantID(c)
	data["is_admin"] = IsAdmin(c)
	data["ownerID"] = GetMyOwnerID(c)

	if !IsAdmin(c) {
		data["tenantFirstname"] = GetMyFirstName(c)
		data["tenantLastname"] = GetMyLastName(c)
	}
	// data["is_verified"] = GetISVerified(c)
	// data["timezone_offset"] = GetMyTimezoneOffset(c)
	// data["paid_plan_id"] = GetMyPaidPlanID(c)
	// data["access_token"], _ = GetAccessToken(c)
	// data["user_subscription"] = helpers.CheckSubscription(GetMyPaidPlanID(c))
	// data["has_expired_subscription"] = helpers.HasExpiredSubscription(GetMyUsername(c))
	// data["user_is_registered"] = helpers.IsUserRegistered(c.Param("username"))
	// userID := strconv.Itoa(int(GetMyUserID(c)))
	// data["my_string_user_id"] = userID
	// localeStr := c.MustGet("locale")
	// locale, ok := localeStr.(string)
	// if ok && locale != "" {
	// 	data["locale"] = locale
	// } else {
	// 	data["locale"] = "en"
	// }
	c.HTML(statusCode, tmpl, data)
}

// RenderHTML ...
func RenderHTML(c *gin.Context, data gin.H) {

	pc, _, _, _ := runtime.Caller(1)
	callerName := runtime.FuncForPC(pc).Name()
	for strings.Contains(callerName, ".") {
		a := strings.Index(callerName, ".")
		callerName = callerName[a+1:]
	}

	tmpl := tmplname.Convert(callerName)
	// check whether the file is existed.
	_, err := os.Stat("app/views/" + tmpl + ".tmpl")
	if err != nil {
		c.String(200, "%s not found", "app/views/"+tmpl+".tmpl")
		return
	}
	amptmpl := tmpl + "_amp"
	if c.Request.URL.Query().Get("amp") == "1" || strings.HasPrefix(c.Request.URL.Path, "/amp/") {
		_, err := os.Stat("app/views/" + amptmpl + ".tmpl")
		if err != nil {
			log.Printf("%s not found, so use %s instead.", "app/views/"+amptmpl+".tmpl", "app/views/"+tmpl+".tmpl")
		} else {
			tmpl = amptmpl
		}
	}
	RenderTemplate(c, tmpl, data, 200)
}

const flashKeyInfo = "flash_key_info"
const flashKeyError = "flash_key_Error"
const flashKeyWarning = "flash_key_warning"
const flashKeySuccess = "flash_key_Success"

// SetAuth ...
func SetAuth(c *gin.Context, account models.Account) {
	session := sessions.Default(c)
	session.Set("account_id", account.ID)
	session.Set("email", account.EmailAddress)
	session.Set("is_admin", account.IsAdmin)
	session.Set("is_login", 1)
	session.Save()
}

// ClearAuth ...
func ClearAuth(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("is_login")
	session.Delete("account_id")
	session.Delete("is_admin")
	session.Delete("email")

	session.Save()
}

// SetFlashInfo ...
func SetFlashInfo(c *gin.Context, msg string) error {
	return setFlash(c, msg, flashKeyInfo)
}

// GetFlashInfo ...
func GetFlashInfo(c *gin.Context) string {
	return getFlash(c, flashKeyInfo)
}

// SetFlashWarning ...
func SetFlashWarning(c *gin.Context, msg string) error {
	return setFlash(c, msg, flashKeyWarning)
}

// GetFlashWarning ...
func GetFlashWarning(c *gin.Context) string {
	return getFlash(c, flashKeyWarning)
}

// SetFlashError ...
func SetFlashError(c *gin.Context, msg string) error {
	return setFlash(c, msg, flashKeyError)
}

// GetFlashError ...
func GetFlashError(c *gin.Context) string {
	return getFlash(c, flashKeyError)
}

// SetFlashSuccess ...
func SetFlashSuccess(c *gin.Context, msg string) error {
	return setFlash(c, msg, flashKeySuccess)
}

// GetFlashSuccess ...
func GetFlashSuccess(c *gin.Context) string {
	return getFlash(c, flashKeySuccess)
}

///

func setFlash(c *gin.Context, msg, key string) error {
	session := sessions.Default(c)
	if msg == "" {
		return nil
	}
	session.Set(key, msg)
	session.Save()
	return nil
}

func getFlash(c *gin.Context, key string) string {
	session := sessions.Default(c)
	obj := session.Get(key)
	if msg, ok := obj.(string); ok {
		session.Delete(key)
		session.Save()
		return msg
	}
	return ""
}

// Redirect redirects to a url
func Redirect(c *gin.Context, url string) {

	c.Redirect(302, url)
	c.Abort()
}

// IsLogin boolean function that checks
// if a user is logged in.
func IsLogin(c *gin.Context) bool {
	isLogin, ok := c.Get("is_login")

	if ok && isLogin.(bool) {
		return true
	}

	session := sessions.Default(c)
	if flag := session.Get("is_login"); flag != nil {
		val, ok := flag.(int)
		if ok && val == 1 {
			return true
		}
	}

	return false
}

// GetMyEmail ...
func GetMyEmail(c *gin.Context) string {
	session := sessions.Default(c)

	if email := session.Get("email"); email != nil {
		valEmail, ok := email.(string)
		if ok {
			return valEmail
		}
	}

	return ""
}

// GetMyAccountID ...
func GetMyAccountID(c *gin.Context) int64 {
	session := sessions.Default(c)

	if userID := session.Get("account_id"); userID != nil {
		valUserID, ok := userID.(int64)
		if ok {
			return valUserID
		}
	}

	return 0
}

// IsAdmin a boolean function that checks
// if an account type is admin or not.
func IsAdmin(c *gin.Context) bool {
	session := sessions.Default(c)

	if isAdmin := session.Get("is_admin"); isAdmin != nil {
		isAdmin, ok := isAdmin.(bool)
		if ok {
			return isAdmin
		}
	}

	return false
}

// GetMyFullname returns the fullname
func GetMyFullname(c *gin.Context) string {
	if IsLogin(c) {
		accountID := GetMyAccountID(c)
		tenant, err := models.GetTenantByAccountID(accountID)
		if err != nil {
			log.Println(err)
			return ""
		}
		return tenant.FirstName + " " + tenant.LastName
	}
	return ""
}

// GetMyFirstName returns the firstname
func GetMyFirstName(c *gin.Context) string {
	if IsLogin(c) {
		accountID := GetMyAccountID(c)
		tenant, err := models.GetTenantByAccountID(accountID)
		if err != nil {
			log.Println(err)
			return ""
		}
		return tenant.FirstName
	}
	return ""
}

// GetMyLastName returns the lastname
func GetMyLastName(c *gin.Context) string {
	if IsLogin(c) {
		accountID := GetMyAccountID(c)
		tenant, err := models.GetTenantByAccountID(accountID)
		if err != nil {
			log.Println(err)
			return ""
		}
		return tenant.LastName
	}
	return ""
}

// GetMyTenantID gets the tenant ID
func GetMyTenantID(c *gin.Context) int64 {
	if IsLogin(c) {
		accountID := GetMyAccountID(c)
		tenant, err := models.GetTenantByAccountID(accountID)
		if err != nil {
			log.Println(err)
			return 0
		}
		return tenant.ID
	}
	return 0
}

// GetMyOwnerID returns an owner ID
func GetMyOwnerID(c *gin.Context) int64 {
	emailAddress := GetMyEmail(c)
	if IsLogin(c) {
		owner, err := models.GetOwnerByEmailAddress(emailAddress)
		if err != nil {
			log.Println(err)
			return 0
		}
		if owner != nil {
			return owner.ID
		}
		return 0
	}

	return 0
}

// RedirectIfNotLogin ...
func RedirectIfNotLogin(c *gin.Context, urlstr, msg string) {
	if IsLogin(c) == false {
		if msg != "" {
			//log.Print("")
			SetFlashError(c, msg)
		}
		Redirect(c, urlstr)
		c.Abort()
	}
}
