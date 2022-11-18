package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/bruh-boys/courses_platform/src/core"
	"github.com/bruh-boys/courses_platform/src/db"
)

func RenderHome(w http.ResponseWriter, r *http.Request) {
	var api core.ApiInformation
	values := r.URL.Query()
	topic := "any"

	if values.Has("topic") {
		topic = values.Get("topic")

	}

	if values.Has("topics") {
		api.Topics = (db.GetTopics())

	}

	page, err := strconv.Atoi(values.Get("page"))
	if err != nil {
		page = 1
	}
	api.Posts = (db.GetPosts(page, topic))
	fmt.Println(api.Posts)

	if values.Has("size") {
		api.Quantity = db.PublicationsSize(topic)
	}
	file, _ := os.ReadFile("public/views/home.html")

	template := template.Must(template.New("main").Parse(string(file)))
	template.Execute(w, api)

}
