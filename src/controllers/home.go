package controllers

import (
	"html/template"
	"net/http"
	"os"
	"strconv"

	"github.com/bruh-boys/courses_platform/src/core"
	"github.com/bruh-boys/courses_platform/src/db"
)

var tmp = template.New("tmp")

func Setup() {
	tmp.Funcs(template.FuncMap{
		"loop": func(from, to int) []int {
			list := []int{}
			for i := from; i <= to; i++ {
				list = append(list, i)
			}
			return list
		},
	})
}
func RenderHome(w http.ResponseWriter, r *http.Request) {
	var api core.ApiInformation
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
