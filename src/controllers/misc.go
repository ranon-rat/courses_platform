package controllers

import (
	"text/template"

	"github.com/bruh-boys/courses_platform/src/core"
	"github.com/bruh-boys/courses_platform/src/db"
)

var Templates = template.New("")

var TemplateFuncs = template.FuncMap{
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
	"posts": func(page int, topic string) <-chan core.ApiGetPublication {
		ch := make(chan core.ApiGetPublication)
		rows := db.GetPostsRows(page, topic)
		go func() {
			for rows.Next() {
				ch <- db.ScanRowPost(rows)
			}
			close(ch)
		}()
		return ch
	},
}
