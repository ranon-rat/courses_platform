package router

import (
	"log"
	"net/http"
	"os"

	"github.com/bruh-boys/courses_platform/src/controllers"
	"github.com/gorilla/mux"
)

func SetupRouter() error {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/sign-in", controllers.SignIn)
	router.HandleFunc("/sign-up", controllers.SignUp)
	router.HandleFunc("/api", controllers.ApiInformation)
	router.HandleFunc("/new-post", controllers.NewPost)

	router.HandleFunc(`/public/{file:[\w\W\/]+}`, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	router.HandleFunc("/post", controllers.GetPost)
	router.HandleFunc("/", controllers.RenderHome)

	port, exist := os.LookupEnv("PORT")
	if !exist {
		log.Println("debuggin mode,working in the port 3000")
		port = "3000"
	}

	return http.ListenAndServe(":"+port, router)
}
