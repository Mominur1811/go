package web

import (
	"JwtToken/web/handlers"
	"JwtToken/web/middlewire"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, manager *middlewire.Manager) {
	mux.Handle(
		"POST /users/register",
		manager.With(
			http.HandlerFunc(handlers.CreateNewUser),
		),
	)

	mux.Handle(
		"POST /users/login",
		manager.With(
			http.HandlerFunc(handlers.UserLogin),
		),
	)

	mux.Handle(
		"POST /users/verify",
		manager.With(
			http.HandlerFunc(handlers.JwtVerification),
		),
	)

}
