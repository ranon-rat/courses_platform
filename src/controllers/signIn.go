package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/bruh-boys/courses_platform/src/core"
	"github.com/bruh-boys/courses_platform/src/db"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())

	switch r.Method {
	case "POST":

		var sign core.SignIn
		if err := json.NewDecoder(r.Body).Decode(&sign); err != nil {
			http.Error(w, "something is wrong", http.StatusBadRequest)
			return
		}
		exist := 0
		ssid, err := r.Cookie("ssid")

		if err == nil {
			exist, _, _ = db.Existence(ssid.Value)

		}
		if exist == 0 || err != nil {
			if db.ExistenceWithPass(sign) > 0 {
				ssid := db.SignIn(sign)

				cookie := &http.Cookie{
					Name:    "ssid",
					Value:   ssid,
					Expires: time.Now().AddDate(1, 0, 0),
				}
				http.SetCookie(w, cookie)
				return
			}
			http.Error(w, "wrong password or the account does not exist", http.StatusBadRequest)

			return
		}
		http.Error(w, "you are already logged in", http.StatusBadRequest)

	default:
		http.Error(w, "not implemented yet", 400)
	}

}
