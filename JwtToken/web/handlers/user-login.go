package handlers

import (
	auth "JwtToken/Auth"
	"JwtToken/db"
	messages "JwtToken/web/message"
	"encoding/json"
	"net/http"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {

	var userLogin db.LoginData

	if err := json.NewDecoder(r.Body).Decode(&userLogin); err != nil {

		messages.SendError(w, http.StatusExpectationFailed, err.Error(), "")
		return
	}

	//User Login Validation
	if err := db.UserLoginValidation(userLogin); err != nil {

		messages.SendError(w, http.StatusExpectationFailed, err.Error(), "")
		return
	}

	//Get Jwt Token
	jwtToken, err := auth.GetJwtToken(userLogin)
	if err != nil {

		messages.SendError(w, http.StatusExpectationFailed, err.Error(), "")
		return
	}

	messages.SendData(w, http.StatusOK, jwtToken)
}
