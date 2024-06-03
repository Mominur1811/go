package web

import (
	handleruser "ecommerce/web/handler-user"
	"ecommerce/web/handlers"
	"ecommerce/web/middlewire"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, manager *middlewire.Manager) {
	mux.Handle(
		"POST /users/register",
		manager.With(
			http.HandlerFunc(handleruser.SignUp),
		),
	)

	mux.Handle(
		"POST /users/login",
		manager.With(
			http.HandlerFunc(handleruser.UserLogin),
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
			http.HandlerFunc(handlers.SearchProduct),
		),
	)

	mux.Handle(

		"GET /users/newtoken",
		manager.With(
			http.HandlerFunc(handleruser.RefreshAccessToken), middlewire.AuthenticateRefreshToken,
		),
	)

	mux.Handle(
		"POST /users/updaterequest",
		manager.With(
			http.HandlerFunc(handleruser.ResetPass), middlewire.AuthenticateAccessToken,
		),
	)

}
