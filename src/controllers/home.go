package controllers

import (
	"net/http"
	"strconv"

	"github.com/bruh-boys/courses_platform/src/core"
	"github.com/bruh-boys/courses_platform/src/db"
	"github.com/bruh-boys/courses_platform/src/middlewares"
)

func RenderHome(w http.ResponseWriter, r *http.Request) {
	var api core.ApiInformation

	if middlewares.Authenticated(w, r) {
		api.Logged = true
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
	/*file, _ := os.ReadFile("templates/home.html")
	template := template.Must(template.New("html").Funcs(tmpFuncs).Parse(string(file)))

	template.Execute(w, api)*/

	if err := Templates.ExecuteTemplate(w, "Home", api); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

}
