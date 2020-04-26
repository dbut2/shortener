package webapp

import (
	"github.com/go-chi/chi"
)

type Router struct {
	*chi.Mux
}

func (r *Router) AddRoutes() {
	router := chi.NewRouter()

	router.Get("/", file)
	router.Get("/index.html", file)
	router.Get("/static/*", file)

	r.Mount("/", router)
}

func (r *Router) AddApiRoutes() {
	router := chi.NewRouter()

	router.Post("/shorten", shorten)

	r.Mount("/api", router)
}

func (r *Router) AddLengthenRouter() {
	router := chi.NewRouter()

	router.Get("/{code}", lengthen)

	r.Mount("/ort", router)
}

func NewRouter() Router {
	r := chi.NewRouter()
	return Router{r}
}
