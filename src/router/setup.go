package router

import (
	"log"
	"net/http"
	"path"

	"github.com/bruh-boys/courses_platform/src/controllers"
	"github.com/gorilla/mux"
)

var TemplateDirectory = path.Join("src", "templates")

func SetupRouter(dirBase string, port string) (err error) {
	router := mux.NewRouter().StrictSlash(true)
	path := path.Join(dirBase, TemplateDirectory)

	if controllers.Templates, err = controllers.InitializeTemplates(path, controllers.TemplateFuncs); err != nil {

		return
	}

	// Static files.
	router.HandleFunc(`/public/{file:[\w\W\/]+}`, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	router.HandleFunc("/sign-in", controllers.SignIn)
	router.HandleFunc("/sign-up", controllers.SignUp)
	router.HandleFunc("/api", controllers.ApiInformation)

	router.HandleFunc("/new-post", controllers.NewPost)

	router.HandleFunc("/post", controllers.GetPost)
	router.HandleFunc("/", controllers.RenderHome)

	// Start server.
	log.Println("Starting server on port " + port + "...")

	return http.ListenAndServe((":" + port), router)
}
