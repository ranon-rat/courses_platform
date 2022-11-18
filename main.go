package main

import (
	"log"

	"github.com/bruh-boys/courses_platform/src/router"
)

func main() {
	//fmt.Println(db.Hash("1234"))

	log.Println(router.SetupRouter())
}
