package handlers

import (
	"encoding/json"
	"gobasic/db"
	"gobasic/web/message"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPut {
		var update_user User
		err := json.NewDecoder(r.Body).Decode(&update_user)

		if err != nil {

			http.Error(w, "Error user to update", http.StatusBadRequest)
			return
		}

		status, value := db.UpdateUser(db.User(update_user))

		if !status {
			message.Send_Json(w, http.StatusPreconditionFailed, value)
			return
		} else {
			message.Send_Json(w, http.StatusAccepted, value)
			return
		}

	}
}
