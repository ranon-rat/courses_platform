package router

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/bruh-boys/courses_platform/src/controllers"
	"github.com/gomarkdown/markdown"
)

func SetupRouter() error {
	http.HandleFunc("/sign-in", controllers.SignIn)
	http.HandleFunc("/sign-up", controllers.SignUp)
	http.HandleFunc("/api", controllers.ApiInformation)
	http.HandleFunc("/new-post", controllers.NewPost)
	http.HandleFunc("/post", controllers.GetPost)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file, _ := os.ReadFile("public/index.html")
		md, _ := os.ReadFile("public/test.md")
		out := markdown.ToHTML(md, nil, nil)

		stri := strings.Replace(string(file), "!#!", string(out), 1)

		w.Write([]byte(stri))
	})

	port, exist := os.LookupEnv("PORT")
	if !exist {
		log.Println("debuggin mode,working in the port 8080")
		port = "8080"
	}
	return http.ListenAndServe(":"+port, nil)
}
