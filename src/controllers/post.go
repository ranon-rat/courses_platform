package controllers

import (
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/bruh-boys/courses_platform/src/db"
)

type Publication struct {
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

	file, _ := os.ReadFile("public/views/post.html")
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

	api.Logged = true

	ssid, err := r.Cookie("ssid")

	if err != nil {
		api.Logged = false
	} else if exist, _, _ := db.Existence(ssid.Value); exist == 0 {
		api.Logged = false
	}

	if api.Content == "" {
		api.Content = "Post not found"
	}

	//stri := strings.Replace(string(file), "!#!", post.Content, 1)
	template := template.Must(template.New("main").Parse(string(file)))
	template.Execute(w, api)

	//w.Write([]byte(stri))
}
