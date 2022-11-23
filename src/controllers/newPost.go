package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bruh-boys/courses_platform/src/core"
	"github.com/bruh-boys/courses_platform/src/db"
	"github.com/bruh-boys/courses_platform/src/middlewares"
	"github.com/gomarkdown/markdown"
	"github.com/microcosm-cc/bluemonday"
)

func ParseContent(content string) string {
	maybeUnsafeHTML := markdown.ToHTML([]byte(content), nil, nil)
	html := bluemonday.UGCPolicy().SanitizeBytes(maybeUnsafeHTML)

	return string(html)
}

func NewPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

		return
	}

	var priv, id int
	var err error

	if ok := middlewares.Authenticated(w, r); !ok {
		http.Error(w, "You are not authorized to view this page.", http.StatusUnauthorized)

		return
	}

	ssid, _ := r.Cookie("ssid")

	if priv, id, err = db.GetSesion(ssid.Value); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println(err.Error())

		return
	}

	if priv < 1 || priv > 3 {
		http.Error(w, "You are not authorized to view this page.", http.StatusUnauthorized)

		return
	}

	var data core.ApiPostPublication

	// Decode the request body into the data variable.
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err.Error())

		return
	}

	if data.Content == "" || data.Title == "" || data.Introduction == "" {
		http.Error(w, "Missing fields data", http.StatusBadRequest)

		return
	}

	if data.Mineature == "" || data.Topic == "" {
		http.Error(w, "Missing fields data", http.StatusBadRequest)

		return
	}

	data.Content = ParseContent(data.Content)

	if err := db.CreatePost(data, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())

		return
	}

	w.WriteHeader(http.StatusCreated)
}
