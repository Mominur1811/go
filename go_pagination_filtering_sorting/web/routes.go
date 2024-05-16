package web

import (
	"net/http"

	"example.com/go_pagination_filtering_sorting/web/handlers"
	"example.com/go_pagination_filtering_sorting/web/middlewire"
)

func InitRoutes(mux *http.ServeMux, manager *middlewire.Manager) {
	mux.Handle(
		"GET /users",
		manager.With(
			http.HandlerFunc(handlers.Search),
		),
	)

	mux.Handle(
		"POST /users",
		manager.With(
			http.HandlerFunc(handlers.Insert),
		),
	)
}
