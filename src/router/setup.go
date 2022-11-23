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

	// Initialize API routes.
	// router.HandleFunc("/api/auth/sign", controllers.Post) // DELETE - Sign out.
	// router.HandleFunc("/api/auth/sign", controllers.Post) // POST - Sign up.
	// router.HandleFunc("/api/auth/sign", controllers.Post) // PUT - Sign in.

	// router.HandleFunc("/api/post/{id:[0-9]+}", controllers.Post) // DELETE - Delete post.
	// router.HandleFunc("/api/post/{id:[0-9]+}", controllers.Post) // GET - Get post.

	// router.HandleFunc("/api/posts", controllers.Post) // GET - Get posts.

	// router.HandleFunc("/api/post", controllers.Post) // POST - Create post.

	// Initialize routes.
	// router.HandleFunc("/auth/signin", controllers.SignIn)
	// router.HandleFunc("/auth/signup", controllers.SignUp)

	// router.HandleFunc("/post/{id:[0-9]+}", controllers.Post)
	// router.HandleFunc("/", controllers.RenderHome)

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
