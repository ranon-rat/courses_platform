package middlewares

import (
	"fmt"
	"net/http"

	"github.com/bruh-boys/courses_platform/src/db"
)

func Authenticated(w http.ResponseWriter, r *http.Request) bool {
	if ssid, err := r.Cookie("ssid"); err == nil {
		if priv, _, _ := db.GetSesion(ssid.Value); priv > 0 || priv < 3 {
			fmt.Println("Auth middleware 12")
			return true
		}

	}

	return false
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Auth middleware 2")
		if !Authenticated(w, r) {
			http.Error(w, "You are not authorized to view this page.", http.StatusUnauthorized)

			return
		}

		next.ServeHTTP(w, r)
	})

}
