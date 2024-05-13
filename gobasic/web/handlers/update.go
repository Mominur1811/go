package handlers

import (
	"gobasic/db"
	"gobasic/web/json_object"
	"gobasic/web/message"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {

	var updateUser db.User
	if err := json_object.JsonDecoding(r, &updateUser); err != nil {

		message.SendError(w, http.StatusBadRequest, "Failed to load json", "")
		return
	}

	// Checking if the user exist or not before deleting
	if value := db.CheckExistenseOfUser(updateUser.Id); value != nil {
		message.SendError(w, http.StatusBadRequest, value.Error(), "")
		return
	}

	//Calling delete function
	if value := db.UpdateUser(updateUser); value != nil {
		message.SendError(w, http.StatusBadRequest, value.Error(), "")
		return
	}
	message.SendJson(w, http.StatusAccepted, updateUser)

}
