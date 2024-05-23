package handlers

import (
	auth "ecommerce/Auth"
	"ecommerce/db"
	"ecommerce/web/messages"
	"encoding/json"
	"net/http"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {

	var userLogin db.Login

	if err := json.NewDecoder(r.Body).Decode(&userLogin); err != nil {

		messages.SendError(w, http.StatusExpectationFailed, err.Error(), "")
		return
	}

	//User Login Validation
	if err := db.LoginUser(userLogin); err != nil {

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
