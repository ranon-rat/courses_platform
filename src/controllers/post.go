package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bruh-boys/courses_platform/src/db"
)

func GetPost(w http.ResponseWriter, r *http.Request) {

	if r.URL.Query().Has("id") {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		post := db.GetPost(id)
		fmt.Println([]byte(post.Content))
		json.NewEncoder(w).Encode(post)
		return
	}
	http.Error(w, "bad request", http.StatusBadRequest)
}
