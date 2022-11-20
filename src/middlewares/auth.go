package middlewares

import (
	"net/http"

	"github.com/bruh-boys/courses_platform/src/db"
	"github.com/bruh-boys/courses_platform/src/tools"
)

func Authenticated(w http.ResponseWriter, r *http.Request) bool {
	if ssid, err := r.Cookie("ssid"); err == nil {
		if priv, _, _ := db.Existence(ssid.Value); priv > 0 || priv < 4 {

			return true
		}

	}

	tools.RemoveCookie("ssid", w)

	return false
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !Authenticated(w, r) {
			http.Error(w, "You are not authorized to view this page.", http.StatusUnauthorized)

			return
		}

		next.ServeHTTP(w, r)
	})
}
