package middlewire

import (
	auth "ecommerce/Auth"
	"ecommerce/web/messages"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (m Manager) Use(middlewares ...Middleware) Manager {
	m.globalMiddlewares = append(m.globalMiddlewares, middlewares...)
	return m
}

func (m *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
	var h http.Handler
	h = handler

	for _, m := range middlewares {
		h = m(h)
	}

	for _, m := range m.globalMiddlewares {
		h = m(h)
	}

	return h
}

func (m *Manager) Authenticate(handler http.Handler, middlewares ...Middleware) http.Handler {
	var h http.Handler
	h = handler

	for _, m := range middlewares {
		h = m(h)
	}

	for _, m := range m.globalMiddlewares {
		h = m(h)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract JWT token from request header
		token := r.Header.Get("Authorization")

		// Verify the JWT token
		if err := auth.CheckAuthorization(token); err != nil {
			messages.SendError(w, http.StatusUnauthorized, err.Error(), "")
			return
		}

		// If token is valid, call the original handler
		h.ServeHTTP(w, r)
	})
}
