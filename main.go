package main

import (
	"log"
	"os"

	"github.com/bruh-boys/courses_platform/src/router"
)

func main() {
	var port string
	var dir string
	var err error

	if port = os.Getenv("PORT"); port == "" {
		log.Println("No PORT environment variable detected. Setting to default 3000.")

		port = "3000"
	}

	if dir, err = os.Getwd(); err != nil {
		log.Println("Error getting working directory.")
		log.Panicln(err)
	}

	if err = router.SetupRouter(dir, port); err != nil {
		log.Panicln(err)
	}

}
