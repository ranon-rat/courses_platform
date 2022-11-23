package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bruh-boys/courses_platform/src/core"
	"github.com/bruh-boys/courses_platform/src/db"
)

/*
GET /api?topics&topic=tal-cosa&page=1&size
*/
func ApiInformation(w http.ResponseWriter, r *http.Request) {
	var api core.ApiInformation
	values := r.URL.Query()
	topic := "any"

	if values.Has("topic") {
		topic = values.Get("topic")
	}

	if values.Has("topics") {
		api.Topics = (db.GetTopics())
	}

	if values.Has("page") {

		page, err := strconv.Atoi(values.Get("page"))
		if err != nil {
			http.Error(w, "that value is not an integer", 400)
			return
		}
		rows := db.GetPostsRows(page, topic)
		for rows.Next() {
			api.Posts = append(api.Posts, db.ScanRowPost(rows))
		}
	}

	if values.Has("size") {
		api.Quantity = db.PublicationsSize(topic)
	}

	json.NewEncoder(w).Encode(api)
}
