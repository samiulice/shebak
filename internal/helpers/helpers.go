package helpers

import (
	"fmt"
	"github.com/samiulru/shebak/internal/config"
	"net/http"
	"runtime/debug"
)

var app *config.AppConfig

// NewHelpers sets up app config for helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with status of", status)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

//IsAuth checks user authenticity and return true if authorized
func IsAuth(r *http.Request) bool {
	exists := app.Session.Exists(r.Context(), "user_id")
	return exists
}
//IsSuperAdmin checks user level and return true if valid SuperAdmin
func IsSuperAdmin(r *http.Request) bool {
	accessLevel := app.Session.Get(r.Context(), "user_access_level").(int)
	return accessLevel < 2
}
//IsAdmin checks user level and return true if valid Admin
func IsAdmin(r *http.Request) bool {
	accessLevel := app.Session.Get(r.Context(), "user_access_level").(int)
	return accessLevel < 3
}
//IsEmployee checks user level and return true if valid Employee
func IsEmployee(r *http.Request) bool {
	accessLevel := app.Session.Get(r.Context(), "user_access_level").(int)
	return accessLevel < 4
}
//IsUser checks user level and return true if valid User
func IsUser(r *http.Request) bool {
	accessLevel := app.Session.Get(r.Context(), "user_access_level").(int)
	return accessLevel < 5	
}
