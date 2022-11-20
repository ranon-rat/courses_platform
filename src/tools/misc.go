package tools

import (
	"math/rand"
	"net/http"
	"time"
)

func RemoveCookie(name string, w http.ResponseWriter) {
	w.Header().Add("Set-Cookie", (name + "=; Max-Age=-1"))
}

func GenerateToken() int {
	return rand.Int() + int(time.Now().Unix())
}
