package web

import (
	"ecommerce/web/handlers"
	"ecommerce/web/middlewire"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, manager *middlewire.Manager) {
	mux.Handle(
		"POST /users/register",
		manager.With(
			http.HandlerFunc(handlers.RegisterUser),
		),
	)

	mux.Handle(
		"POST /users/login",
		manager.With(
			http.HandlerFunc(handlers.UserLogin),
		),
	)

	mux.Handle(
		"POST /users/addproduct",
		manager.With(
			http.HandlerFunc(handlers.AddProduct),
		),
	)

	mux.Handle(
		"POST /users/addorder",
		manager.With(
			http.HandlerFunc(handlers.NewOrder),
		),
	)

	mux.Handle(

		"GET /users/search",
		manager.With(
			http.HandlerFunc(handlers.SearchProduct), []middlewire.Middleware{middlewire.Authenticate}...,
		),
	)

}
