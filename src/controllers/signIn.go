package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bruh-boys/courses_platform/src/core"
	"github.com/bruh-boys/courses_platform/src/db"
	"github.com/bruh-boys/courses_platform/src/tools"
)

func IsAlreadyRegistered(w http.ResponseWriter, email string) (ok bool, err error) {
	if ok, err = db.IsAlreadyRegistered(email); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	return
}

func IsPasswordCorrect(w http.ResponseWriter, email, password string) (ok bool, err error) {
	if ok, err = db.IsPasswordCorrect(email, password); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	return
}

func SignInUser(w http.ResponseWriter, data core.SignIn) (ssid string, err error) {
	if ssid, err = db.SignIn(data); err != nil {

		return "", err
	}

	tools.CreateCookie("ssid", ssid, w)

	return
}

//

func SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

		return
	}

	var data core.SignIn

	// Decode the request body into the data variable.
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err.Error())

		return
	}

	// Check if the email is already registered.
	if ok, _ := IsAlreadyRegistered(w, data.Email); !ok {
		http.Error(w, "Email is not registered", http.StatusNotFound)

		return
	}

	// Check if the password is correct.
	if ok, _ := IsPasswordCorrect(w, data.Email, data.Password); !ok {
		http.Error(w, "Password is incorrect", http.StatusUnauthorized)

		return
	}

	// Sign in the user.
	if _, err := SignInUser(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("SignInUser:", err)

		return
	}

	w.WriteHeader(http.StatusOK)
}
