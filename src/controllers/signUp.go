package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/bruh-boys/courses_platform/src/core"
	"github.com/bruh-boys/courses_platform/src/db"
)

// Optimized.
func SignUp(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)

	rand.Seed(time.Now().Unix())

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

		return
	}

	var data core.SignUp

	// Decode the request body into the data variable.
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	// Check if the email is already registered.
	if ok, _ := IsAlreadyRegistered(w, data.Email); !ok {
		http.Error(w, "Email is already registered", http.StatusNotFound)

		return
	}

	// Create the user in the database.
	if err := db.SignUp(data); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		log.Println(err.Error())

		return
	}
	fmt.Println(data)
	// Sign in the user.

	if err := db.SignUp(data); err != nil {
		http.Error(w, "username or email are already registered", http.StatusConflict)
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusCreated)
}

/*
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
*/
