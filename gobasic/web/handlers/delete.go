package handlers

import (
	"fmt"

	"gobasic/db"
	"gobasic/web/json_object"
	"gobasic/web/message"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {

	var deleteUserId db.UserId
	if err := json_object.JsonDecoding(r, &deleteUserId); err != nil {

		message.SendError(w, http.StatusBadRequest, err.Error(), "")
		return
	}

	// Checking if the user exist or not before deleting
	if value := db.CheckExistenseOfUser(deleteUserId.ID); value != nil {
		message.SendError(w, http.StatusBadRequest, value.Error(), "")
		return
	}

	//Calling delete function
	if value := db.DeleteUser(deleteUserId); value != nil {
		message.SendError(w, http.StatusBadRequest, value.Error(), "")
		return
	}
	message.SendJson(w, http.StatusAccepted, fmt.Sprint(deleteUserId)+" deleted")

}
