package middlewares

import (
	"net/http"

	"github.com/bruh-boys/courses_platform/src/db"
)

func Authenticated(w http.ResponseWriter, r *http.Request) bool {
	if ssid, err := r.Cookie("ssid"); err == nil {
		if priv, _, _ := db.GetSession(ssid.Value); priv > 0 || priv < 3 {

			return true
		}

	}

	return false
}
