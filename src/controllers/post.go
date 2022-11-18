package controllers

import (
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/bruh-boys/courses_platform/src/db"
)

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

	if post.Content == "" {
		post.Content = "Post not found"
	}

	//stri := strings.Replace(string(file), "!#!", post.Content, 1)
	template := template.Must(template.New("main").Parse(string(file)))
	template.Execute(w, post)

	//w.Write([]byte(stri))
}
