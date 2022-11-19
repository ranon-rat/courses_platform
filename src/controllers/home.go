package controllers

import (
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/bruh-boys/courses_platform/src/core"
	"github.com/bruh-boys/courses_platform/src/db"
)

var tmp = template.New("tmp")

func Setup() {
	tmp.Funcs(template.FuncMap{
		"loop": func(from, to int) <-chan int {
			ch := make(chan int)
			go func() {
				for i := from; i <= to; i++ {
					ch <- i
				}
				close(ch)
			}()
			return ch
		},
	})
}
func RenderHome(w http.ResponseWriter, r *http.Request) {
	var api core.ApiInformation
	api.Logged = true

	ssid, err := r.Cookie("ssid")

	if err != nil {
		api.Logged = false
	} else if priv, _ := db.Existence(ssid.Value); priv == 0 {
		api.Logged = false
	}

	values := r.URL.Query()
	topic := "any"

	if values.Has("topic") {
		topic = values.Get("topic")

	}

	api.Topics = db.GetTopics()

	page, err := strconv.Atoi(values.Get("page"))
	if err != nil {
		page = 1
	}
	api.Posts = (db.GetPosts(page, topic))
	api.Quantity = db.PublicationsSize(topic)
	to := ((api.Quantity) / core.PostPerPage)
	if to > 8 {
		to = page + 8
	}
	// ye
	api.Page = page
	api.To = to + 1
	file, _ := os.ReadFile("public/views/home.html")
	template := template.Must(tmp.Parse(string(file)))

	template.Execute(w, api)

}
