package infra

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HTTPRoute struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

func StartHTTPServer(routes []HTTPRoute, port string) {
	r := chi.NewRouter()
	registerMiddlewares(r)
	registerRoutes(r, routes)
	http.ListenAndServe(":"+port, r)
}

func registerRoutes(r *chi.Mux, routes []HTTPRoute) {
	for _, route := range routes {
		switch route.Method {
		case "GET":
			r.Get(route.Path, route.Handler)
		case "POST":
			r.Post(route.Path, route.Handler)
		}
	}
}

func registerMiddlewares(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

}
