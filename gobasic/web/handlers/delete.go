package handlers

import (
	"encoding/json"

	"gobasic/model"
	"gobasic/web/message"
	"net/http"

	"fmt"
)

type UserData struct {
	ID int `json:"id"`
}

func Delete(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodDelete {

		var delete_user UserData
		err := json.NewDecoder(r.Body).Decode(&delete_user)

		if err != nil {

			message.Send_Error(w, http.StatusBadRequest, "Failed to load json", "")
			return
		}

		db := model.GetDB()

		query := `DELETE FROM employee WHERE id = $1`
		idToDelete := delete_user.ID // Replace 123 with the ID of the record you want to delete
		fmt.Println(idToDelete)
		_, err = db.Exec(query, idToDelete)
		if err != nil {
			fmt.Println("Error deleting record:", err)
			return
		} else {
			fmt.Println("deleted")
		}
		return

	}

}
