package web

import (
	"gobasic/web/handlers"
	"gobasic/web/middlewires"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, manager *middlewires.Manager) {
	mux.Handle(
		"GET /users",
		manager.With(
			http.HandlerFunc(handlers.View),
		),
	)

	mux.Handle(
		"POST /users",
		manager.With(
			http.HandlerFunc(handlers.Create),
		),
	)

	mux.Handle(
		"PUT /users",
		manager.With(
			http.HandlerFunc(handlers.Update),
		),
	)
	mux.Handle(
		"DELETE /users",
		manager.With(
			http.HandlerFunc(handlers.Delete),
		),
	)
}
