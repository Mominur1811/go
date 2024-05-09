package handlers

import (
	"encoding/json"
	"gobasic/db"
	"gobasic/web/message"
	"net/http"
)

type User struct {
	ID       int    `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Password string `db:"password" json:"password"`
}

func Create(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		var new_user User
		err := json.NewDecoder(r.Body).Decode(&new_user)

		if err == nil {

			err = db.InsertUser(db.User(new_user))
			if err != nil {

				message.Send_Json(w, http.StatusPreconditionFailed, err)
				return
			} else {

				message.Send_Json(w, http.StatusCreated, "User created successfully")
				return
			}
		} else {

			message.Send_Error(w, http.StatusInternalServerError, "Error in json message", "")
			return
		}

	} else {

		message.Send_Error(w, http.StatusMethodNotAllowed, "Method not allowed", "")
		return
	}

}
