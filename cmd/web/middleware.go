package main

import (
	"github.com/justinas/nosurf"
	"github.com/samiulru/shebak/internal/helpers"
	"net/http"
)

// NoSurf adds CSRF protection to all post request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads and saves the session in every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

//AuthSuperAdmin is the super admin authenticator
func AuthSuperAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !helpers.IsAuth(r){
			session.Put(r.Context(), "error", "please log in first!")
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			return
		}
		if !helpers.IsSuperAdmin(r){
			session.Put(r.Context(), "error", "Invalid credentials: Unauthorized Super-Admin\n!Please try again")
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			return
		}
		session.Put(r.Context(), "flash", "Logged in as SuperAdmin")
		session.Put(r.Context(), "userType", "SuperAdmin")
		next.ServeHTTP(w, r)
	})
}


//AuthAdmin is the admin authenticator
func AuthAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !helpers.IsAuth(r){
			session.Put(r.Context(), "error", "please log in first!")
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			return
		}
		if !helpers.IsAdmin(r){
			session.Put(r.Context(), "error", "Invalid credentials: Unauthorized admin!Please try again")
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			return
		}
		session.Put(r.Context(), "flash", "Logged in as Admin")
		session.Put(r.Context(), "userType", "Admin")
		next.ServeHTTP(w, r)
	})

}

//AuthEmployee is the Employee authenticator
func AuthEmployee(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !helpers.IsAuth(r){
			session.Put(r.Context(), "error", "please log in first!")
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			return
		}
		if !helpers.IsEmployee(r){
			session.Put(r.Context(), "error", "Invalid credentials: Unauthorized Employee!Please try again")
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			return
		}
		session.Put(r.Context(), "flash", "Logged in as Employee")
		session.Put(r.Context(), "userType", "Employee")
		next.ServeHTTP(w, r)
	})

}

//AuthUser is the Employee authenticator
func AuthUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !helpers.IsAuth(r){
			session.Put(r.Context(), "error", "please log in first!")
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			return
		}
		if !helpers.IsUser(r){
			session.Put(r.Context(), "error", "Invalid credentials: Unauthorized User!Please try again")
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			return
		}
		session.Put(r.Context(), "flash", "Logged in successfully ")
		session.Put(r.Context(), "userType", "User")
		next.ServeHTTP(w, r)
	})

}

