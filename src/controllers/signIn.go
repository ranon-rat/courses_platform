package controllers

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/bruh-boys/courses_platform/src/core"
	"github.com/bruh-boys/courses_platform/src/db"
)

func IsAlreadyRegistered(w http.ResponseWriter, email string) (ok bool, err error) {
	if ok, err = db.IsAlreadyRegistered(email); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	if !ok {
		http.Error(w, "Email is not registered", http.StatusNotFound)

		return
	}

	return
}

func IsPasswordCorrect(w http.ResponseWriter, email, password string) (ok bool, err error) {
	if ok, err = db.IsPasswordCorrect(email, password); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	if !ok {
		http.Error(w, "Password is incorrect", http.StatusUnauthorized)

		return
	}

	return
}

func SignIn2(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

		return
	}

	var data core.SignIn

	// Decode the request body into the data variable.
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	// Check if the email is already registered.
	if ok, err := IsAlreadyRegistered(w, data.Email); err != nil || !ok {
		log.Println(err)

		return
	}

	// Check if the password is correct.
	if ok, err := IsPasswordCorrect(w, data.Email, data.Password); err != nil || !ok {
		log.Println(err)

		return
	}

	w.WriteHeader(http.StatusOK)
}

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
