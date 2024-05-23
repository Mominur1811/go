package handlers

import (
	"ecommerce/db"
	"ecommerce/web/messages"
	"encoding/json"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	var newUser db.User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {

		messages.SendError(w, http.StatusPreconditionFailed, err.Error(), "")
		return
	}

	if err := db.InsertNewUser(newUser); err != nil {
		messages.SendError(w, http.StatusExpectationFailed, err.Error(), "")
		return
	}

	messages.SendData(w, "", newUser)

}
