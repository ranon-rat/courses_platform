package controllers

import (
	"net/http"
	"strconv"

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

	api.Logged = false

	ssid, err := r.Cookie("ssid")

	if err == nil {
		if exist, _, _ := db.Existence(ssid.Value); exist > 0 {
			api.Logged = true
		} else {
			r.AddCookie(&http.Cookie{
				Name:   "ssid",
				Value:  "",
				MaxAge: -1,
			})
		}
	}

	if api.Content == "" {
		api.Content = "Post not found"
	}

	if err := Templates.ExecuteTemplate(w, "Post", api); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	//stri := strings.Replace(string(file), "!#!", post.Content, 1)
	//template := template.Must(template.New("main").Parse(string(file)))
	//template.Execute(w, api)

	//w.Write([]byte(stri))
}
