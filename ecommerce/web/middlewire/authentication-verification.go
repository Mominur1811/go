package middlewire

import (
	auth "ecommerce/Auth"
	"ecommerce/web/utils"
	"encoding/json"
	"net/http"
)

type RefreshToken struct {
	RefreshToken string `json:"refreshtoken"`
}

func AuthenticateAccessToken(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if err := auth.CheckAuthorization(token); err != nil {
			utils.SendError(w, http.StatusUnauthorized, err.Error())
			return
		}
		h.ServeHTTP(w, r)
	})
}

func AuthenticateRefreshToken(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var token RefreshToken
		if err := json.NewDecoder(r.Body).Decode(&token); err != nil {
			utils.SendError(w, http.StatusBadRequest, err.Error())
			return
		}

		var err error
		if err = auth.CheckAuthorization(token.RefreshToken); err != nil {
			utils.SendError(w, http.StatusUnauthorized, err.Error())
			return
		}
		h.ServeHTTP(w, r)
	})
}
