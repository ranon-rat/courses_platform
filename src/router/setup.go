package router

import (
	"log"
	"net/http"
	"os"

	"github.com/bruh-boys/courses_platform/src/controllers"
)

func SetupRouter() error {
	http.HandleFunc("/sign-in", controllers.SignIn)
	port, exist := os.LookupEnv("PORT")
	if !exist {
		log.Println("debuggin mode,working in the port 8080")
		port = "8080"
	}
	return http.ListenAndServe(":"+port, nil)
}
