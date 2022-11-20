package controllers

import "text/template"

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
}
