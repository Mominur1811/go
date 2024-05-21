package handlers

import (
	"JwtToken/db"
	messages "JwtToken/web/message"
	"encoding/json"
	"net/http"
)

func CreateNewUser(w http.ResponseWriter, r *http.Request) {

	var newEntry db.User

	if err := json.NewDecoder(r.Body).Decode(&newEntry); err != nil {

		messages.SendError(w, http.StatusBadRequest, "Failed to load json", "")
		return
	}

	if err := db.InsertNewUser(newEntry); err != nil {

		messages.SendError(w, http.StatusPreconditionFailed, err.Error(), "")
		return
	}

	messages.SendJson(w, http.StatusCreated, newEntry)

}
