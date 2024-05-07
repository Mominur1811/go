package handlers

import (
	"encoding/json"
	"fmt"
	"gobasic/web/message"
	"net/http"
)

func deleteElement(slice []User, index int) []User {

	// Create a new slice with the element at the specified index removed
	return append(slice[:index], slice[index+1:]...)
}

func Delete(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodDelete {

		var delete_user User
		err := json.NewDecoder(r.Body).Decode(&delete_user)

		if err != nil {

			message.Send_Error(w, http.StatusBadRequest, "Failed to load json", "")
			return
		}

		to_del := -1

		for i, user := range users {

			if user.ID == delete_user.ID {
				to_del = i
				break
			}
		}

		if to_del != -1 {

			users = deleteElement(users, to_del)
			fmt.Fprintf(w, "User deleted successfully: %s", delete_user.Name)
			return
		}

		message.Send_Error(w, http.StatusBadRequest, "User not found", "")
		return

	}

}
