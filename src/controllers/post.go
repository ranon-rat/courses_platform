package controllers

import (
	"net/http"
	"strconv"

	"github.com/bruh-boys/courses_platform/src/db"
	"github.com/bruh-boys/courses_platform/src/middlewares"
)

type Publication struct {
	Admin        bool
	Logged       bool
	ID           string
	Content      string
	Title        string
	Mineature    string
	Author       string
	Date         int
	Topic        string
	Introduction string
}

func GetPost2(w http.ResponseWriter, r *http.Request) {

}

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	values := r.URL.Query()

	if !values.Has("id") {
		values.Set("id", "1")
	}

	id, err := strconv.Atoi(values.Get("id"))

	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)

		return
	}

	post := db.GetPost(id)

	var api = Publication{
		ID:           post.ID,
		Content:      post.Content,
		Title:        post.Title,
		Mineature:    post.Mineature,
		Author:       post.Author,
		Date:         post.Date,
		Topic:        post.Topic,
		Introduction: post.Introduction,
	}

	if middlewares.Authenticated(w, r) {
		api.Logged = true
	}

	if ssid, err := r.Cookie("ssid"); err == nil {
		priv := 0

		if priv, _, err = db.GetSession(ssid.Value); err != nil {
			http.Error(w, "Internal server error.", http.StatusInternalServerError)

			return
		}

		if priv > 0 && priv < 3 {
			api.Admin = true
		}
	}

	if api.Content == "" {
		api.Content = "Post not found"
	}

	if err := Templates.ExecuteTemplate(w, "Post", api); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
}
