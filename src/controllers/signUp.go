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

// Optimized.
func SignUp(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

		return
	}

	var data core.SignUp
	data.Privileges = 3

	// Decode the request body into the data variable.
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	// Check if the email is already registered.
	if ok, _ := IsAlreadyRegistered(w, data.Email); ok {
		http.Error(w, "Email is already registered", http.StatusConflict)

		return
	}

	// Create the user in the database.
	if err := db.SignUp(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())

		return
	}

	// Sign in the user.
	if _, err := SignInUser(w, core.SignIn{
		Password: data.Password,
		Email:    data.Email,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("SignInUser:", err)

		return
	}

	w.WriteHeader(http.StatusCreated)
}
