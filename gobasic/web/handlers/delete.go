package handlers

import (
	"encoding/json"

	"gobasic/db"
	"gobasic/web/message"
	"net/http"
)

type UserData struct {
	ID int `db:"id" json:"id"`
}

func Delete(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodDelete {

		var delete_user UserData
		err := json.NewDecoder(r.Body).Decode(&delete_user)

		if err != nil {

			message.Send_Error(w, http.StatusBadRequest, "Failed to load json", "")
			return
		}

		status, value := db.DeleteUser(db.UserData(delete_user))
		if !status {
			message.Send_Json(w, http.StatusPreconditionFailed, value)
			return
		} else {
			message.Send_Json(w, http.StatusAccepted, value)
			return
		}

	}

}
