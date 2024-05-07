package handlers

import (
	"encoding/json"
	"fmt"
	"gobasic/web/message"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

var users []User

func Create(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		var new_user User
		err := json.NewDecoder(r.Body).Decode(&new_user)

		if err == nil {
			users = append(users, new_user)
			fmt.Fprintf(w, "Register Successfuly: %s", new_user.Name)
			return
		}
		message.Send_Error(w, http.StatusInternalServerError, "Error in json message", "")

		return
	}
	message.Send_Error(w, http.StatusMethodNotAllowed, "Method not allowed", "")
}
