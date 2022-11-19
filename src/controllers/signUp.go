package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/bruh-boys/courses_platform/src/core"
	"github.com/bruh-boys/courses_platform/src/db"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	switch r.Method {
	case "POST":
		var sign core.SignUp

		if err := json.NewDecoder(r.Body).Decode(&sign); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		ssid, err := r.Cookie("ssid")

		_, priv, _ := db.Existence(ssid.Value)

		if priv != core.Admin {
			sign.Privileges = core.Pupil
		}

		if sign.Privileges < core.Admin || sign.Privileges > core.Pupil {
			http.Error(w, "unkown kind of user", http.StatusBadRequest)
		}

		if err != nil || priv == core.Admin {

			if err = db.SignUp(sign); err != nil {
				http.Error(w, "username or email are already registered", http.StatusConflict)
			}

		}

	default:
		http.Error(w, "Not implemented yet", http.StatusUnauthorized)
	}
}
