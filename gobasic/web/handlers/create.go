package handlers

import (
	"gobasic/db"
	"gobasic/web/json_object"
	"gobasic/web/message"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {

	var newUser db.User

	if err := json_object.JsonDecoding(r, &newUser); err != nil {

		message.SendError(w, http.StatusBadRequest, "Failed to load json", "")
		return
	}

	if err := db.InsertUser(newUser); err != nil {

		message.SendError(w, http.StatusPreconditionFailed, err.Error(), "")
		return
	}

	message.SendJson(w, http.StatusCreated, newUser)

}
