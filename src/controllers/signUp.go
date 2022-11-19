package controllers

import (
	"encoding/json"
	"fmt"
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
		priv := core.Pupil
		if err == nil {
			_, priv, _ = db.Existence(ssid.Value)
		}
		if priv != core.Admin {
			sign.Privileges = core.Pupil
		}

		if err != nil || priv == core.Admin {

			if err = db.SignUp(sign); err != nil {
				http.Error(w, "username or email are already registered", http.StatusConflict)
				fmt.Println(err)
			}

		}

	default:
		http.Error(w, "Not implemented yet", http.StatusUnauthorized)
	}
}
