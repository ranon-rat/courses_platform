package tools

import (
	"math/rand"
	"net/http"
	"time"
)

func CreateCookie(name, value string, w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:  name,
		Value: value,
	})
}

func RemoveCookie(name string, w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:   name,
		Value:  "",
		MaxAge: -1,
	})
}

func GenerateToken() int {
	return rand.Int() + int(time.Now().Unix())
}
