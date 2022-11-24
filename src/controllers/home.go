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

	var err error = nil

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

	values := r.URL.Query()
	topic := "any"

	if !values.Has("id") {
		http.Redirect(w, r, "/?id=1", http.StatusTemporaryRedirect)

		return
	}

	if values.Has("topic") {
		topic = values.Get("topic")

	}

	api.Topics = db.GetTopics()

	page, err := strconv.Atoi(values.Get("page"))
	if err != nil {
		page = 1
	}
	api.Quantity = db.PublicationsSize(topic)
	to := ((api.Quantity) / core.PostPerPage)
	if to > 8 {
		to = page + 8
	}
	// ye
	api.Page = page
	api.To = to + 1
	api.Topic = topic

	/*file, _ := os.ReadFile("templates/home.html")
	template := template.Must(template.New("html").Funcs(tmpFuncs).Parse(string(file)))

	template.Execute(w, api)*/

	if err := Templates.ExecuteTemplate(w, "Home", api); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

}
