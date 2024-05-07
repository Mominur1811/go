package handlers

import (
	"encoding/json"
	"fmt"
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

		var detect_user *User

		for i, user := range users {

			if user.ID == update_user.ID {
				detect_user = &users[i]
				break
			}
		}

		if detect_user == nil {

			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		detect_user.Name = update_user.Name
		detect_user.Password = update_user.Password

		// Return a success response
		fmt.Fprintf(w, "User updated successfully: %s", detect_user.Name)
		return

	}
}
