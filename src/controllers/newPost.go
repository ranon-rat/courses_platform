package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/bruh-boys/courses_platform/src/core"
	"github.com/bruh-boys/courses_platform/src/db"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

func Get(wr http.ResponseWriter, r *http.Request) {

}

func Add(wr http.ResponseWriter, r *http.Request) {

}

func NewPost(w http.ResponseWriter, r *http.Request) {

	ssid, err := r.Cookie("ssid")
	if err != nil {
		http.Error(w, "you doesnt have an account", http.StatusUnauthorized)
		return
	}

	priv, id := db.Existence(ssid.Value)
	if priv == 0 {
		http.Error(w, "you doesnt have an account", http.StatusNotFound)
		return
	}

	if priv == core.Pupil {
		http.Error(w, "you cant do that", http.StatusUnauthorized)
		return
	}

	switch r.Method {
	case "POST":
		var post core.ApiPostPublication

		if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
			http.Error(w, "error parsing the body request", http.StatusBadRequest)
			return
		}

		post.Content = ParseContent(post.Content)

		db.NewPost(post, id)
	default:

		http.Error(w, "not implented yet", 400)
	}
}

func ParseContent(content string) (body string) {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	parse := template.HTMLEscapeString(content)

	body = string(markdown.ToHTML([]byte(parse), parser, nil))
	return
}
