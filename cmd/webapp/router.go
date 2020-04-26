package webapp

import (
	"github.com/go-chi/chi"
)

type Router struct {
	*chi.Mux
}

func (r *Router) AddRoutes() {
	router := chi.NewRouter()

	router.Get("/", home)

	r.Mount("/", router)
}

func (r *Router) AddApiRoutes() {
	router := chi.NewRouter()

	router.Post("/shorten", shorten)

	r.Mount("/api", router)
}

func NewRouter() Router {
	r := chi.NewRouter()
	return Router{r}
}
