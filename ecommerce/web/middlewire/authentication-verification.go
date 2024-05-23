package middlewire

import (
	auth "ecommerce/Auth"
	"ecommerce/web/messages"
	"net/http"
)

func Authenticate(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if err := auth.CheckAuthorization(token); err != nil {
			messages.SendError(w, http.StatusUnauthorized, err.Error(), "")
			return
		}
		h.ServeHTTP(w, r)
	})
}
