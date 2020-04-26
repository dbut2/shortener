package main

import (
	"github.com/dbut2/shortener/cmd/webapp"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	start(port)
}

func start(port string) {
	router := webapp.NewRouter()

	router.AddRoutes()
	router.AddApiRoutes()
	router.AddLengthenRouter()

	http.ListenAndServe(":"+port, router)
}
