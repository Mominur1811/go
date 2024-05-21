package handlers

import (
	auth "JwtToken/Auth"
	messages "JwtToken/web/message"
	"encoding/json"
	"net/http"
)

func JwtVerification(w http.ResponseWriter, r *http.Request) {

	var token map[string]string
	if err := json.NewDecoder(r.Body).Decode(&token); err != nil {

		messages.SendError(w, http.StatusForbidden, err.Error()+" hihi", "")
		return
	}

	payload, err := auth.GetVerifyJwt(token)
	if err != nil {
		messages.SendError(w, http.StatusUnauthorized, err.Error(), "")
		return
	}

	messages.SendData(w, "Authorized", payload)

}
